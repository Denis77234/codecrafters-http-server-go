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
	fmt.Println(rows)

	firstRowContent := strings.Split(rows[0], " ")

	url := parseURL(firstRowContent[1])

	rows = rows[1:]

	headers := parseHeader(rows)
	fmt.Println(headers)
	request := Request{
		Method: firstRowContent[0],
		URL:    url,
		Header: headers,
	}

	return request
}

func parseHeader(headerArr []string) map[string]string {
	headers := make(map[string]string)

	for _, h := range headerArr {
		fmt.Println(h)
		if h == "\r\n" {
			continue
		}
		header := strings.Split(h, ":")
		headers[header[0]] = header[1]
	}

	return headers
}

func parseURL(urlStr string) URL {
	url := URL{}

	url.Path = urlStr

	return url
}
