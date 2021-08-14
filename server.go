package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Server struct {
	http.Handler
	Store
}

type TODO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewServer(store Store) *Server {
	s := &Server{Store: store}
	router := mux.NewRouter()
	router.HandleFunc("/todos", s.getTodoList).Method("GET")
	router.HandleFunc("/todos", s.createTodo).Method("POST")
	router.HandleFunc("/todos/{id}", s.deleteTodo).Method("DELETE")
	router.HandleFunc("/todos/{id}/search", s.searchTodoContent)
	router.HandleFunc("/todos/{id}", s.getTodo).Method("GET")
	router.HandleFunc("/todos/{id}", s.updateTodoContent).Method("PUT")
	s.Handler = router
	return s
}

func (s *Server) getTodoList(w http.ResponseWriter, r *http.Request) {
	todos, err := s.Store.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
	}

	if err := json.NewEncoder(w).Encode(&todos); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
	}
}

func (s *Server) getTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	todo, err := s.Store.GetTodoByID(id)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
	}

	if err := json.NewEncoder(w).Encode(&todo); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
	}
}

func (s *Server) createTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) searchTodoContent(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) updateTodoContent(w http.ResponseWriter, r *http.Request) {
}
