package nlibgo

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

func (c *Client) connect() error {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return err
	}
	if u.Scheme == "https" {
		u.Scheme = "wss"
	} else {
		u.Scheme = "ws"
	}
	u.Path = fmt.Sprintf("/api/app/%s/ws", c.AppID)
	c.websocketConnection, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) handleWebSocketRequests(message *WebSocketMessage) error {
	switch message.SubType {
	case WebSocketSubTypeCallFunction:
		var callFunctionReq WebSocketCallFunctionReq
		if err := mapstructure.Decode(message.Payload, &callFunctionReq); err != nil {
			return err
		}
		funcRaw, ok := c.registeredFunctions.Load(callFunctionReq.FuncName)
		if !ok {
			return nil
		}
		nlibFunc, ok := funcRaw.(NLIBFunc)
		if !ok {
			return nil
		}
		res := nlibFunc(callFunctionReq.Params)
		callFunctionRes := WebSocketCallFunctionRes{
			FuncName: callFunctionReq.FuncName,
			Response: res,
		}
		websocketRes := WebSocketMessage{
			MessageID:     uuid.NewString(),
			PairMessageID: message.MessageID,
			Type:          WebSocketTypeResponse,
			SubType:       WebSocketSubTypeCallFunction,
			Timestamp:     time.Now().UnixMilli(),
			Payload:       callFunctionRes,
		}
		err := c.websocketConnection.WriteJSON(websocketRes)
		if err != nil {
			print(234)
		}
	}
	return nil
}

func (c *Client) handleClose(code int, text string) error {
	println("disconnected", code, text)
	return nil
}

func (c *Client) listenWebSocketMessages() error {
	c.websocketConnection.SetCloseHandler(c.handleClose)
	for {
		var message WebSocketMessage
		err := c.websocketConnection.ReadJSON(&message)
		if err != nil {
			return err
		}
		switch message.Type {
		case WebSocketTypeRequest:
			c.handleWebSocketRequests(&message)
		case WebSocketTypeResponse:
			print(123)
		}
	}
}
