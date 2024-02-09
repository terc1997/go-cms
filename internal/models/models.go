package models

import (
	"gorm.io/gorm"
)

type JSONAuthor struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type JSONArticle struct {
	Title    string `form:"title"`
	Content  string `form:"content"`
	AuthorID uint   `form:"authorid"`
	Email    string `form:"email"`
}

type APIResponseModel struct {
	Message string `json:"message"`
	Body    any    `json:"body"`
}

type Author struct {
	gorm.Model
	Name  string
	Email string
}

type Article struct {
	gorm.Model
	Title    string
	Content  string
	AuthorID uint
}
