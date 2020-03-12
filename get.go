package services

import (
	"net/http"
)

type get struct {
	url string
	params map[string]string
	headers map[string]string
}

func Get(url string)*get{
	var obj = &get{url:url}
	return obj
}

func (s *get) Header(key string, value string) *get{
	if s.headers == nil {
		s.headers = make(map[string]string)
	}
	s.headers[key] = value
	return s
}

func (s *get) Param(key string, value string) *get{
	if s.params == nil {
		s.params = make(map[string]string)
	}
	s.params[key] = value
	return s
}

func (s *get) Send(){
	var req *http.Request
	var err error
	var resp *http.Response

	req, err = http.NewRequest("GET", s.url, nil)

	if err != nil {
		panic(err)
	}
	if s.headers != nil {
		for key, value := range s.headers{
			req.Header.Set(key, value)
		}
	}
	if s.params != nil {
		q := req.URL.Query()

		for key, value := range s.params {
			q.Add(key, value)
		}

		req.URL.RawQuery = q.Encode()
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
}
