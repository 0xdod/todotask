package main

import "errors"

type Store interface {
	GetTodos() ([]TODO, error)
	GetTodoByID(id int) (*TODO, error)
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
