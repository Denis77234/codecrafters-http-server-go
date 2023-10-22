package server

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
)

var (
	errBadRequest = errors.New("bad request")
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
	req, err := parseRequest(buf)
	if err != nil {
		return Request{}, err
	}
	return req, nil
}

func parseRequest(req []byte) (Request, error) {
	header, body, err := splitHeaderAndBody(req)
	if err != nil {
		return Request{}, err
	}
	headerRows := bytes.Split(header, []byte("\r\n"))

	firstRowBytes := headerRows[0]
	firstRowStr := string(firstRowBytes)
	firstRowContent := strings.Split(firstRowStr, " ")

	rawHeaders := headerRows[0:]
	fmt.Println("HEREHEREHEREHEREHERE")
	for _, rawHeader := range rawHeaders {
		fmt.Println(string(rawHeader))
	}

	headers := parseHeader(rawHeaders)

	request := Request{
		Method: firstRowContent[0],
		URL:    firstRowContent[1],
		Header: headers,
		Body:   body,
	}

	return request, nil
}

func splitHeaderAndBody(rows []byte) (header []byte, body []byte, err error) {

	indOfSeparator := bytes.Index(rows, []byte("\r\n\r\n"))
	if indOfSeparator == -1 {
		return nil, nil, errBadRequest
	}

	header = rows[:indOfSeparator]
	body = rows[indOfSeparator+4:]

	return header, body, nil
}

func parseHeader(headerArr [][]byte) map[string]string {
	headers := make(map[string]string)
	if len(headerArr) == 0 {
		return headers
	}
	for _, h := range headerArr {
		h := string(h)
		header := strings.Split(h, ":")

		headers[header[0]] = strings.Trim(header[1], " ")
	}

	return headers
}
