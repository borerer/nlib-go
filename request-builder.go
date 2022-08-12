package nlibgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RequestBuilder struct {
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

func NewRequestBuilder(endpoint string, appID string) *RequestBuilder {
	r := &RequestBuilder{
		Endpoint: endpoint,
		AppID:    appID,
	}
	r.buildEndpoints()
	return r
}

func (r *RequestBuilder) buildEndpoints() {
	r.EndpointGetFile = fmt.Sprintf("%s/api/file/get", r.Endpoint)
	r.EndpointPutFile = fmt.Sprintf("%s/api/file/put", r.Endpoint)
	r.EndpointRegisterFunction = fmt.Sprintf("%s/api/function/register", r.Endpoint)
	r.EndpointAddLogs = fmt.Sprintf("%s/api/logs", r.Endpoint)
}

func (r *RequestBuilder) AddLogs(message string) (*http.Request, error) {
	u, err := url.Parse(r.EndpointAddLogs)
	if err != nil {
		return nil, err
	}
	values := u.Query()
	values.Add("app", r.AppID)
	values.Add("message", message)
	u.RawQuery = values.Encode()
	req, _ := http.NewRequest("POST", u.String(), nil)
	return req, nil
}

func (r *RequestBuilder) RegisterFunction(payload *RegisterFunctionRequest) (*http.Request, error) {
	u, err := url.Parse(r.EndpointRegisterFunction)
	if err != nil {
		return nil, err
	}
	values := u.Query()
	values.Add("app", r.AppID)
	u.RawQuery = values.Encode()
	buf, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(buf)
	req, _ := http.NewRequest("POST", u.String(), reader)
	return req, nil
}
