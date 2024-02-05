package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/terc1997/go-cms/internal/models"
)

type JSONAuthor struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type APIResponseModel struct {
	Message string
	Body    any
}

type Controller struct {
	m models.Models
}

func (ctrl *Controller) Init() {
	path := os.Getenv("DB_PATH")
	if path != "" {
		ctrl.m = *models.NewModel(path)
	} else {
		ctrl.m = *models.NewModel("/Users/tarsis/Documents/Studies/Go/go-cms/cms.db")
	}

}

func (ctrl *Controller) GetAuthor(c *gin.Context) {
	var jsonAuthor JSONAuthor
	c.Bind(&jsonAuthor)
	log.Println(jsonAuthor)
	author, err := ctrl.m.GetAuthor(jsonAuthor.Email)
	log.Println(author)
	if err != nil {
		c.JSON(http.StatusNotFound, &APIResponseModel{
			Message: "User not found",
			Body:    "",
		})
	} else {
		c.JSON(http.StatusOK, &APIResponseModel{
			Message: "User succesfully found",
			Body:    &author,
		})
	}

}

func (ctrl *Controller) CreateAuthor(c *gin.Context) {
	var jsonAuthor JSONAuthor
	c.Bind(&jsonAuthor)
	log.Printf("Received from API: %v\n", jsonAuthor)
	status, err := ctrl.m.CreateAuthor(jsonAuthor.Name, jsonAuthor.Email)
	if err != nil {
		panic("Failed to create user")
	}
	log.Printf("Creation status: %d\n", status)
	c.JSON(http.StatusOK, "User successfully created")
}
