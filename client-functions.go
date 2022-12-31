package nlibgo

import (
	"errors"
	"net/http"

	"github.com/borerer/nlib-go/utils"
	nlibshared "github.com/borerer/nlib-shared/go"
)

var (
	ErrInvalidFunction  = errors.New("invalid function")
	ErrFunctionNotFound = errors.New("function not found")
)

// expose to nlib-go users
type SimpleFunctionIn = nlibshared.SimpleFunctionIn
type SimpleFunctionOut = nlibshared.SimpleFunctionOut
type SimpleFunction = nlibshared.SimpleFunction

type HARFunctionIn = nlibshared.HARFunctionIn
type HARFunctionOut = nlibshared.HARFunctionOut
type HARFunction = nlibshared.HARFunction

func (c *Client) RegisterFunction(name string, f interface{}) error {
	req := &nlibshared.PayloadRegisterFunctionRequest{
		Name: name,
	}
	switch f.(type) {
	case SimpleFunction:
		req.UseHAR = false
	case HARFunction:
		req.UseHAR = true
	default:
		return ErrInvalidFunction
	}
	if _, err := c.registerFunction(req); err != nil {
		return err
	}
	c.registeredFunctions.Store(name, f)
	return nil
}

func (c *Client) registerFunction(req *nlibshared.PayloadRegisterFunctionRequest) (*nlibshared.PayloadRegisterFunctionResponse, error) {
	raw, err := c.socket.SendRequest(&nlibshared.WebSocketMessage{
		SubType: nlibshared.WebSocketMessageSubTypeRegisterFunction,
		Payload: req,
	})
	if err != nil {
		return nil, err
	}
	var res nlibshared.PayloadRegisterFunctionResponse
	if err := utils.DecodeStruct(raw, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) callSimpleFunction(req *nlibshared.PayloadCallFunctionRequest) *nlibshared.PayloadCallFunctionResponse {
	raw, ok := c.registeredFunctions.Load(req.Name)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: ErrFunctionNotFound,
		}
	}
	f, ok := raw.(SimpleFunction)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: ErrInvalidFunction,
		}
	}
	var input SimpleFunctionIn
	if err := utils.DecodeStruct(req.Request, &input); err != nil {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: err.Error(),
		}
	}
	output := f(input)
	return &nlibshared.PayloadCallFunctionResponse{
		Response: output,
	}
}

func (c *Client) callHARFunction(req *nlibshared.PayloadCallFunctionRequest) *nlibshared.PayloadCallFunctionResponse {
	raw, ok := c.registeredFunctions.Load(req.Name)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: HARFunctionOut{
				Status: http.StatusNotFound,
			},
		}
	}
	f, ok := raw.(HARFunction)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: HARFunctionOut{
				Status: http.StatusInternalServerError,
			},
		}
	}
	var input HARFunctionIn
	if err := utils.DecodeStruct(req.Request, &input); err != nil {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: HARFunctionOut{
				Status: http.StatusInternalServerError,
			},
		}
	}
	output := f(input)
	return &nlibshared.PayloadCallFunctionResponse{
		Response: output,
	}
}

func (c *Client) callFunction(req *nlibshared.PayloadCallFunctionRequest) *nlibshared.PayloadCallFunctionResponse {
	if req.UseHAR {
		return c.callHARFunction(req)
	} else {
		return c.callSimpleFunction(req)
	}
}
