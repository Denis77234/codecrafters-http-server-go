package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

func main() {

	dir := flag.String("directory", "", "file directory")
	flag.Parse()

	serv := server.New("tcp", "0.0.0.0:4221")

	serv.AddHandler("/file", func(req server.Request, w server.ResponseWriter) {
		filename := strings.TrimPrefix(req.URL.Path, "/file/")
		path := filepath.Join(*dir, filename)
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			w.WriteStatus(server.STATUS_404_NOTFOUND)
			return
		}
		file, err := os.ReadFile(path)
		if err != nil {
			log.Println(err)
			w.WriteStatus(server.STATUS_404_NOTFOUND)
			return
		}

		w.WriteStatus(server.STATUS_200_OK)
		w.WriteContentType("application/octet-stream")
		w.WriteBody(file)
	})

	serv.AddHandler("/user-agent", func(req server.Request, w server.ResponseWriter) {
		agent := req.Header["User-Agent"]

		value := []byte(agent)

		w.WriteStatus(server.STATUS_200_OK)
		w.WriteContentType("text/plain")
		w.WriteBody(value)
	})

	serv.AddHandler("/echo", func(req server.Request, w server.ResponseWriter) {
		str := strings.TrimPrefix(req.URL.Path, "/echo/")
		value := []byte(str)
		w.WriteStatus(server.STATUS_200_OK)
		w.WriteContentType("text/plain")
		w.WriteBody(value)
	})

	serv.AddHandler("/", func(req server.Request, w server.ResponseWriter) {
		if req.URL.Path != "/" {
			w.WriteStatus(server.STATUS_404_NOTFOUND)
			return
		}
		w.WriteStatus(server.STATUS_200_OK)
	})

	serv.Start()

}
