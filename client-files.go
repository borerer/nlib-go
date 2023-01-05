package nlibgo

import (
	"fmt"
	"io"

	"github.com/borerer/nlib-go/network"
)

func (c *Client) GetFile(filename string) (io.ReadCloser, error) {
	req, err := network.NewHTTPRequestBuilder("GET", fmt.Sprintf("%s/api/app/files/get", c.Endpoint)).Query("file", filename).Build()
	if err != nil {
		return nil, err
	}
	res, err := network.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if !network.StatusOK(res.StatusCode) {
		return nil, fmt.Errorf("status code %d", res.StatusCode)
	}
	return res.Body, nil
}

func (c *Client) PutFile(filename string, reader io.Reader) error {
	req, err := network.NewHTTPRequestBuilder("PUT", fmt.Sprintf("%s/api/app/files/put", c.Endpoint)).Query("file", filename).Body(reader).Build()
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

func (c *Client) DeleteFile(filename string) error {
	req, err := network.NewHTTPRequestBuilder("DELETE", fmt.Sprintf("%s/api/app/files/delete", c.Endpoint)).Query("file", filename).Build()
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

func (c *Client) MustGetFile(filename string) io.ReadCloser {
	r, err := c.GetFile(filename)
	Must(err)
	return r
}

func (c *Client) MustPutFile(filename string, reader io.Reader) {
	Must(c.PutFile(filename, reader))
}

func (c *Client) MustDeleteFile(filename string) {
	Must(c.DeleteFile(filename))
}
