package server

import (
	"fmt"
	"net"
	"strings"
)

func (s *Server) getRequest(conn net.Conn) (Request, error) {
	buf := make([]byte, 0, 10)
	_, err := conn.Read(buf)
	if err != nil {
		return Request{}, err
	}

	req := parseRequest(buf)

	return req, nil
}

func parseRequest(req []byte) Request {
	str := string(req)

	rows := strings.Split(str, "\r\n")

	fmt.Println(str)

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

	urlValues := strings.Split(urlStr, "/")
	valuesLen := len(urlValues)

	if valuesLen == 2 {
		pathStr := strings.Join(urlValues, "/")
		url.Path = pathStr
		return url
	}

	path := urlValues[:valuesLen-1]
	pathStr := strings.Join(path, "/")

	valueIndex := valuesLen - (valuesLen - 2)
	values := urlValues[valueIndex:]
	url.Path = pathStr
	url.Value = values[0]

	return url
}
