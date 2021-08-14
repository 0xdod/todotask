package main

type Store interface {
	GetTodos() ([]TODO, error)
	GetTodoByID(id int) (*TODO, error)
	CreateTodo(*TODO) error
	FilterTodos(TodoOptions) ([]TODO, error)
	UpdateTodo(id int, opt TodoOptions) (*TODO, error)
	DeleteTodo(id int) error
}
