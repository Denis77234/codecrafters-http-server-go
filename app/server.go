package main

import (
	"fmt"
	"net"
	"os"
)

var responseOk []byte = []byte("HTTP/1.1 200 OK\r\n\r\n")

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Printf("my programm read:%v\n", err)
	}

	_, err = conn.Write(responseOk)
	if err != nil {
		fmt.Printf("my programm write:%v\n", err)
	}

}
