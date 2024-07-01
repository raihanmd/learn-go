package controller

import (
	"net/http"
	"strconv"

	"github.com/raihanmd/dependency_injection/helper"
	"github.com/raihanmd/dependency_injection/model/web"
	"github.com/raihanmd/dependency_injection/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateReq := web.CategoryCreateRequest{}

	helper.ReadFromReqBody(r, &categoryCreateReq)

	categoryResponse := c.CategoryService.Create(r.Context(), categoryCreateReq)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    categoryResponse,
	}

	helper.WriteToResBody(w, webResponse)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateReq := web.CategoryUpdateRequest{}

	helper.ReadFromReqBody(r, &categoryUpdateReq)

	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryUpdateReq.Id = id

	categoryResponse := c.CategoryService.Update(r.Context(), categoryUpdateReq)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    categoryResponse,
	}

	helper.WriteToResBody(w, webResponse)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	c.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
	}

	helper.WriteToResBody(w, webResponse)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryResponse := c.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    categoryResponse,
	}

	helper.WriteToResBody(w, webResponse)
}

func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := c.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    categoryResponses,
	}

	helper.WriteToResBody(w, webResponse)
}
