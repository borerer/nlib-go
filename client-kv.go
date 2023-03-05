package nlibgo

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/borerer/nlib-go/network"
)

func (c *Client) GetKey(key string) (string, error) {
	req, err := network.NewHTTPRequestBuilder("GET", fmt.Sprintf("%s/api/app/kv/get", c.Endpoint)).Query("key", key).Build()
	if err != nil {
		return "", err
	}
	res, err := network.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	if !network.StatusOK(res.StatusCode) {
		return "", fmt.Errorf("http error, path: %s, status code %d", "/api/app/kv/get", res.StatusCode)
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (c *Client) GetJSON(key string, res interface{}) error {
	val, err := c.GetKey(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), res)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SetKey(key string, value string) error {
	req, err := network.NewHTTPRequestBuilder("GET", fmt.Sprintf("%s/api/app/kv/set", c.Endpoint)).Query("key", key).Query("value", value).Build()
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

func (c *Client) MustGetKey(key string) string {
	value, err := c.GetKey(key)
	Must(err)
	return value
}

func (c *Client) MustGetJSON(key string, res interface{}) {
	Must(c.GetJSON(key, res))
}

func (c *Client) MustSetKey(key string, value string) {
	Must(c.SetKey(key, value))
}
