package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
	assertDeepEqual(t, got, testStore.todos)
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
	assertDeepEqual(t, got, testStore.todos[0])
}

func TestCreateTodos(t *testing.T) {
	testStore := newTestStore()

	todo := `
{
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

func TestDeleteTodo(t *testing.T) {
	testStore := newTestStore()
	server := NewServer(testStore)
	request := newDeleteRequest(fmt.Sprintf("/todos/%d", testStore.todos[0].ID))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := M{}

	err := json.NewDecoder(response.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	assertStatusCode(t, response.Code, http.StatusOK)

	if status, ok := got["status"].(string); ok && status != "success" {
		t.Errorf("Expected success status, but got %s", status)
	}

	if tl := len(testStore.todos); tl != 1 {
		t.Errorf("Expected length of 1, but got %d", tl)
	}
}

func TestUpdateTodo(t *testing.T) {
	todoUpdate := `
{
    "title": "Change",
	"content": "Change is constant"
}
`
	testStore := newTestStore()
	server := NewServer(testStore)
	request := newPutRequest(fmt.Sprintf("/todos/%d", testStore.todos[0].ID), strings.NewReader(todoUpdate))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := TODO{}

	err := json.NewDecoder(response.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	assertStatusCode(t, response.Code, http.StatusOK)
	if title := got.Title; title != "Change" {
		t.Errorf("Expected %s, but got %s", "Change", title)
	}
	assertDeepEqual(t, got, testStore.todos[0])
}
