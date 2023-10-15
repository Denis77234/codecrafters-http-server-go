package main

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

//
//var responseOk []byte = []byte("HTTP/1.1 200 OK\r\n\r\n")
//
//var responseNotOk []byte = []byte("HTTP/1.1 404 Not Found\r\n\r\n")
//
//func getPath(buf []byte) string {
//	str := string(buf)
//	rows := strings.Split(str, "\r\n")
//
//	firstRowContent := strings.Split(rows[0], " ")
//
//	path := firstRowContent[1]
//
//	return path
//}
//
//func getValue(path string) string {
//	content := strings.Split(path, "/")
//
//	value := content[1]
//
//	return value
//}
//
//func makeBody(value string) []byte {
//
//	contentLength := len(value)
//
//	bodyStr := fmt.Sprintf("Content-Type: text/plain\r\nContent-Length: %v\r\n\r\n%s\r\n", contentLength, value)
//
//	body := []byte(bodyStr)
//
//	return body
//}

func main() {

	serv := server.New("tcp", "0.0.0.0:4221")

	serv.Handle("/", func(req server.Request, w server.ResponseWriter) {
		if req.URL.Path != "/" {
			w.WriteStatus(server.STATUS_404_NOTFOUND)
			return
		}
		w.WriteStatus(server.STATUS_200_OK)
	})

	serv.Handle("/echo", func(req server.Request, w server.ResponseWriter) {
		value := []byte(req.URL.Value)
		w.WriteStatus(server.STATUS_200_OK)
		w.WriteContentType("text/plain")
		w.WriteBody(value)
	})
	serv.Start()

}
