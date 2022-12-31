package network

import (
	"net/http"
)

var (
	HttpClient *http.Client
)

func init() {
	HttpClient = http.DefaultClient
}

func StatusOK(code int) bool {
	return code >= 200 && code < 300
}
