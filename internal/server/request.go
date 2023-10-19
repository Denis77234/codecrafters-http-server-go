package server

type Request struct {
	Method string
	URL    URL
	Header map[string]string
	Body   []byte
}

type URL struct {
	Path  string
	Value string
}
