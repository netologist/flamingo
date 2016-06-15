package repositories

import (
	"github.com/hasanozgan/flamingo"
	"github.com/hasanozgan/flamingo/examples/basic/models"
)

type UserRepository struct {
	dbClientFactory *flamingo.DbClientFactory
	logger          flamingo.Logger
}

func NewUserRepository(dbClientFactory *flamingo.DbClientFactory, logger flamingo.Logger) *UserRepository {
	return &UserRepository{dbClientFactory, logger}
}

func (r *UserRepository) FindUserById(id int) (models.User, error) {
	result := models.User{}
	results := []models.User{}

	session, err := r.dbClientFactory.MongoClient.OpenSession()

	if err != nil {
		return result, err
	}

	session.Db("mydb").Collection("users").FindById(123).One(&result)
	session.Db("mydb").Collection("users").FindById(123).All(&results)
	session.Db("mydb").Collection("users").Query(flamingo.Selectors{"id": 123}).All(&results)

	defer session.Close()

	r.logger.Info("Find User By Id %d", id)
	return models.User{1234, "Ekin Ozgan", "Istanbul"}, nil
}
