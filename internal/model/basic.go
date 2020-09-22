package model

import "time"

type DbBasicModel struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
