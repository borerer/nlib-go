package nlibgo

import (
	"encoding/json"
	"net/http"
)

const (
	ContentTypeTextPlain       = "text/plain"
	ContentTypeApplicationJSON = "application/json"
)

func NewResponse(statusCode int, content string, contentType string) *Response {
	res := &Response{}
	res.Status = int64(statusCode)
	res.Content = Content{
		Text: &content,
	}
	res.Headers = append(res.Headers, Header{
		Name:  "Content-Type",
		Value: contentType,
	})
	return res
}

var Err404 = NewResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), "")
var Err405 = NewResponse(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), "")
var Err500 = NewResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")

func Text(s string) *Response {
	return NewResponse(http.StatusOK, s, ContentTypeTextPlain)
}

func JSON(v interface{}) *Response {
	buf, err := json.Marshal(v)
	if err != nil {
		return Error_(err)
	}
	return NewResponse(http.StatusOK, string(buf), ContentTypeApplicationJSON)
}

func Error_(err error) *Response {
	return NewResponse(http.StatusInternalServerError, err.Error(), ContentTypeTextPlain)
}

func GetQuery(req *Request, key string) string {
	for _, query := range req.QueryString {
		if query.Name == key {
			return query.Value
		}
	}
	return ""
}

func GetHeader(req *Request, key string) string {
	for _, header := range req.Headers {
		if header.Name == key {
			return header.Value
		}
	}
	return ""
}

func GetCookie(req *Request, key string) string {
	for _, cookie := range req.Cookies {
		if cookie.Name == key {
			return cookie.Value
		}
	}
	return ""
}
