package helper

import (
	"github.com/raihanmd/dependency_injection/model/entity"
	"github.com/raihanmd/dependency_injection/model/web"
)

func ToCategoryResponse(category entity.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []entity.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, web.CategoryResponse{
			Id:   category.Id,
			Name: category.Name,
		})
	}
	return categoryResponses
}
