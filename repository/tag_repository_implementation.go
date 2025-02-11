package repository

import (
	"errors"
	"github.com/imvahid/golang-api-sample/data/request"
	"github.com/imvahid/golang-api-sample/helper"
	"github.com/imvahid/golang-api-sample/model"
	"gorm.io/gorm"
)

type TagRepositoryImplementation struct {
	DB *gorm.DB
}

func NewTagRepositoryImplementation(DB *gorm.DB) TagRepository {
	return &TagRepositoryImplementation{DB: DB}
}

func (t *TagRepositoryImplementation) Create(tag model.Tag) {
	result := t.DB.Create(&tag)
	helper.ErrorPanic(result.Error)
}

func (t *TagRepositoryImplementation) Update(tag model.Tag) {
	var updateTag = request.UpdateTagRequest{
		ID:   tag.ID,
		Name: tag.Name,
	}

	result := t.DB.Model(&tag).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func (t *TagRepositoryImplementation) Delete(tagID int) {
	var tag model.Tag
	result := t.DB.Where("id = ?", tagID).Delete(&tag)
	helper.ErrorPanic(result.Error)
}

func (t *TagRepositoryImplementation) FindByID(tagID int) (tags model.Tag, err error) {
	var tag model.Tag
	result := t.DB.Find(&tag, tagID)

	if result != nil {
		return tag, nil
	}
	return tag, errors.New("tag is not found")
}

func (t *TagRepositoryImplementation) FindAll() []model.Tag {
	var tags []model.Tag
	result := t.DB.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
