package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

type TodoOptions struct {
	Title   string
	Content string
}

func (t *TODO) ToJSON() string {
	data, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func NewServer(store Store) *Server {
	s := &Server{Store: store}
	router := mux.NewRouter()
	router.HandleFunc("/todos", s.getTodoList).Methods("GET")
	router.HandleFunc("/todos", s.createTodo).Methods("POST")
	router.HandleFunc("/todos/search", s.searchTodoContent).Methods("GET")
	router.HandleFunc("/todos/{id}", s.getTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", s.updateTodoContent).Methods("PUT")
	router.HandleFunc("/todos/{id}", s.deleteTodo).Methods("DELETE")
	s.Handler = router
	return s
}

func (s *Server) getTodoList(w http.ResponseWriter, r *http.Request) {
	todos, err := s.Store.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
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
		return
	}

	if err := json.NewEncoder(w).Encode(&todo); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
	}
}

func (s *Server) createTodo(w http.ResponseWriter, r *http.Request) {
	requestData := TODO{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
		return
	}

	if err := s.Store.CreateTodo(&requestData); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, requestData.ToJSON())
}

func (s *Server) searchTodoContent(w http.ResponseWriter, r *http.Request) {
	searchTerm := strings.ToLower(r.FormValue("q"))
	todos, err := s.Store.Filter(TodoOptions{searchTerm, searchTerm})
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(&todos); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
	}

}

func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) updateTodoContent(w http.ResponseWriter, r *http.Request) {
}
