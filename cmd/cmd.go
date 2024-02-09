package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/terc1997/go-cms/internal/controllers"
)

type Application struct {
	Router *gin.Engine
}

func NewConfig() *Application {
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
	articleRouter := router.Group("/api/article")
	{
		articleRouter.POST("/", ctrl.CreateArticle)
	}

	return &Application{
		Router: router,
	}
}

func (sc *Application) Run(address string) {
	sc.Router.Run(address)
}
