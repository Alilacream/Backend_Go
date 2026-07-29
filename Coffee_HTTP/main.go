package main

import (
	"fmt"
	"strings"
)

type Request struct {
	path    string
	method  string
	version string
}

func main() {
	requests := []Request{
		{method: "GET", path: "/coffees", version: "HTTP/1.1"},
		{method: "GET", path: "/tea", version: "HTTP/1.1"},
		{method: "POST", path: "/hey", version: "HTTP/1.1"},
		{method: "POST", path: "beans", version: "HTTP/1.1"},
	}

	for _, raw := range requests {
		resp := route(parse(&raw))
		fmt.Println(resp)
	}
}

func route(path, method string) string {
	if path == "" && method == "" {
		return ""
	}

	switch path {
	case "/tea":
		return fmt.Sprintf("Handling the %s for %s", method, path)
	case "/coffees":
		return fmt.Sprintf("Handling the %s for %s", method, path)
	case "/beans":
		return fmt.Sprintf("Handling the %s for %s", method, path)
	default:
		return "404: not found"
	}
}

func parse(req *Request) (string, string) {
	if req == nil {
		fmt.Println("400: bad request")
		return "", ""
	}

	if !strings.HasPrefix(req.version, "HTTP/1.") {
		fmt.Println("505: version not supported")
		return "", ""
	}

	if !strings.HasPrefix(req.path, "/") {
		fmt.Print("400: bad request because the PATH is malformed")
		return "", ""
	}

	return req.path, req.method
}
