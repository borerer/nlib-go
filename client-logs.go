package nlibgo

import (
	"fmt"
	"log"
)

func (c *Client) logToStdout(level string, message string, details interface{}) {
	log.Println(level, message, details)
}

func (c *Client) log(level string, message string, details interface{}) error {
	c.logToStdout(level, message, details)
	req, err := c.requestBuilder.AddLogs(level, message, details)
	if err != nil {
		return err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if !StatusOK(res.StatusCode) {
		return fmt.Errorf("status code %d", res.StatusCode)
	}
	return nil
}
