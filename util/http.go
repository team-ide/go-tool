package util

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	HttpClient      = NewHttpClient()
	ContentTypeJson = "application/json"
)

func NewHttpClient() *http.Client {

	// 创建传输对象
	transport := &http.Transport{
		MaxIdleConns:    10,
		MaxConnsPerHost: 10,
		IdleConnTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			// 指定不校验 SSL/TLS 证书
			InsecureSkipVerify: true,
		},
	}
	// 创建 HTTP 客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}
	return client
}

func BindJsonBody(body io.ReadCloser, obj interface{}) (str string, err error) {
	if body == nil {
		err = errors.New("body is null")
		return
	}
	defer func() { _ = body.Close() }()
	bs, err := io.ReadAll(body)
	if err != nil {
		return
	}
	str = string(bs)
	//fmt.Println("response body:", str)
	err = json.Unmarshal(bs, obj)
	if err != nil {
		err = errors.New("json [" + str + "] to object error:" + err.Error())
		return
	}
	return
}

func BindXmlBody(body io.ReadCloser, obj interface{}) (bodyXml string, err error) {
	if body == nil {
		err = errors.New("body is null")
		return
	}
	defer func() { _ = body.Close() }()
	bs, err := io.ReadAll(body)
	if err != nil {
		return
	}
	bodyXml = string(bs)
	err = xml.Unmarshal(bs, obj)
	if err != nil {
		err = errors.New("xml [" + string(bs) + "] to object error:" + err.Error())
		return
	}
	return
}

func GetJson[T any](url string, response T) (res T, body string, err error) {
	cR, err := HttpClient.Get(url)
	if err != nil {
		return
	}
	body, err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}

func GetJsonHeader[T any](url string, header http.Header, response T) (res T, body string, err error) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	r.Header = header
	cR, err := HttpClient.Do(r)
	if err != nil {
		return
	}
	body, err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}

func PostJson[T any](url string, data any, response T) (res T, body string, err error) {
	bs, err := ObjToJsonBytes(data)
	if err != nil {
		return
	}

	cR, err := HttpClient.Post(url, ContentTypeJson, bytes.NewReader(bs))
	if err != nil {
		return
	}
	body, err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}

func PostJsonHeader[T any](url string, header http.Header, data any, response T) (res T, body string, err error) {
	bs, err := ObjToJsonBytes(data)
	if err != nil {
		return
	}

	r, err := http.NewRequest("POST", url, bytes.NewReader(bs))
	if err != nil {
		return
	}
	r.Header = header
	cR, err := HttpClient.Do(r)
	if err != nil {
		return
	}
	//fmt.Println("response code:", cR.StatusCode)
	body, err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}

func PostForm[T any](url string, data url.Values, response T) (res T, body string, err error) {

	cR, err := HttpClient.PostForm(url, data)
	if err != nil {
		return
	}
	body, err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}

func PostFormHeader[T any](url string, header http.Header, data url.Values, response T) (res T, body string, err error) {

	r, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	r.Header = header
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	cR, err := HttpClient.Do(r)
	if err != nil {
		return
	}
	body, err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}

func PostFormBytes(url string, data url.Values) (res []byte, err error) {
	cR, err := HttpClient.PostForm(url, data)
	if err != nil {
		return
	}
	body := cR.Body
	if body == nil {
		err = errors.New("body is null")
		return
	}
	defer func() { _ = body.Close() }()
	res, err = io.ReadAll(body)
	if err != nil {
		return
	}
	return
}
