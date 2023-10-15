package server

import (
	"fmt"
	"net"
)

type rw struct {
	header        []byte
	contentType   []byte
	contentLength []byte
	body          []byte
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
	length := len(body)
	cl := fmt.Sprintf("Content-Length: %v\r\n", length)
	w.contentLength = []byte(cl)

	row := []byte("\r\n")
	w.body = append(row, body...)
}

func (w *rw) makeResponse() []byte {
	response := append(w.header, w.contentType...)
	response = append(response, w.contentLength...)
	response = append(response, w.body...)
	response = append(response, []byte("\r\n")...)
	return response
}

func (w rw) write(conn net.Conn) {
	conn.Write(w.makeResponse())
}
