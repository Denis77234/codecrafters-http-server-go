package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

func main() {

	dir := flag.String("directory", "", "file directory")
	flag.Parse()

	serv := server.New("tcp", "0.0.0.0:4221")

	serv.AddHandler("/files", func(req server.Request, w server.ResponseWriter) {
		if req.Method == server.METHOD_GET {
			filename := strings.TrimPrefix(req.URL.Path, "/files/")
			path := filepath.Join(*dir, filename)
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				w.WriteStatus(server.STATUS_404_NOTFOUND)
				return
			}
			file, err := os.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				w.WriteStatus(server.STATUS_404_NOTFOUND)
				return
			}

			w.WriteStatus(server.STATUS_200_OK)
			w.WriteContentType("application/octet-stream")
			w.WriteBody(file)
		}
		if req.Method == server.METHOD_POST {

			filename := strings.TrimPrefix(req.URL.Path, "/files/")
			path := filepath.Join(*dir, filename)

		}
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
