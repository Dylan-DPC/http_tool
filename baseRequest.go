package main

import (
	"strings"
)

type BaseRequest struct {
	Url      string
	Method   string
	Response Response
	Payload  PayloadInterface
}

func (b *BaseRequest) ParseUrl() {
	url := strings.Trim(b.Url, "/")

	if strings.Index(url, "http://") == -1 && strings.Index(url, "https://") == -1 {
		url = "http://" + url
	}

	b.Url = url
}

func NewBaseRequest(method string, url string, payload PayloadInterface) *BaseRequest {
	base := BaseRequest{
		Url:     url,
		Method:  method,
		Payload: payload,
	}

	base.ParseUrl()

	return &base
}
