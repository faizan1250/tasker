package repository

import "github.com/faizan1250/tasker/internal/server"

type Repositories struct {
	Todo *TodoRepository
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{
		Todo: NewTodoRepository(s),
	}
}
