package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	sjson "github.com/bitly/go-simplejson"
)

type RequestHeader struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Request struct {
	Url        string            `json:"url"`
	Method     string            `json:"method"`
	Params     url.Values        `json:"params"`
	Payload    interface{}       `json:"payload"`
	OperatorId string            `json:"operatorId"`
	Header     map[string]string `json:"header"`
}

func (r *Request) SetHeader(key string, value string) {
	if r.Header == nil {
		r.Header = map[string]string{}
	}
	r.Header[key] = value
}

func (r *Request) Do() ([]byte, error) {
	fmt.Println("request:", r.Url)
	js, err := json.Marshal(r.Payload)
	if err != nil {
		return nil, err
	}
	if r.Params.Encode() != "" {
		r.Url = r.Url + "?" + r.Params.Encode()
	}
	request, err := http.NewRequest(r.Method, r.Url, bytes.NewReader(js))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	if r.OperatorId != "" {
		request.Header.Set("operatorId", r.OperatorId)
	}
	for key, value := range r.Header {
		request.Header.Set(key, value)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 300 {
		sjError, err := sjson.NewJson(result)
		if err != nil {
			return nil, err
		}
		errStr := sjError.GetPath("errors").GetIndex(0).GetPath("detail").MustString()
		return nil, errors.New(errStr)
	}

	return result, nil
}

//get helper
func RequestHelper(url string, params url.Values, operatorId string) *Request {
	req := new(Request)
	req.OperatorId = operatorId

	// 1： 获取所有子级列表
	req.Url = url
	req.Params = params

	return req
}

//get helper
func RequestGetHelper(url string, params url.Values, operatorId string) *Request {
	req := new(Request)
	req.Method = "GET"
	req.OperatorId = operatorId

	// 1： 获取所有子级列表
	req.Url = url
	req.Params = params

	return req
}

//post helper
func RequestPostHelper(url string, Payload interface{}, operatorId string) *Request {
	req := new(Request)
	req.Method = "POST"
	req.OperatorId = operatorId

	// 1： 获取所有子级列表
	req.Url = url
	req.Payload = Payload

	return req
}
