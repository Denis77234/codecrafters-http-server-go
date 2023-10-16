package server

type Request struct {
	Method string
	URL    URL
	Header *map[string]string
}

type URL struct {
	Path  string
	Value string
}
