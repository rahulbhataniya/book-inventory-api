package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Category  string `json:"category"`
	Published string `json:"published" binding:"required"` // for simplicity
}
