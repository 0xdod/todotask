package main

import (
	"errors"
	"strings"
)

type InMemStore struct {
	todos []TODO
}

func NewInMemStore() *InMemStore {
	return new(InMemStore)
}

func (im *InMemStore) GetTodos() ([]TODO, error) {
	return im.todos, nil
}

func (im *InMemStore) GetTodoByID(id int) (*TODO, error) {
	for _, todo := range im.todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("Object not found")
}

func (im *InMemStore) CreateTodo(todo *TODO) error {
	if len(im.todos) == 0 {
		todo.ID = 1
	} else {
		todo.ID = im.todos[len(im.todos)-1].ID + 1
	}
	im.todos = append(im.todos, *todo)
	return nil
}

func (im *InMemStore) FilterTodos(opt TodoOptions) ([]TODO, error) {

	title := strings.ToLower(opt.Title)
	content := strings.ToLower(opt.Content)
	todos := []TODO{}

	for _, todo := range im.todos {
		if strings.Contains(todo.Title, title) || strings.Contains(todo.Content, content) {
			todos = append(todos, todo)
		}
	}
	return todos, nil
}

func (im *InMemStore) DeleteTodo(id int) error {
	for i, todo := range im.todos {
		if todo.ID == id {
			im.todos = append(im.todos[:i], im.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Object not found")
}

func (im *InMemStore) UpdateTodo(id int, opt TodoOptions) (*TODO, error) {
	todo, err := im.GetTodoByID(id)
	if err != nil {
		return nil, err
	}

	if title := opt.Title; title != "" {
		todo.Title = title
	}
	if content := opt.Content; content != "" {
		todo.Content = content
	}
	im.todos[0] = *todo
	return todo, nil
}
