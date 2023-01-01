package nlibgo

import (
	"fmt"
	"log"

	"github.com/borerer/nlib-go/network"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelFatal = "fatal"
)

func (c *Client) Debug(message string, args ...interface{}) error {
	return c.log(LevelDebug, message, args...)
}

func (c *Client) Info(message string, args ...interface{}) error {
	return c.log(LevelInfo, message, args...)
}

func (c *Client) Warn(message string, args ...interface{}) error {
	return c.log(LevelWarn, message, args...)
}

func (c *Client) Error(message string, args ...interface{}) error {
	return c.log(LevelError, message, args...)
}

func (c *Client) Fatal(message string, args ...interface{}) error {
	return c.log(LevelFatal, message, args...)
}

func arrayToMap(args ...interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	for i := 0; i+1 < len(args); i += 2 {
		s, ok := args[i].(string)
		if !ok {
			continue
		}
		res[s] = args[i+1]
	}
	return res
}

func (c *Client) logToStdout(level string, message string, details map[string]interface{}) {
	if len(details) > 0 {
		log.Printf("[%s] %s %+v", level, message, details)
	} else {
		log.Printf("[%s] %s", level, message)
	}
}

func (c *Client) log(level string, message string, args ...interface{}) error {
	details := arrayToMap(args...)
	c.logToStdout(level, message, details)
	body := map[string]interface{}{
		"level":   level,
		"message": message,
		"details": details,
	}
	req, err := network.NewHTTPRequestBuilder("POST", fmt.Sprintf("%s/api/app/logs/log", c.Endpoint)).Body(body).Build()
	if err != nil {
		return err
	}
	res, err := network.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if !network.StatusOK(res.StatusCode) {
		return fmt.Errorf("status code %d", res.StatusCode)
	}
	return nil
}
