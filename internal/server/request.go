package server

type Request struct {
	Method string
	URL    URL
}

type URL struct {
	Path  string
	Value string
}
