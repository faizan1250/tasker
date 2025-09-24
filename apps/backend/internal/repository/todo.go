package repository

import (
	"context"

	"github.com/faizan1250/tasker/internal/model/todo"
	"github.com/faizan1250/tasker/internal/server"
)

type TodoRepository struct {
	Server *server.Server
}

func NewTodoRepository(server *server.Server) *TodoRepository {
	return &TodoRepository{
		Server: server,
	}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, userID string, payload *todo.CreateTodoPayload) *todo.Todo {
	return nil
}
