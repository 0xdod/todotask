package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgStore struct {
	*gorm.DB
}

func NewPgStore(dsn string) (Store, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PgStore{db}, nil

}

func (ps *PgStore) GetTodos() ([]TODO, error) {
	todos := []TODO{}
	return todos, ps.DB.Find(&todos).Error
}

func (ps *PgStore) GetTodoByID(id int) (*TODO, error) {
	todo := &TODO{}
	return todo, ps.DB.First(todo, id).Error

}

func (ps *PgStore) CreateTodo(todo *TODO) error {
	return ps.DB.Create(todo).Error
}

func (ps *PgStore) FilterTodos(TodoOptions) ([]TODO, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *PgStore) UpdateTodo(id int, opt TodoOptions) (*TODO, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *PgStore) DeleteTodo(id int) error {
	panic("not implemented") // TODO: Implement
}
