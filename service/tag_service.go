package service

import (
	"github.com/imvahid/golang-api-sample/data/request"
	"github.com/imvahid/golang-api-sample/data/response"
)

type TagService interface {
	Create(tag request.CreateTagRequest)
	Update(tag request.UpdateTagRequest)
	Delete(tagID int)
	FindByID(tagID int) response.TagResponse
	FindAll() []response.TagResponse
}
