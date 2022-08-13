package nlibgo

import (
	"fmt"
	"io"
)

func (c *Client) getFile(filename string) (io.ReadCloser, error) {
	req, err := c.requestBuilder.GetFile(filename)
	if err != nil {
		return nil, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if !StatusOK(res.StatusCode) {
		return nil, fmt.Errorf("status code %d", res.StatusCode)
	}
	return res.Body, nil
}

func (c *Client) putFile(filename string, reader io.Reader) error {
	req, err := c.requestBuilder.PutFile(filename, reader)
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
