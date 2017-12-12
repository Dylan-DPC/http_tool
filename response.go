package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Headers http.Header
	Body    io.ReadCloser
}

func (r Response) GetBodyContent() []byte {
	bodyContent, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	return bodyContent
}

func (r Response) BodyContentToStdout() {
	fmt.Print(string(r.GetBodyContent()))
}

func NewResponse(headers http.Header, body io.ReadCloser) Response {
	return Response{
		Headers: headers,
		Body:    body,
	}
}
