package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETTodoList(t *testing.T) {
	server := NewServer()
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	expected := []TODO{
		{"Hello", "hello world"},
		{"Hi", "hi world"},
	}

	got := []TODO{}

	err := json.NewDecoder(response.Body).Decode(&got)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	assertStatusCode(t, response.Code, http.StatusOK)

}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}
