package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestGetTodos(t *testing.T) {
	testStore := newTestStore()
	server := NewServer(testStore)
	request := newGetRequest("/todos")
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := []TODO{}

	err := json.NewDecoder(response.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	assertStatusCode(t, response.Code, http.StatusOK)
	assertDeepEqual(t, testStore.todos, got)
}

func TestGetTodoByID(t *testing.T) {
	testStore := newTestStore()
	server := NewServer(testStore)
	request := newGetRequest(fmt.Sprintf("/todos/%d", testStore.todos[0].ID))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := TODO{}

	err := json.NewDecoder(response.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	assertStatusCode(t, response.Code, http.StatusOK)
	assertDeepEqual(t, testStore.todos[0], got)
}

func assertDeepEqual(t testing.TB, got, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}

func TestCreateTodos(t *testing.T) {
	testStore := newTestStore()

	todo := `
{
    "id": 3,
	"title": "Money",
	"content": "kudi"
}
`
	request := newPostRequest("/todos", strings.NewReader(todo))
	response := httptest.NewRecorder()
	server := NewServer(testStore)
	server.ServeHTTP(response, request)
	got := TODO{}
	err := json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Fatalf("Could not parse json: %v", err)
	}

	assertStatusCode(t, response.Code, http.StatusCreated)
	if len(testStore.todos) != 3 {
		t.Errorf("Expected list length of 3, but got %d", len(testStore.todos))
	}
	assertDeepEqual(t, got, testStore.todos[2])
}

func TestSearchTodoContent(t *testing.T) {
	testStore := newTestStore()
	server := NewServer(testStore)
	request := newGetRequest("/todos/search?q=hello")
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := []TODO{}

	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	assertStatusCode(t, response.Code, http.StatusOK)
	assertDeepEqual(t, got, []TODO{testStore.todos[0]})
}
