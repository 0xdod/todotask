package main

import (
	"encoding/json"
	"errors"
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

type M map[string]interface{}

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

	JSONResponse(w, http.StatusOK, todos)
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
	JSONResponse(w, http.StatusOK, todo)
}

func (s *Server) createTodo(w http.ResponseWriter, r *http.Request) {
	requestData := TODO{}
	if err := parseJSON(r, &requestData); err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
		return
	}

	if err := s.Store.CreateTodo(&requestData); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}

	JSONResponse(w, http.StatusCreated, requestData)
}

func (s *Server) searchTodoContent(w http.ResponseWriter, r *http.Request) {
	searchTerm := strings.ToLower(r.FormValue("q"))
	todos, err := s.Store.FilterTodos(TodoOptions{searchTerm, searchTerm})
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}

	JSONResponse(w, http.StatusOK, todos)
}

func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := s.Store.DeleteTodo(id); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	JSONResponse(w, http.StatusOK, M{"status": "success",
		"message": "todo with id " + vars["id"] + " deleted successfully"})
}

func (s *Server) updateTodoContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	opt := TodoOptions{}

	if err := parseJSON(r, &opt); err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
		return
	}
	todo, err := s.Store.UpdateTodo(id, opt)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}
	JSONResponse(w, http.StatusOK, todo)
}

func JSONResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(&body)
	if err != nil {
		panic(err)
	}
}

func parseJSON(r *http.Request, d interface{}) error {
	if r.Header.Get("content-type") != "application/json" {
		return errors.New("invalid content-type")
	}
	return json.NewDecoder(r.Body).Decode(d)
}
