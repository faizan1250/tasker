package service

import (
	"github.com/faizan1250/tasker/internal/lib/job"
	"github.com/faizan1250/tasker/internal/repository"
	"github.com/faizan1250/tasker/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
