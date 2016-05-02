package repositories

import (
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/models"
)

type UserRepository struct {
	dbClientFactory *vodka.DbClientFactory
	loggerFactory   *vodka.LoggerFactory
}

func NewUserRepository(dbClientFactory *vodka.DbClientFactory, loggerFactory *vodka.LoggerFactory) *UserRepository {
	return &UserRepository{dbClientFactory, loggerFactory}
}

func (r *UserRepository) FindUserById(id int) models.User {
	return models.User{1234, "Ekin Ozgan", "Istanbul"}
}
