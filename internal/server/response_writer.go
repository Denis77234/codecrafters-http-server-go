package server

import (
	"fmt"
)

type rw struct {
	header      []byte
	contentType []byte
	body        []byte
}

func (w *rw) WriteStatus(header status) {
	str := fmt.Sprintf("HTTP/1.1 %v\r\n", header)
	w.header = []byte(str)
}

func (w *rw) WriteContentType(ct string) {
	str := fmt.Sprintf("Content-Type: %v\r\n", ct)
	w.contentType = []byte(str)
}

func (w *rw) WriteBody(body []byte) {
	w.body = body
}
