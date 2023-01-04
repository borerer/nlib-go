package nlibgo

import (
	"errors"
	"net/http"

	"github.com/borerer/nlib-go/har"
	"github.com/borerer/nlib-go/utils"
	nlibshared "github.com/borerer/nlib-shared/go"
)

var (
	ErrInvalidFunctionType = errors.New("invalid function type")
	ErrFunctionNotFound    = errors.New("function not found")
)

type FunctionIn = har.Request
type FunctionOut = har.Response
type Function = func(*FunctionIn) (*FunctionOut, error)

func (c *Client) RegisterFunction(name string, f Function) error {
	req := &nlibshared.PayloadRegisterFunctionRequest{
		Name: name,
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

func (c *Client) callFunction(req *nlibshared.PayloadCallFunctionRequest) *nlibshared.PayloadCallFunctionResponse {
	raw, ok := c.registeredFunctions.Load(req.Name)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: *har.NewResponse(http.StatusNotFound, "not found", har.ContentTypeTextPlain),
		}
	}
	f, ok := raw.(Function)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: *har.Error(ErrInvalidFunctionType),
		}
	}
	output, err := f(&req.Request)
	if err != nil {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: *har.Error(ErrInvalidFunctionType),
		}
	}
	return &nlibshared.PayloadCallFunctionResponse{
		Response: *output,
	}
}
