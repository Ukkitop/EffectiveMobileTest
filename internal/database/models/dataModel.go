package models

import (
	"gorm.io/gorm"
)

type DataModel struct {
	gorm.Model
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Sex        bool
	Nation     string
}
