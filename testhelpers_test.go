package main

import (
	"io"
	"net/http"
	"reflect"
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
	return newHTTPRequest("GET", url, nil)
}
func newDeleteRequest(url string) *http.Request {
	return newHTTPRequest("DELETE", url, nil)
}

func newPutRequest(url string, body io.Reader) *http.Request {
	return newHTTPRequest("PUT", url, body)
}

func newPostRequest(url string, body io.Reader) *http.Request {
	return newHTTPRequest("POST", url, body)
}

func newHTTPRequest(method, url string, body io.Reader) *http.Request {
	r, _ := http.NewRequest(method, url, body)
	r.Header.Set("content-type", "application/json")
	return r
}

func newTestStore() *InMemStore {
	testTodos := []TODO{}
	m := map[string]string{
		"Hello": "hello world",
		"Hi":    "hi world",
	}
	for k, v := range m {
		testTodos = append(testTodos, *NewTodo(k, v))
	}
	return &InMemStore{testTodos}
}
