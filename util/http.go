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
		Timeout:   15 * time.Second,
	}
	return client
}

func BindJsonBody(body io.ReadCloser, obj interface{}) (err error) {
	if body == nil {
		err = errors.New("body is null")
		return
	}
	defer func() { _ = body.Close() }()
	bs, err := io.ReadAll(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, obj)
	if err != nil {
		err = errors.New("json[" + string(bs) + "] to object error:" + err.Error())
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
		err = errors.New("xml[" + string(bs) + "] to object error:" + err.Error())
		return
	}
	return
}

func GetJson[T any](url string, response T) (res T, err error) {
	cR, err := HttpClient.Get(url)
	if err != nil {
		return
	}
	err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}

func PostJson[T any](url string, data any, response T) (res T, err error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return
	}

	cR, err := HttpClient.Post(url, ContentTypeJson, bytes.NewReader(bs))
	if err != nil {
		return
	}
	err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}

func PostForm[T any](url string, data url.Values, response T) (res T, err error) {

	cR, err := HttpClient.PostForm(url, data)
	if err != nil {
		return
	}
	err = BindJsonBody(cR.Body, response)
	if err != nil {
		return
	}
	res = response
	return
}
