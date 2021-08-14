package main

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}

func assertDeepEqual(t testing.TB, got, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}

func newGetRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

func newDeleteRequest(url string) *http.Request {
	req, _ := http.NewRequest("DELETE", url, nil)
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

func (ts *testStore) FilterTodos(opt TodoOptions) ([]TODO, error) {
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

func (ts *testStore) DeleteTodo(id int) error {
	for i, todo := range ts.todos {
		if todo.ID == id {
			before := ts.todos[:i]
			after := ts.todos[i+1:]
			ts.todos = append(before, after...)
			return nil
		}
	}
	return errors.New("Object not found")
}
