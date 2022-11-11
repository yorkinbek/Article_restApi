package main

import (
	"github/yorqinbek/CRUD/article/handlers"
	"github/yorqinbek/CRUD/article/storage"
	"github/yorqinbek/CRUD/article/storage/postgres"

	docs "github/yorqinbek/CRUD/article/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	docs.SwaggerInfo.Title = "Swagger Examle"
	docs.SwaggerInfo.Description = "About Article with SWagger"
	docs.SwaggerInfo.Version = "1.0"
	r := gin.Default()
	var err error
	var stg storage.StorageI

	stg, err = postgres.InitDB("user=yorqin dbname=article password='yorqinbek' sslmode=disable")
	if err != nil {
		panic(err)
	}

	h := handlers.Handlers{
		In: stg,
	}

	r.POST("/article", h.CreateArticle)
	r.GET("/article", h.GetArticleList)
	r.GET("/article/:id", h.GetArticleByID)
	r.PUT("/article", h.UpdateArticle)
	r.DELETE("/article", h.DeleteArticle)

	r.POST("/author", h.CreateAuthor)
	r.GET("/author", h.GetAuthorList)
	r.GET("/author/:id", h.GetAuthorByID)
	r.PUT("/author", h.UpdateAuthor)
	r.DELETE("/author", h.DeleteAuthor)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":3000")
}
