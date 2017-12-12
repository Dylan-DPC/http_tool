package main

type UrlEncodedPayload struct {
	*BasePayload
}

func NewUrlEncodedPayload(data string) *UrlEncodedPayload {
	return &UrlEncodedPayload{
		BasePayload: NewBasePayload(data),
	}
}
