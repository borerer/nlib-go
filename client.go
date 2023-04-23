package nlibgo

import (
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/borerer/nlib-go/logs"
	"github.com/borerer/nlib-go/socket"
	"go.uber.org/zap"
)

type Client struct {
	Endpoint  string
	AppID     string
	DebugMode bool
	LogsSkip  int

	socket              *socket.Socket
	registeredFunctions sync.Map
}

var (
	ErrMissingEndpoint = errors.New("missing endpoint")
	ErrMissingAppID    = errors.New("missing app id")
)

func NewClient(endpoint string, appID string) *Client {
	c := &Client{
		Endpoint:  endpoint,
		AppID:     appID,
		DebugMode: false,
		LogsSkip:  1,
	}
	return c
}

func (c *Client) SetDebugMode(debugMode bool) {
	if debugMode {
		c.DebugMode = true
		logs.SetLevel(zap.DebugLevel)
	} else {
		c.DebugMode = false
		logs.SetLevel(zap.InfoLevel)
	}
}

func (c *Client) connectOnce() error {
	if len(c.Endpoint) == 0 {
		return ErrMissingEndpoint
	}
	if len(c.AppID) == 0 {
		return ErrMissingAppID
	}
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
	c.socket = socket.NewSocket(u.String())
	c.socket.SetRequestHandler(c.requestHandler)
	if err = c.socket.Connect(); err != nil {
		return err
	}
	return nil
}

func (c *Client) Connect() error {
	if err := c.connectOnce(); err != nil {
		return err
	}
	skipOnce := true
	go func() {
		for {
			if skipOnce {
				skipOnce = false
			} else {
				if err := c.connectOnce(); err != nil {
					logs.Error("error", zap.Error(err))
					time.Sleep(time.Second)
					continue
				}
			}
			if err := c.socket.ListenWebSocketMessages(); err != nil {
				logs.Error("error", zap.Error(err))
				time.Sleep(time.Second)
			}
		}
	}()
	return nil
}
