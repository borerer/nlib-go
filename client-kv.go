package nlibgo

import (
	"fmt"
	"io"
)

func (c *Client) getKey(key string) (string, error) {
	req, err := c.requestBuilder.GetKey(key)
	if err != nil {
		return "", err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	if !StatusOK(res.StatusCode) {
		return "", fmt.Errorf("status code %d", res.StatusCode)
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (c *Client) setKey(key string, value string) error {
	req, err := c.requestBuilder.SetKey(key, value)
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
