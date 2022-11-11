package handlers

import (
	"net/http"
	"strconv"

	"github/yorqinbek/CRUD/article/moduls"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     moduls.CreateAuthorModel true "author body"
// @Success     201    {object} moduls.JSONResponse{data=moduls.Author}
// @Failure     400    {object} moduls.JSONErrorResponse
// @Router      /author [post]
func (h Handlers) CreateAuthor(c *gin.Context) {
	var body moduls.CreateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	id := uuid.New()

	err := h.In.AddAuthor(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := h.In.GetAuthorByID(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, moduls.JSONResponse{
		Message: "Author | GetList",
		Data:    author,
	})
}

// GetAuthorByID godoc
// @Summary     get author by id
// @Description get an author by id
// @Tags        authors
// @Accept      json
// @Param       id path string true "Author ID"
// @Produce     json
// @Success     200 {object} moduls.JSONResponse{data=moduls.Author}
// @Failure     400 {object} moduls.JSONErrorResponse
// @Router      /author/{id} [get]
func (h Handlers) GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation

	author, err := h.In.GetAuthorByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.In.GetArticlesByAuthorID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author.Article = article

	c.JSON(http.StatusOK, moduls.JSONResponse{
		Message: "OK",
		Data:    author,
	})
}

// GetAuthorList godoc
// @Summary     List author
// @Description get author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset query    int false "0"
// @Param       limit  query    int false "10"
// @Param       search query    string false "smth"
// @Success     200 {object} moduls.JSONResponse{data=[]moduls.Author}
// @Router      /author [get]
func (h Handlers) GetAuthorList(c *gin.Context) {
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

	articleList, err := h.In.GetAuthorList(offset, limit, searchStr)
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

// UpdateAuthor godoc
// @Summary     Update author
// @Description update a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     moduls.UpdateAuthorModel true "author body"
// @Success     200     {object} moduls.JSONResponse{data moduls.Author}
// @Failure     400     {object} moduls.JSONErrorResponse
// @Router      /author [put]
func (h Handlers) UpdateAuthor(c *gin.Context) {
	var body moduls.UpdateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.In.UpdateAuthor(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := h.In.GetAuthorByID(body.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, moduls.JSONResponse{
		Message: "OK",
		Data:    author,
	})

}

// DeleteAuthor godoc
// @Summary     delete author by id
// @Description delete an author by id
// @Tags        authors
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Failure     400 {object} moduls.JSONErrorResponse
// @Router      /author/{id} [delete]
func (h Handlers) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	author, err := h.In.GetAuthorByID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.In.DeleteAuthor(author.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, moduls.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, moduls.JSONResponse{
		Message: "OK",
		Data:    author,
	})
}
