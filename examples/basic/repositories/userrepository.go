package repositories

import (
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/models"
)

type UserRepository struct {
	dbClientFactory *vodka.DbClientFactory
	logger          vodka.Logger
}

func NewUserRepository(dbClientFactory *vodka.DbClientFactory, logger vodka.Logger) *UserRepository {
	return &UserRepository{dbClientFactory, logger}
}

func (r *UserRepository) FindUserById(id int) models.User {
	r.logger.Info("Find User By Id %d", id)
	return models.User{1234, "Ekin Ozgan", "Istanbul"}
}
