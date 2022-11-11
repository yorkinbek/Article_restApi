package handlers

import (
	"net/http"
	"strconv"

	"github/yorqinbek/CRUD/article/moduls"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateArticle godoc
// @Summary     Create article
// @Description create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     moduls.CreateArticleModel true "article body"
// @Success     201     {object} moduls.JSONResponse{data moduls.Article}
// @Failure     400     {object} moduls.JSONErrorResponse
// @Router      /article [post]
func (h Handlers) CreateArticle(c *gin.Context) {
	var body moduls.CreateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	id := uuid.New()

	err := h.In.AddArticle(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.In.GetArticleByID(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, moduls.JSONResponse{
		Message: "Article | GetList",
		Data:    article,
	})
}

// GetArticleByID godoc
// @Summary     get article by id
// @Description get an article by id
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} moduls.JSONResponse{data moduls.FullArticleModuls}
// @Failure     400 {object} moduls.JSONErrorResponse
// @Router      /article/{id} [get]
func (h Handlers) GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation

	article, err := h.In.GetArticleByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, moduls.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}

// GetArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset query    int false "0"
// @Param       limit  query    int false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} moduls.JSONResponse{data  moduls.Article}
// @Router      /article [get]
func (h Handlers) GetArticleList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.In.GetArticleList(offset, limit, searchStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, moduls.JSONResponse{
		Message: "OK",
		Data:    articleList,
	})
}

// UpdateArticle godoc
// @Summary     Update article
// @Description update a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     moduls.UpdateArticleModel true "article body"
// @Success     200     {object} moduls.JSONResponse{data moduls.Article}
// @Failure     400     {object} moduls.JSONErrorResponse
// @Router      /article [put]
func (h Handlers) UpdateArticle(c *gin.Context) {
	var body moduls.UpdateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.In.UpdateArticle(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.In.GetArticleByID(body.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, moduls.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}

// DeleteArticle godoc
// @Summary     delete article by id
// @Description delete an article by id
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} moduls.JSONResponse{data moduls.FullArticleModuls}
// @Failure     400 {object} moduls.JSONErrorResponse
// @Router      /article/{id} [delete]
func (h Handlers) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	article, err := h.In.GetArticleByID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.In.DeleteArticle(article.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, moduls.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}
