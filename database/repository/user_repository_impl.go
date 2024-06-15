package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang_database/model"
)

type userRepoImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepoImpl{db}
}

func (r *userRepoImpl) Insert(ctx context.Context, user model.User) (model.User, error) {
	sql := "INSERT INTO users (id, name, password, email, balance, rating, birth_date, married) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := r.DB.ExecContext(ctx, sql, user.Id, user.Name, user.Password, user.Email, user.Balance, user.Rating, user.BirthDate, user.Married)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (r *userRepoImpl) FindById(ctx context.Context, id string) (model.User, error) {
	sql := "SELECT id, name, password, email, balance, rating, birth_date, created_at, married FROM users WHERE id=? LIMIT 1"
	row, err := r.DB.QueryContext(ctx, sql, id)
	user := model.User{}
	if err != nil {
		return user, err
	}
	defer row.Close()

	if row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Balance, &user.Rating, &user.BirthDate, &user.Created, &user.Married)
		return user, nil
	} else {
		return user, fmt.Errorf("user with id %s not found", id)
	}
}
func (r *userRepoImpl) FindAll(ctx context.Context) ([]model.User, error) {
	sql := "SELECT id, name, password, email, balance, rating, birth_date, created_at, married FROM users"
	row, err := r.DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var users []model.User
	for row.Next() {
		user := model.User{}
		row.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Balance, &user.Rating, &user.BirthDate, &user.Created, &user.Married)
		users = append(users, user)
	}
	return users, nil
}
