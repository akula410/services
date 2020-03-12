package services


import (
	"io"
	"io/ioutil"
	"net/http"
)

type abstract struct {
	url string
	method string
	params map[string]string
	headers map[string]string
	cookies []*http.Cookie
	body []byte
	request *http.Request
	response *http.Response
}

func (s *abstract) setUrl(url string){
	s.url = url
}
func (s *abstract) setMethod(method string){
	s.method = method
}

func (s *abstract) Header(key string, value string) _interface{
	if s.headers == nil {
		s.headers = make(map[string]string)
	}
	s.headers[key] = value
	return s
}

func (s *abstract) Param(key string, value string) _interface{
	if s.params == nil {
		s.params = make(map[string]string)
	}
	s.params[key] = value
	return s
}

func (s *abstract) Cookie(name string, value string) _interface{
	if s.cookies == nil {
		s.cookies = make([]*http.Cookie, 0)
	}
	s.cookies = append(s.cookies, &http.Cookie{Name:name, Value:value})
	return s
}

func (s *abstract) GetRequest() *http.Request{
	return s.request
}

func (s *abstract) GetResponse() *http.Response{
	return s.response
}

func (s *abstract) GetResponseBody() []byte{
	return s.body
}

func (s *abstract) GetHeader(key string)string{
	var result string
	if s.response != nil {
	Exit:
		for HeaderKey, values := range s.response.Header {
			if HeaderKey == key {
				for _, value := range values{
					result = value
					break Exit
				}
			}
		}
	}
	return result
}

func (s *abstract) GetHeaders()map[string]string{
	var result = make(map[string]string)
	if s.response != nil {
		for key, values := range s.response.Header {
			for _, value := range values{
				result[key] = value
				break
			}
		}
	}
	return result
}

func (s *abstract) GetCookie(name string)string{
	var result string
	if s.response != nil {
		for _, cookie := range s.response.Cookies() {
			if cookie.Name == name {
				result = cookie.Value
				break
			}
		}
	}
	return result
}

func (s *abstract) SetCookie(cookie *http.Cookie) _interface{
	if s.cookies == nil {
		s.cookies = make([]*http.Cookie, 0)
	}
	s.cookies = append(s.cookies, cookie)
	return s
}

func (s *abstract) Send(body ...io.Reader){
	var err error
	var send io.Reader
	if body != nil && len(body) > 0 {
		send = body[0]
	}

	s.request, err = http.NewRequest(s.method, s.url, send)

	if err != nil {
		panic(err)
	}
	if s.headers != nil {
		for key, value := range s.headers{
			s.request.Header.Set(key, value)
		}
	}
	if s.cookies != nil {
		for _, cookie := range s.cookies{
			s.request.AddCookie(cookie)
		}
	}

	if s.params != nil {
		q := s.request.URL.Query()

		for key, value := range s.params {
			q.Add(key, value)
		}

		s.request.URL.RawQuery = q.Encode()
	}
	s.response, err = http.DefaultClient.Do(s.request)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = s.response.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	s.body, err = ioutil.ReadAll(s.response.Body)
}

