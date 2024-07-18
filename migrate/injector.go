//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/raihanmd/dependency_injection/app"
	"github.com/raihanmd/dependency_injection/controller"
	"github.com/raihanmd/dependency_injection/middleware"
	"github.com/raihanmd/dependency_injection/repository"
	"github.com/raihanmd/dependency_injection/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

func InitializedRouter() *middleware.AuthMiddleware {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Value([]validator.Option{}),
	)
	return nil
}
