package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terc1997/go-cms/internal/db"
	"github.com/terc1997/go-cms/internal/models"
	"gorm.io/gorm"
)

type Controller struct {
	dbc db.DBConfig
	DB  *gorm.DB
}

func (ctrl *Controller) Init() {
	ctrl.dbc = *db.NewDBConfig()
}

func (ctrl *Controller) GetAuthor(c *gin.Context) {
	method := "GetAuthor\t"
	var jsonAuthor models.JSONAuthor
	c.Bind(&jsonAuthor)
	log.Printf("%s Received from API %v\n", method, jsonAuthor)
	author, err := ctrl.dbc.GetAuthor(jsonAuthor.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, &models.APIResponseModel{
			Message: "User not found",
			Body:    "",
		})
	} else {
		c.JSON(http.StatusOK, &models.APIResponseModel{
			Message: "User succesfully found",
			Body:    &author,
		})
	}

}

func (ctrl *Controller) CreateAuthor(c *gin.Context) {
	var jsonAuthor models.JSONAuthor
	method := "CreateAuthor\t"
	c.Bind(&jsonAuthor)
	log.Printf("%s Received from API: %v\n", method, jsonAuthor)
	status, err := ctrl.dbc.CreateAuthor(jsonAuthor.Name, jsonAuthor.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponseModel{
			Message: "User not created",
			Body:    "",
		})
	} else {
		log.Printf("Creation status: %d\n", status)
		c.JSON(http.StatusOK, &models.APIResponseModel{
			Message: "User created sucessfully",
			Body:    "",
		})
	}
}

func (ctrl *Controller) UpdateAuthor(c *gin.Context) {
	method := "UpdateAuthor\t"
	var jsonAuthor models.JSONAuthor
	c.Bind(&jsonAuthor)
	log.Printf("%s Received from API: %v", method, jsonAuthor)
	err := ctrl.dbc.UpdateAuthor(jsonAuthor.Name, jsonAuthor.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponseModel{
			Message: "User not updated",
			Body:    "",
		})
	} else {
		c.JSON(http.StatusOK, &models.APIResponseModel{
			Message: "User updated sucessfully",
			Body:    "",
		})
	}

}

func (ctrl *Controller) DeleteAuthor(c *gin.Context) {
	method := "DeleteAuthor\t"
	var jsonAuthor models.JSONAuthor
	c.Bind(&jsonAuthor)
	log.Printf("%s Received from API: %v", method, jsonAuthor)
	err := ctrl.dbc.DeleteAuthor(jsonAuthor.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponseModel{
			Message: "User not deleted",
			Body:    "",
		})
	} else {
		c.JSON(http.StatusOK, &models.APIResponseModel{
			Message: "User deleted sucessfully",
			Body:    "",
		})
	}
}

func (ctrl *Controller) CreateArticle(c *gin.Context) {
	method := "CreateArticle\t"
	var jsonArticle models.JSONArticle
	c.Bind(&jsonArticle)
	log.Printf("%s Received from API: %v\n", method, jsonArticle)
	status, err := ctrl.dbc.CreateArticle(jsonArticle.Title, jsonArticle.Content, jsonArticle.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponseModel{
			Message: "User not created",
			Body:    "",
		})
	} else {
		log.Printf("Creation status: %d\n", status)
		c.JSON(http.StatusOK, &models.APIResponseModel{
			Message: "Article created sucessfully",
			Body:    "",
		})
	}
}
