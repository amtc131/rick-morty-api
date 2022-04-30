package server

import (
	"fmt"
	"net/http"
	"net/url"
)

type HttpClientApi struct {
	client  *http.Client
	host    string
	path    string
	params  url.Values
	method  string
	headers http.Header
	body    interface{}
}

func NewHttpClientApi(host string, client *http.Client) *HttpClientApi {
	return &HttpClientApi{
		client: client,
		host:   host,
	}
}

func (r *HttpClientApi) Path(path string) *HttpClientApi {
	r.path = path
	return r
}

func (r *HttpClientApi) Params(params url.Values) *HttpClientApi {
	r.params = params
	return r
}

func (r *HttpClientApi) Method(method string) *HttpClientApi {
	r.method = method
	return r
}

func (r *HttpClientApi) Headers(headers http.Header) *HttpClientApi {
	r.headers = headers
	return r
}

func (r *HttpClientApi) Body(body interface{}) *HttpClientApi {
	r.body = body
	return r
}

func (r *HttpClientApi) Do() (*http.Response, error) {
	req, err := http.NewRequest(r.method, r.host, nil)
	if err != nil {
		fmt.Errorf("Unable to new Request", err.Error())
	}

	resp, err := r.client.Do(req)
	if err != nil {
		fmt.Errorf("Unable to client Do", err.Error())
	}

	return resp, err
}
