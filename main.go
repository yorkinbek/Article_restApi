package main

import (
	"fmt"
	"github/yorqinbek/CRUD/article/config"
	"github/yorqinbek/CRUD/article/handlers"
	"github/yorqinbek/CRUD/article/storage"
	"github/yorqinbek/CRUD/article/storage/postgres"
	"net/http"

	docs "github/yorqinbek/CRUD/article/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.Load()

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	var err error
	var stg storage.StorageI
	stg, err = postgres.InitDB(psqlConnString)
	if err != nil {
		panic(err)
	}

	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	if cfg.Environment != "production" {
		r.Use(gin.Logger(), gin.Recovery())
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	h := handlers.NewHandler(stg, cfg)

	docs.SwaggerInfo.Description = "About Article with SWagger"
	docs.SwaggerInfo.Version = "1.0"

	stg, err = postgres.InitDB(psqlConnString)
	if err != nil {
		panic(err)
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort)
}
