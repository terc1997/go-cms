package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/terc1997/go-cms/internal/controllers"
)

type ServerConfig struct {
	Router *gin.Engine
}

func NewConfig() *ServerConfig {
	router := gin.Default()
	ctrl := controllers.Controller{}
	ctrl.Init()
	authorRouter := router.Group("/author")
	{
		authorRouter.GET("/get", ctrl.GetAuthor)
		authorRouter.POST("/create", ctrl.CreateAuthor)
	}
	return &ServerConfig{
		Router: router,
	}
}

func (sc *ServerConfig) Run(address string) {
	sc.Router.Run(address)
}
