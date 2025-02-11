package request

type UpdateTagRequest struct {
	ID   int    `validate:"required"`
	Name string `validate:"required,min=1,max=255" json:"name"`
}
