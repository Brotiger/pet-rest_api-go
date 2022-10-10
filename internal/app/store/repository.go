package store

import "github.com/Brotiger/rest_api_gopher.git/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
