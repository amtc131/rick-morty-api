package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type HttpClientApi struct {
	client  *http.Client
	host    string
	path    string
	params  map[string]interface{}
	method  string
	headers http.Header
	body    interface{}
}

func NewHttpClientApi(host string, client *http.Client) *HttpClientApi {
	return &HttpClientApi{
		client: client,
		host:   host,
		params: make(map[string]interface{}),
	}
}

func (r *HttpClientApi) Path(path string) *HttpClientApi {
	r.path = path
	return r
}

func (r *HttpClientApi) Params(params map[string]interface{}) *HttpClientApi {
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

	endpoint := ""
	hasParams := false

	if len(r.params) != 0 {
		for k, v := range r.params {
			switch v.(type) {
			case map[string]string:
				if k == "params" {
					hasParams = true
				}
			case int:
				if k == "page" {
					hasParams = true
					r.params[k] = strconv.FormatInt(int64(v.(int)), 10)
				}
				delete(r.params, k)
			case map[string]int:
				if k == "params" { // id
					integer := strconv.FormatInt(int64(v.(map[string]int)["integer"]), 10)
					endpoint = endpoint + integer
					delete(r.params, "params")
				}
			case []int:
				params := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v.([]int))), ","), "[]")
				endpoint = endpoint + params
			default:
				err := sliceIntToString(v.([]int), ",")
				return nil, errors.New(err)
			}
		}
	}

	req, err := http.NewRequest(r.method, r.host+endpoint, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()

	if hasParams {
		for key, value := range r.params["params"].(map[string]string) {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := r.client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp, err
}

func sliceIntToString(slice []int, join string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), join), "[]")
}
