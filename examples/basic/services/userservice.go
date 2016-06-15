package services

import (
	"github.com/hasanozgan/flamingo"
	"github.com/hasanozgan/flamingo/examples/basic/models"
	"github.com/hasanozgan/flamingo/examples/basic/repositories"
)

type UserService struct {
	userRepository    *repositories.UserRepository
	httpClientFactory *flamingo.HttpClientFactory
	logger            flamingo.Logger
}

func NewUserService(userRepository *repositories.UserRepository, httpClientFactory *flamingo.HttpClientFactory, logger flamingo.Logger) *UserService {
	return &UserService{userRepository, httpClientFactory, logger}
}

func (s *UserService) GetUserById(id int) (models.User, error) {
	s.logger.Info("Get User By Id %d", id)
	return s.userRepository.FindUserById(id)
}
