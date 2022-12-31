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

func (c *Client) RegisterFunction(name string, f interface{}) error {
	req := &nlibshared.PayloadRegisterFunctionRequest{
		Name: name,
	}
	switch f.(type) {
	case nlibshared.SimpleFunction:
		req.UseHAR = false
	case nlibshared.HARFunction:
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
	f, ok := raw.(nlibshared.SimpleFunction)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: ErrInvalidFunction,
		}
	}
	var input nlibshared.SimpleFunctionIn
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
			Response: nlibshared.HARFunctionOut{
				Status: http.StatusNotFound,
			},
		}
	}
	f, ok := raw.(nlibshared.HARFunction)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: nlibshared.HARFunctionOut{
				Status: http.StatusInternalServerError,
			},
		}
	}
	var input nlibshared.HARFunctionIn
	if err := utils.DecodeStruct(req.Request, &input); err != nil {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: nlibshared.HARFunctionOut{
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
