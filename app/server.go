package main

import (
	"fmt"

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
		w.WriteStatus(server.STATUS_200_OK)
	})

	err := serv.Start()
	if err != nil {
		fmt.Println(err)
	}

	//l, err := net.Listen("tcp", "0.0.0.0:4221")
	//if err != nil {
	//	fmt.Println("Failed to bind to port 4221")
	//	os.Exit(1)
	//}
	//
	//_, err = l.Accept()
	//if err != nil {
	//	fmt.Println("Error accepting connection: ", err.Error())
	//	os.Exit(1)
	//}

}
