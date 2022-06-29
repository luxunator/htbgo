package htbgo

import (
	"net/http"
)

type Session struct {
	Token string
	Client *http.Client
	Headers http.Header
}
