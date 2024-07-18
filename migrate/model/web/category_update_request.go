package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=2,max=50" json:"name"`
}
