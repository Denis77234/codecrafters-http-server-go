package main

import (
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

func main() {

	serv := server.New("tcp", "0.0.0.0:4221")

	serv.Handle("/user-agent", func(req server.Request, w server.ResponseWriter) {
		agent := req.Header["User-Agent"]

		value := []byte(agent)

		w.WriteStatus(server.STATUS_200_OK)
		w.WriteContentType("text/plain")
		w.WriteBody(value)
	})

	serv.Handle("/echo", func(req server.Request, w server.ResponseWriter) {
		str := strings.TrimPrefix(req.URL.Path, "/echo/")
		value := []byte(str)
		w.WriteStatus(server.STATUS_200_OK)
		w.WriteContentType("text/plain")
		w.WriteBody(value)
	})

	serv.Handle("/", func(req server.Request, w server.ResponseWriter) {
		if req.URL.Path != "/" {
			w.WriteStatus(server.STATUS_404_NOTFOUND)
			return
		}
		w.WriteStatus(server.STATUS_200_OK)
	})

	serv.Start()

}
