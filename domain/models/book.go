package models

import (
	"gorm.io/gorm"
)

type Book struct {
	Name   string `json:"name" gorm:"unique""`
	Author string `json:"author"`
	Year   int    `json:"year"`
	gorm.Model
}

type Books = []Book
