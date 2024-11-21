package services

import (
	"context"
	"sync"

	"github.com/KhachikAstoyan/go-api-template/db/repository"
)

type UserService interface {
	CountUsers() (int64, error)
}

type userService struct {
	repo *repository.Queries
}

var (
	service *userService
	once    sync.Once
)

func GetUserService(repo *repository.Queries) UserService {
	once.Do(func() {
		service = &userService{
			repo: repo,
		}
	})

	return service
}

func (s *userService) CountUsers() (int64, error) {
	count, err := s.repo.CountActiveUsers(context.Background())

	if err != nil {
		return 0, err
	}

	return count, nil
}
