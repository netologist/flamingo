package services

import (
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/repositories"
)

type Services struct {
	dbClientFactory   *vodka.DbClientFactory
	httpClientFactory *vodka.HttpClientFactory
	loggerFactory     *vodka.LoggerFactory
}

func New(dbClientFactory *vodka.DbClientFactory, httpClientFactory *vodka.HttpClientFactory, loggerFactory *vodka.LoggerFactory) *Services {
	return &Services{dbClientFactory, httpClientFactory, loggerFactory}
}

func (s *Services) UserService() *UserService {
	userRepository := repositories.NewUserRepository(s.dbClientFactory, s.loggerFactory.NewLogger("User Repository"))
	return NewUserService(userRepository, s.httpClientFactory, s.loggerFactory.NewLogger("User Service"))
}
