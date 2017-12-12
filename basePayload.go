package main

import (
	"io"
	"strings"
)

type BasePayload struct {
	Data   string
	reader io.Reader
}

func (b BasePayload) GetData() string {
	return b.Data
}

func (b *BasePayload) Read(p []byte) (n int, err error) {
	return b.reader.Read(p)
}

func NewBasePayload(data string) *BasePayload {
	return &BasePayload{
		Data:   data,
		reader: strings.NewReader(data),
	}
}
