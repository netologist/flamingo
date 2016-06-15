package services

import (
	"github.com/hasanozgan/flamingo"
	"github.com/hasanozgan/flamingo/examples/basic/repositories"
)

type Services struct {
	dbClientFactory   *flamingo.DbClientFactory
	httpClientFactory *flamingo.HttpClientFactory
	loggerFactory     *flamingo.LoggerFactory
}

func New(dbClientFactory *flamingo.DbClientFactory, httpClientFactory *flamingo.HttpClientFactory, loggerFactory *flamingo.LoggerFactory) *Services {
	return &Services{dbClientFactory, httpClientFactory, loggerFactory}
}

func (s *Services) UserService() *UserService {
	userRepository := repositories.NewUserRepository(s.dbClientFactory, s.loggerFactory.NewLogger("User Repository"))
	return NewUserService(userRepository, s.httpClientFactory, s.loggerFactory.NewLogger("User Service"))
}
