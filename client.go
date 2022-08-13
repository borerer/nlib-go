package nlibgo

import (
	"io"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Endpoint string
	AppID    string

	requestBuilder      *APIRequestBuilder
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

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
	Fatal = "fatal"
)

func (c *Client) Debug(message string) error {
	return c.log(Debug, message, nil)
}

func (c *Client) Info(message string) error {
	return c.log(Info, message, nil)
}

func (c *Client) Warn(message string) error {
	return c.log(Warn, message, nil)
}

func (c *Client) Error(message string) error {
	return c.log(Error, message, nil)
}

func (c *Client) Fatal(message string) error {
	return c.log(Fatal, message, nil)
}

func (c *Client) GetFile(filename string) (io.ReadCloser, error) {
	return c.getFile(filename)
}

func (c *Client) PutFile(filename string, reader io.Reader) error {
	return c.putFile(filename, reader)
}

func (c *Client) RegisterFunction(f NLIBFunc, opt RegisterFunctionOptions) error {
	c.registeredFunctions.Store(opt.FuncName, f)
	return nil
}
