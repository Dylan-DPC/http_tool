package main

import (
	"net/http"
	"strings"
)

type GetRequest struct {
	*BaseRequest
}

func (g *GetRequest) Send() {
	url := g.Url

	if strings.Index(url, "?") > -1 {
		url += "&" + g.Payload.GetData()
	} else {
		url += "?" + g.Payload.GetData()
	}

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	g.Response = NewResponse(response.Header, response.Body)
}

func NewGetRequest(url string, payload PayloadInterface) *GetRequest {
	get := GetRequest{
		BaseRequest: NewBaseRequest(http.MethodGet, url, payload),
	}

	return &get
}
