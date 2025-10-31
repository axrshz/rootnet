package http

import (
	"fmt"
	"strings"

	"github.com/axrshz/rootnet/http/status"
)

type Response struct {
	Version    string
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       []byte
}

func NewResponse() *Response {
	return &Response{
		Version: "HTTP/1.1",
		Headers: make(map[string]string),
	}
}

func FormatResponse(r *Response) []byte {
	var builder strings.Builder

	// Status line
	builder.WriteString(fmt.Sprintf("%s %d %s\r\n", r.Version, r.StatusCode, r.StatusText))

	// Headers
	for key, value := range r.Headers {
		builder.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	// Empty line
	builder.WriteString("\r\n")

	// Body
	return append([]byte(builder.String()), r.Body...)
}

func StatusText(code int) string {
	return status.Text(code)
}

func (r *Response) SetStatus(code int) {
    r.StatusCode = code
    r.StatusText = status.Text(code)
}

func (r *Response) SetHeader(key, value string) {
    r.Headers[key] = value
}

func (r *Response) SetBody(body []byte) {
    r.Body = body
}