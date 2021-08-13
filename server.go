package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	*http.ServeMux
}

type TODO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewServer() *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", todoListHandler)
	return &Server{mux}
}

func todoListHandler(w http.ResponseWriter, r *http.Request) {
	data := `
[
{
  "title": "Hello",
  "content": "hello world"
},
{
  "title": "Hi",
  "content": "hi world"
}
]
`
	fmt.Fprint(w, data)
}
