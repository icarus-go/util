package data

import (
	"net/http"
	"net/url"
)

type I struct {
	Body        Body
	FormData    url.Values
	QueryString url.Values
	Header      http.Header
}
