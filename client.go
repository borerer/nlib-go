package nlibgo

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Endpoint string
	AppID    string

	requestBuilder      *RequestBuilder
	registeredFunctions sync.Map

	websocketConnection *websocket.Conn
}

func NewClient(endpoint string, appID string) *Client {
	c := &Client{
		Endpoint:       endpoint,
		AppID:          appID,
		requestBuilder: NewRequestBuilder(endpoint, appID),
	}
	Must(c.connect())
	go c.listenWebSocketMessages()
	return c
}

func (c *Client) LogMessage(message string) error {
	req, err := c.requestBuilder.AddLogs(message)
	if err != nil {
		return err
	}
	err = DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

type RegisterFunctionRequest struct {
	FuncName string `json:"func_name"`
	Endpoint string `json:"endpoint"`
}

type RegisterFunctionResponse struct {
	ID string `json:"id"`
}

type RegisterFunctionOptions struct {
	FuncName string
	Endpoint string
}

type NLIBFunc func(string) string

func (c *Client) RegisterFunction(f NLIBFunc, opt RegisterFunctionOptions) error {
	c.registeredFunctions.Store(opt.FuncName, f)
	// reqBody := &RegisterFunctionRequest{
	// 	ID: options.ID,
	// }
	// req, err := c.requestBuilder.RegisterFunction(reqBody)
	// if err != nil {
	// 	return err
	// }
	// var res RegisterFunctionResponse
	// err = DoRequest(req, &res)
	// if err != nil {
	// 	return err
	// }
	return nil
}
