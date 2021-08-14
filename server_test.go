package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type testStore struct {
	todos []TODO
}

func (ts *testStore) GetTodoByID(id int) (*TODO, error) {
	for _, todo := range ts.todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("Object not found")
}

func (ts *testStore) GetTodos() ([]TODO, error) {
	return ts.todos, nil
}

func TestGetTodos(t *testing.T) {
	testTodos := []TODO{
		{1, "Hello", "hello world"},
		{2, "Hi", "hi world"},
	}
	testStore := &testStore{testTodos}
	server := NewServer(testStore)
	request, _ := http.NewRequest("GET", "/todos", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := []TODO{}

	err := json.NewDecoder(response.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(testStore.todos, got) {
		t.Errorf("Expected %v, got %v", testStore.todos, got)
	}

	assertStatusCode(t, response.Code, http.StatusOK)
}

func TestGetTodoByID(t *testing.T) {
	testTodos := []TODO{
		{1, "Hello", "hello world"},
		{2, "Hi", "hi world"},
	}
	testStore := &testStore{testTodos}
	server := NewServer(testStore)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/todos/%d", testStore.todos[0].ID), nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)
	got := TODO{}

	err := json.NewDecoder(response.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(testStore.todos[0], got) {
		t.Errorf("Expected %v, got %v", testStore.todos[0], got)
	}
	assertStatusCode(t, response.Code, http.StatusOK)
}
