package nlibgo

import (
	"net/http"
)

var (
	httpClient *http.Client
)

func init() {
	httpClient = http.DefaultClient
}

func StatusOK(code int) bool {
	return code >= 200 && code < 300
}
