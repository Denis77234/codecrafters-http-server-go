package server

import (
	"fmt"
	"net"
	"strings"
)

func (s *Server) getRequest(conn net.Conn) (Request, error) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		return Request{}, err
	}

	if len(buf) == 0 {
		return Request{}, nil
	}
	req := parseRequest(buf)
	return req, nil
}

func parseRequest(req []byte) Request {
	str := string(req)

	rows := strings.Split(str, "\r\n")

	firstRowContent := strings.Split(rows[0], " ")

	url := parseURL(firstRowContent[1])

	request := Request{
		Method: firstRowContent[0],
		URL:    url,
	}

	return request
}

func parseURL(urlStr string) URL {
	url := URL{}

	url.Path = urlStr
	fmt.Println(urlStr)

	return url
}
