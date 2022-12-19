package nlibgo

import (
	"fmt"
	"log"
)

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

func (c *Client) logToStdout(level string, message string, details interface{}) {
	log.Println(level, message, details)
}

func (c *Client) log(level string, message string, args ...interface{}) error {
	details := arrayToMap(args...)
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
