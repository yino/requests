package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"encoding/json"
)

type HttpRequest struct {
	Header  http.Header
	Request *http.Request
	Response Response
	FormData map[string][]string
}

func (req *HttpRequest) Post() Response {
	return Response{}
}

func (req *HttpRequest) Get() Response {
	return Response{}
}

func (req *HttpRequest) SetHeader(key, value string) Request {
	req.Header.Add(key, value)
	return req
}
func (req *HttpRequest) SetJsonBody(data interface{}) Request {
	btys, e := json.Marshal(data)
	if e != nil {
		req.Response.Error = e
		return req
	}

	req.Request.Body = ioutil.NopCloser(bytes.NewReader(btys))
	req.Request.ContentLength = int64(len(btys))
	req.Request.Header.Set("Content-Type", "application/json")
	return req
}
func (req *HttpRequest) SetBody(data map[string]interface{}) Request {

	var postValue url.Values

	for key,val := range data {
		var value []string
		value = append(value, val.(string))
		postValue[key] =value
	}
	req.FormData = postValue
	return req
}
func (req *HttpRequest) SetTimeout(duration time.Duration) Request {
	return req
}
func (req *HttpRequest) SetKeepAlive(duration time.Duration) Request {
	return req
}
