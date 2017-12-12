package main

import (
	"net/http"
)

type PostRequest struct {
	*BaseRequest

	ContentType string
}

func (p *PostRequest) Send() {
	request, err := http.NewRequest(http.MethodPost, p.Url, p.Payload)

	request.Header.Set("content-type", p.ContentType)
	request.ContentLength = int64(len(p.Payload.GetData()))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	p.Response = NewResponse(response.Header, response.Body)
}

func NewPostRequest(url string, contentType string, payload PayloadInterface) *PostRequest {
	post := PostRequest{
		BaseRequest: NewBaseRequest(http.MethodPost, url, payload),
		ContentType: contentType,
	}

	return &post
}
