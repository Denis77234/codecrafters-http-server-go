package main

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

var responseOk []byte = []byte("HTTP/1.1 200 OK\r\n\r\n")

func main() {

	s := server.New("tcp", "0.0.0.0:4221")

	s.Start()

}
