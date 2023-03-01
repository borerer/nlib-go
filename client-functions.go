package nlibgo

import (
	"errors"
	"net/http"

	"github.com/borerer/nlib-go/utils"
	nlibshared "github.com/borerer/nlib-shared/go"
)

var (
	ErrInvalidFunctionType = errors.New("invalid function type")
	ErrFunctionNotFound    = errors.New("function not found")
)

func (c *Client) RegisterFunction(name string, f func(*nlibshared.Request) *nlibshared.Response) error {
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
			Response: *NewResponse(http.StatusNotFound, "not found", ContentTypeTextPlain),
		}
	}
	f, ok := raw.(func(*nlibshared.Request) *nlibshared.Response)
	if !ok {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: *Error_(ErrInvalidFunctionType),
		}
	}
	var output *nlibshared.Response
	panicError := Safe(func() {
		output = f(&req.Request)
	})
	if panicError != nil {
		return &nlibshared.PayloadCallFunctionResponse{
			Response: *Error_(panicError),
		}
	}
	return &nlibshared.PayloadCallFunctionResponse{
		Response: *output,
	}
}
