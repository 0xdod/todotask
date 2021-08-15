package main

import (
	"fmt"
	"strings"

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

func (ps *PgStore) FilterTodos(opt TodoOptions) ([]TODO, error) {
	todos := []TODO{}

	stmt := "SELECT * FROM todos"

	conditions, args := []string{}, []interface{}{} // not using position 0

	if t := opt.Title; t != "" {
		t = "%" + strings.ToLower(t) + "%"
		conditions, args = append(conditions, "title"), append(args, t)
	}

	if c := opt.Content; c != "" {
		c = "%" + strings.ToLower(c) + "%"
		conditions, args = append(conditions, "content"), append(args, c)
	}

	if l := len(conditions); l > 0 && len(args) > 0 {
		stmt += " WHERE "

		for i, cond := range conditions {
			stmt += fmt.Sprintf(" %s LIKE ? ", cond)
			if i < l-1 {
				stmt += " OR "
			}
		}
	}

	return todos, ps.DB.Raw(stmt, args...).Scan(&todos).Error
}

func (ps *PgStore) UpdateTodo(id int, opt TodoOptions) (*TODO, error) {
	todo := &TODO{ID: id}
	todoUpdates := &TODO{}
	if t := opt.Title; t != "" {
		todoUpdates.Title = t
	}

	if c := opt.Content; c != "" {
		todoUpdates.Content = c
	}

	return todo, ps.DB.Model(todo).Updates(todoUpdates).Error
}

func (ps *PgStore) DeleteTodo(id int) error {
	return ps.DB.Delete(&TODO{}, id).Error
}
