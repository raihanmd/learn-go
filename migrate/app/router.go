package app

import (
	"github.com/raihanmd/dependency_injection/controller"
	"github.com/raihanmd/dependency_injection/exeption"
	"github.com/raihanmd/dependency_injection/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *middleware.AuthMiddleware {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exeption.ErrorHandler

	return middleware.NewAuthMiddleware(router)
}
