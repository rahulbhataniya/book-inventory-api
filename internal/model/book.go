package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	Category  string `json:"category"`
	Published string `json:"published"` // for simplicity
}
