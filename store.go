package main

import "errors"

type Store interface {
	GetTodos() ([]TODO, error)
	GetTodoByID(id int) (*TODO, error)
	CreateTodo(*TODO) error
	FilterTodos(TodoOptions) ([]TODO, error)
	DeleteTodo(id int) error
}

type InMemStore struct {
	todos []TODO
}

type PsqlStore struct {
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

func (im *InMemStore) CreateTodo(t *TODO) error {
	return nil
}

func (im *InMemStore) FilterTodos(opt TodoOptions) ([]TODO, error) {
	return nil, nil
}

func (im *InMemStore) DeleteTodo(id int) error {
	for i, todo := range im.todos {
		if todo.ID == id {
			before := im.todos[:i]
			after := im.todos[i+1:]
			im.todos = append(before, after...)
			return nil
		}
	}
	return errors.New("Object not found")
}
