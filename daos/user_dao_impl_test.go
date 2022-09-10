package daos_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/arikama/koran-backend/daos"
	"github.com/arikama/koran-backend/models"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	container, err := mysqltestcontainer.Create("test")
	assert.Nil(t, err)

	db := container.GetDb()

	err = arctictern.Migrate(db, "./../migrations")
	assert.Nil(t, err)

	var userDao daos.UserDao
	userDao, err = daos.NewUserDaoImpl(db)
	assert.Nil(t, err)

	faker := faker.New()

	email := faker.Internet().Email()
	name := faker.Person().Name()
	token := faker.Internet().Password()
	picture := faker.Internet().URL()

	newUser := models.User{
		Email:   email,
		Name:    name,
		Token:   token,
		Picture: picture,
	}

	err = userDao.CreateUser(newUser.Email, newUser.Token)
	assert.Nil(t, err)

	queriedUser, err := userDao.QueryUserByToken(token)
	assert.Nil(t, err)
	assert.Equal(t, newUser.Email, queriedUser.Email)
}
