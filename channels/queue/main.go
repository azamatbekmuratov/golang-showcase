package main

import "net/http"

const MAX_QUEUE = 5

type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

type Payload struct {
	// [redacted]
}

type Content []Payload

var Queue chan Payload

func (p *Payload) UploadToS3() error {
	return nil
}

func init() {
	Queue = make(chan Payload, MAX_QUEUE)
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	payloads := Content{}
	for _, payload := range payloads {
		Queue <- payload
	}
}
