package services

import (
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/models"
	"github.com/hasanozgan/vodka/examples/basic/repositories"
)

type UserService struct {
	userRepository    *repositories.UserRepository
	httpClientFactory *vodka.HttpClientFactory
	loggerFactory     *vodka.LoggerFactory
}

func NewUserService(userRepository *repositories.UserRepository, httpClientFactory *vodka.HttpClientFactory, loggerFactory *vodka.LoggerFactory) *UserService {
	return &UserService{userRepository, httpClientFactory, loggerFactory}
}

func (s *UserService) GetUserById(id int) models.User {
	return s.userRepository.FindUserById(id)
}
