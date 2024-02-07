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
	authorRouter := router.Group("/api/author")
	{
		authorRouter.GET("/", ctrl.GetAuthor)
		authorRouter.POST("/", ctrl.CreateAuthor)
		authorRouter.PUT("/", ctrl.UpdateAuthor)
		authorRouter.DELETE("/", ctrl.DeleteAuthor)
	}
	return &ServerConfig{
		Router: router,
	}
}

func (sc *ServerConfig) Run(address string) {
	sc.Router.Run(address)
}
