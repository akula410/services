package services

import (
	"io"
	"net/http"
	)

type _interface interface {
	Header(key string, value string) _interface
	Param(key string, value string) _interface
	Cookie(name string, value string) _interface
	GetRequest() *http.Request
	GetResponse() *http.Response
	GetResponseBody() []byte
	GetHeader(key string)string
	GetHeaders()map[string]string
	GetCookie(name string)string
	SetCookie(cookie *http.Cookie) _interface
	Send(body ...io.Reader)
}
