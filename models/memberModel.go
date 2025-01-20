package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name string
	Age  int
}
