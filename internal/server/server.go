package server

import (
	"fmt"
	"net"
	"os"
)

type status string

const (
	STATUS_200_OK       status = "200 OK"
	STATUS_404_NOTFOUND status = "404 Not Found"
)

type (
	ResponseWriter interface {
		WriteStatus(header status)
		WriteContentType(ct string)
		WriteBody(body []byte)
		write(conn net.Conn)
	}

	Server struct {
		listener net.Listener
		handlers []handler
	}

	handler struct {
		path           string
		hadlerFunc     HandlerFunc
		responseWriter ResponseWriter
	}

	HandlerFunc func(req Request, w ResponseWriter)
)

func New(network, address string) Server {
	l, err := net.Listen(network, address)
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	handlers := make([]handler, 0, 10)

	serv := Server{
		listener: l,
		handlers: handlers,
	}
	return serv
}

func (s *Server) Handle(path string, handlerFunc HandlerFunc) {
	h := handler{
		path:           path,
		hadlerFunc:     handlerFunc,
		responseWriter: &rw{},
	}

	s.handlers = append(s.handlers, h)
}

func (s *Server) Start() {
	//conn, err := s.listener.Accept()
	//if err != nil {
	//
	//}
	//defer s.listener.Close()
	//defer conn.Close()
	//
	//_, err = s.getRequest(conn)
	//if err != nil {
	//
	//}
	//
	//conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	//for _, h := range s.handlers {
	//	if h.path == req.URL.Path {
	//		h.hadlerFunc(req, h.responseWriter)
	//		h.responseWriter.write(conn)
	//	}
	//}

	defer s.listener.Close()

	conn, err := s.listener.Accept()
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

	fmt.Println(string(buffer))
	
	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	if err != nil {
		fmt.Printf("my programm write:%v\n", err)
	}
}
