package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string
	Email string
}

type Article struct {
	gorm.Model
	Title    string
	Content  string
	AuthorID string
}

type Models struct {
	DB *gorm.DB
}

func NewModel(path string) *Models {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("Could not open connection to DB")
	}
	db.AutoMigrate(&Author{})
	return &Models{
		DB: db,
	}
}

func (m *Models) CreateAuthor(name, email string) (uint, error) {
	author := Author{
		Name:  name,
		Email: email,
	}

	result := m.DB.Create(&author)

	return author.ID, result.Error
}

func (m *Models) GetAuthor(email string) (result Author, err error) {
	log.Printf("Author Email: %v\n", email)
	ret := m.DB.Where(&Author{Email: email}).First(&result)
	err = ret.Error
	return
}

func (m *Models) DeleteAuthor(email string) error {
	author, err := m.GetAuthor(email)
	if err != nil {
		log.Fatal("User not found")
		return err
	}
	result := m.DB.Unscoped().Delete(&author)
	return result.Error
}

func (m *Models) UpdateAuthor(name, email string) error {
	var author Author
	m.DB.Where("email = ?", email).Find(&author)
	author.Name = name
	author.Email = email

	result := m.DB.Save(&author)

	return result.Error
}
