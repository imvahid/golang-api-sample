package repository

import "github.com/imvahid/golang-api-sample/model"

type TagRepository interface {
	Save(tag model.Tag)
	Update(tag model.Tag)
	Delete(tagID int)
	FindByID(tagID int) (tag model.Tag, err error)
	FindAll() []model.Tag
}
