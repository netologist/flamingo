package services

import (
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/models"
	"github.com/hasanozgan/vodka/examples/basic/repositories"
)

type UserService struct {
	userRepository    *repositories.UserRepository
	httpClientFactory *vodka.HttpClientFactory
	logger            vodka.Logger
}

func NewUserService(userRepository *repositories.UserRepository, httpClientFactory *vodka.HttpClientFactory, logger vodka.Logger) *UserService {
	return &UserService{userRepository, httpClientFactory, logger}
}

func (s *UserService) GetUserById(id int) (models.User, error) {
	s.logger.Info("Get User By Id %d", id)
	return s.userRepository.FindUserById(id)
}
