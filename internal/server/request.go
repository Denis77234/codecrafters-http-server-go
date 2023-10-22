package server

type Request struct {
	Method string
	URL    string
	Header map[string]string
	Body   []byte
}
