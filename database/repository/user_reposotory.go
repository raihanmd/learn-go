package repository

import (
	"context"
	"golang_database/model"
)

type UserRepository interface {
	Insert(ctx context.Context, user model.User) (model.User, error)
	FindById(ctx context.Context, id string) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
}
