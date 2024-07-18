package web

type CategoryCreateRequest struct {
	Name string `validate:"required,min=2,max=50" json:"name"`
}
