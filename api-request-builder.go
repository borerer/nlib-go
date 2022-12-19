package nlibgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIRequestBuilder struct {
	Endpoint string
	AppID    string

	// File APIs
	EndpointGetFile    string
	EndpointPutFile    string
	EndpointDeleteFile string
	EndpointFileStats  string
	EndpointListFolder string

	// KV APIs
	EndpointGetKey string
	EndpointSetKey string

	// Database APIs
	EndpointInsertDocument string
	EndpointUpdateDocument string
	EndpointDeleteDocument string
	EndpointFindDocuments  string

	// Function APIs
	EndpointRegisterFunction string

	// Logs APIs
	EndpointAddLogs string
}

func NewRequestBuilder(endpoint string, appID string) *APIRequestBuilder {
	r := &APIRequestBuilder{
		Endpoint: endpoint,
		AppID:    appID,
	}
	r.buildEndpoints()
	return r
}

func (r *APIRequestBuilder) buildEndpoints() {
	r.EndpointGetFile = fmt.Sprintf("%s/api/file/get", r.Endpoint)
	r.EndpointPutFile = fmt.Sprintf("%s/api/file/put", r.Endpoint)
	r.EndpointRegisterFunction = fmt.Sprintf("%s/api/function/register", r.Endpoint)
	r.EndpointAddLogs = fmt.Sprintf("%s/api/app/logs/log", r.Endpoint)
	r.EndpointGetKey = fmt.Sprintf("%s/api/app/kv/get", r.Endpoint)
	r.EndpointSetKey = fmt.Sprintf("%s/api/app/kv/set", r.Endpoint)
}

func (b *APIRequestBuilder) AddLogs(level string, message string, details interface{}) (*http.Request, error) {
	body := AddLogsRequest{
		Level:   level,
		Message: message,
		Details: details,
	}
	buf, _ := json.Marshal(body)
	println(string(buf))
	return NewHTTPRequestBuilder().Method("POST").BaseURL(b.EndpointAddLogs).Body(body).Build()
}

func (b *APIRequestBuilder) GetFile(filename string) (*http.Request, error) {
	return NewHTTPRequestBuilder().Method("GET").BaseURL(b.EndpointGetFile).Query("app", b.AppID).Query("file", filename).Build()
}

func (b *APIRequestBuilder) PutFile(filename string, reader io.Reader) (*http.Request, error) {
	return NewHTTPRequestBuilder().Method("PUT").BaseURL(b.EndpointPutFile).Query("app", b.AppID).Query("file", filename).Body(reader).Build()
}

func (b *APIRequestBuilder) GetKey(key string) (*http.Request, error) {
	return NewHTTPRequestBuilder().Method("GET").BaseURL(b.EndpointGetKey).Query("key", key).Build()
}

func (b *APIRequestBuilder) SetKey(key string, value string) (*http.Request, error) {
	return NewHTTPRequestBuilder().Method("GET").BaseURL(b.EndpointSetKey).Query("key", key).Query("value", value).Build()
}
