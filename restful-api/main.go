package main

import (
	"net/http"
	"restful_api/app"
	"restful_api/controller"
	"restful_api/helper"
	"restful_api/middleware"
	"restful_api/repository"
	"restful_api/service"

	"github.com/go-playground/validator/v10"
)

func main() {

	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	err := http.ListenAndServe(":3000", middleware.NewAuthMiddleware(router))
	helper.PanicIfError(err)
}
