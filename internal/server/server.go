package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type status string

const (
	STATUS_200_OK         status = "200 OK"
	STATUS_404_NOTFOUND   status = "404 Not Found"
	STATUS_405_NOTALLOWED status = "405 Method Not Allowed"
	STATUS_201_CREATED    status = "201 Created"
	METHOD_GET            string = "GET"
	METHOD_POST           string = "POST"
)

type (
	Server struct {
		listener net.Listener
		handlers []handler
	}

	ResponseWriter interface {
		WriteStatus(header status)
		WriteContentType(ct string)
		WriteBody(body []byte)
		write(conn net.Conn)
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

func (s *Server) AddHandler(path string, handlerFunc HandlerFunc) {
	h := handler{
		path:           path,
		hadlerFunc:     handlerFunc,
		responseWriter: &rw{},
	}

	s.handlers = append(s.handlers, h)
}

func (s *Server) handle(conn net.Conn) error {
	req, err := s.getRequest(conn)
	if err != nil {
		return err
	}

	handlerExists := false
	for _, h := range s.handlers {
		if strings.HasPrefix(req.URL, h.path) {
			handlerExists = true
			h.hadlerFunc(req, h.responseWriter)
			h.responseWriter.write(conn)
			break
		}
	}

	if !handlerExists {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
	conn.Close()
	return nil
}

func (s *Server) Start() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}

		go s.handle(conn)
	}
	return nil
}
