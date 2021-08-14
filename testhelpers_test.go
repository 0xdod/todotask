package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}

func newGetRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

func newPostRequest(url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest("POST", url, body)
	return req
}

type testStore struct {
	todos []TODO
}

func newTestStore() *testStore {
	testTodos := []TODO{
		{1, "Hello", "hello world"},
		{2, "Hi", "hi world"},
	}
	return &testStore{testTodos}
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

func (ts *testStore) CreateTodo(t *TODO) error {
	ts.todos = append(ts.todos, *t)
	return nil
}

func (ts *testStore) Filter(opt TodoOptions) ([]TODO, error) {
	title := strings.ToLower(opt.Title)
	content := strings.ToLower(opt.Content)
	todos := []TODO{}

	for _, todo := range ts.todos {
		if strings.Contains(todo.Title, title) || strings.Contains(todo.Content, content) {
			todos = append(todos, todo)
		}
	}
	return todos, nil
}
