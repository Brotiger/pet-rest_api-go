package pgstore_test

import (
	"testing"

	"github.com/Brotiger/rest_api_gopher.git/internal/app/model"
	"github.com/Brotiger/rest_api_gopher.git/internal/app/store"
	"github.com/Brotiger/rest_api_gopher.git/internal/app/store/pgstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := pgstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := pgstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := pgstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := pgstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := pgstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := pgstore.New(db)
	u := model.TestUser(t)

	s.User().Create(u)

	u, err := s.User().Find(u.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u)

}
