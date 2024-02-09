package db

import (
	"log"
	"os"

	"github.com/terc1997/go-cms/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB *gorm.DB
}

func NewDBConfig() *DBConfig {
	path := os.Getenv("DB_PATH")
	var db *gorm.DB
	var err error
	if path != "" {
		db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("/Users/tarsis/Documents/Studies/Go/go-cms/cms.db"), &gorm.Config{})
	}
	if err != nil {
		panic("Could not open connection to DB")
	}
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Article{})
	return &DBConfig{
		DB: db,
	}
}

func (dbc *DBConfig) CreateAuthor(name, email string) (uint, error) {
	author := models.Author{
		Name:  name,
		Email: email,
	}

	result := dbc.DB.Create(&author)

	return author.ID, result.Error
}

func (dbc *DBConfig) GetAuthor(email string) (result []models.Author, err error) {
	log.Printf("Author Email: %v\n", email)
	if email == "" {
		ret := dbc.DB.Find(&result)
		err = ret.Error
	} else {
		ret := dbc.DB.Where(&models.Author{Email: email}).First(&result)
		err = ret.Error
	}
	return
}

func (dbc *DBConfig) DeleteAuthor(email string) error {
	author, err := dbc.GetAuthor(email)
	if err != nil {
		log.Fatal("User not found")
		return err
	}
	result := dbc.DB.Unscoped().Delete(&author)
	return result.Error
}

func (dbc *DBConfig) UpdateAuthor(name, email string) error {
	var author models.Author
	dbc.DB.Where("email = ?", email).Find(&author)
	author.Name = name
	author.Email = email

	result := dbc.DB.Save(&author)

	return result.Error
}

func (dbc *DBConfig) GetArticles() (result []models.Article, err error) {
	ret := dbc.DB.Find(&result)
	err = ret.Error
	return
}

func (dbc *DBConfig) CreateArticle(title, content, email string) (uint, error) {
	author, err := dbc.GetAuthor(email)
	if err != nil {
		log.Fatal("Invalid author")
	}
	article := models.Article{
		Title:    title,
		Content:  content,
		AuthorID: author[0].ID,
	}

	result := dbc.DB.Create(&article)

	return article.ID, result.Error
}
