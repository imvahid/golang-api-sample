package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/imvahid/golang-api-sample/data/request"
	"github.com/imvahid/golang-api-sample/data/response"
	"github.com/imvahid/golang-api-sample/helper"
	"github.com/imvahid/golang-api-sample/model"
	"github.com/imvahid/golang-api-sample/repository"
)

type TagServiceImplementation struct {
	TagRepository repository.TagRepository
	Validate      *validator.Validate
}

func NewTagServiceImplementation(tagRepository repository.TagRepository, validate *validator.Validate) TagService {
	return &TagServiceImplementation{
		TagRepository: tagRepository,
		Validate:      validate,
	}
}

func (t *TagServiceImplementation) Create(tag request.CreateTagRequest) {
	err := t.Validate.Struct(tag)
	helper.ErrorPanic(err)

	tagModel := model.Tag{
		Name: tag.Name,
	}
	t.TagRepository.Create(tagModel)
}

func (t *TagServiceImplementation) Update(tag request.UpdateTagRequest) {
	result, err := t.TagRepository.FindByID(tag.ID)
	helper.ErrorPanic(err)

	result.Name = tag.Name
	t.TagRepository.Update(result)
}

func (t *TagServiceImplementation) Delete(tagID int) {
	t.TagRepository.Delete(tagID)
}

func (t *TagServiceImplementation) FindByID(tagID int) response.TagResponse {
	result, err := t.TagRepository.FindByID(tagID)
	helper.ErrorPanic(err)

	tag := response.TagResponse{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
	}
	return tag
}

func (t *TagServiceImplementation) FindAll() []response.TagResponse {
	result := t.TagRepository.FindAll()

	var tags []response.TagResponse

	for _, value := range result {
		tag := response.TagResponse{
			ID:        value.ID,
			Name:      value.Name,
			CreatedAt: value.CreatedAt,
		}
		tags = append(tags, tag)
	}

	return tags
}
