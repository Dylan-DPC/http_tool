package main

import (
	"io"
)

type PayloadInterface interface {
	io.Reader

	GetData() string
}
