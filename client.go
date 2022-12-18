package nlibgo

import (
	"io"
	"sync"
	"time"

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
	go func() {
		for {
			err := c.connect()
			if err != nil {
				println(err.Error())
				time.Sleep(time.Second)
				continue
			}
			err = c.listenWebSocketMessages()
			if err != nil {
				println(err.Error())
				time.Sleep(time.Second)
			}
		}
	}()
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

func (c *Client) GetKey(key string) (string, error) {
	return c.getKey(key)
}

func (c *Client) SetKey(key string, value string) error {
	return c.setKey(key, value)
}

func (c *Client) RegisterFunction(funcName string, f NLIBFunc, opts ...RegisterFunctionOptions) error {
	c.registeredFunctions.Store(funcName, f)
	return nil
}
