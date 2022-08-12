package nlibgo

import (
	"encoding/json"
	"io"
	"net/http"
)

var (
	client *http.Client
)

func init() {
	client = http.DefaultClient
}

func DoRequest(req *http.Request, result interface{}) error {
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if result != nil {
		defer res.Body.Close()
		buf, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(buf, result)
		if err != nil {
			return err
		}
	}
	return nil
}
