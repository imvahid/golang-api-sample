package model

import (
	"time"
)

type Tag struct {
	ID        int        `gorm:"type:int;primary_key;autoIncrement"`
	Name      string     `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoCreateTime;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}
