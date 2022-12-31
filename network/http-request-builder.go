package network

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HTTPRequestBuilder struct {
	method  string
	baseURL string
	query   url.Values
	body    interface{}
}

func NewHTTPRequestBuilder(method string, baseURL string) *HTTPRequestBuilder {
	b := &HTTPRequestBuilder{
		method:  method,
		baseURL: baseURL,
	}
	b.query = make(url.Values)
	return b
}

func (b *HTTPRequestBuilder) Method(method string) *HTTPRequestBuilder {
	b.method = method
	return b
}

func (b *HTTPRequestBuilder) BaseURL(baseURL string) *HTTPRequestBuilder {
	b.baseURL = baseURL
	return b
}

func (b *HTTPRequestBuilder) Query(key string, value string) *HTTPRequestBuilder {
	b.query.Add(key, value)
	return b
}

func (b *HTTPRequestBuilder) Body(body interface{}) *HTTPRequestBuilder {
	b.body = body
	return b
}

func (b *HTTPRequestBuilder) Build() (*http.Request, error) {
	u, err := url.Parse(b.baseURL)
	if err != nil {
		return nil, err
	}
	u.RawQuery = b.query.Encode()
	var reader io.Reader
	if b.body == nil {
		// no-op
	} else if bodyReader, ok := b.body.(io.Reader); ok {
		reader = bodyReader
	} else if bodyStr, ok := b.body.(string); ok {
		reader = strings.NewReader(bodyStr)
	} else if bodyBytes, ok := b.body.([]byte); ok {
		reader = bytes.NewReader(bodyBytes)
	} else {
		buf, err := json.Marshal(b.body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(buf)
	}
	return http.NewRequest(b.method, u.String(), reader)
}
