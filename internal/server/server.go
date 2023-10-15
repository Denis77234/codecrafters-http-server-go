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

func (s *Server) Start() error {
	conn, err := s.listener.Accept()
	if err != nil {
		return err
	}

	req, err := s.getRequest(conn)
	if err != nil {
		return err
	}

	for _, h := range s.handlers {
		if h.path == req.URL.Path {
			h.hadlerFunc(req, h.responseWriter)
		}
	}
	return nil
}