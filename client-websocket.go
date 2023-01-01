package nlibgo

import (
	"errors"

	"github.com/borerer/nlib-go/logs"
	"github.com/borerer/nlib-go/utils"
	nlibshared "github.com/borerer/nlib-shared/go"
	"go.uber.org/zap"
)

var (
	ErrUnknownRequestType = errors.New("unknown request type")
)

func (c *Client) requestHandler(req *nlibshared.WebSocketMessage) (*nlibshared.WebSocketMessage, error) {
	switch req.SubType {
	case nlibshared.WebSocketMessageSubTypeCallFunction:
		var input nlibshared.PayloadCallFunctionRequest
		if err := utils.DecodeStruct(req.Payload, &input); err != nil {
			return nil, err
		}
		res := c.callFunction(&input)
		return &nlibshared.WebSocketMessage{
			SubType: nlibshared.WebSocketMessageSubTypeCallFunction,
			Payload: res,
		}, nil
	default:
		logs.Error("unknown request type", zap.Error(ErrUnknownRequestType), zap.Any("req", req))
		return nil, ErrUnknownRequestType
	}
}
