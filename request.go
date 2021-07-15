package requests

import "time"

// Request req
type Request interface {
	Post() Response
	Get() Response
	SetHeader(key, value string) Request
	SetJsonBody(data interface{}) Request
	SetBody(data map[string]interface{}) Request
	SetTimeout(duration time.Duration) Request
	SetKeepAlive(duration time.Duration) Request
}
