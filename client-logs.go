package nlibgo

import "fmt"

func (c *Client) log(level string, message string, details interface{}) error {
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
