package handler

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/tuyentv96/lamda-echo-template/models"
	"net/http"
	"strconv"
	"time"
)

var articleStore map[int]*models.Article
var articleStoreSeq int

type ArticleHandler struct {
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(e *echo.Echo) {
	articleStore = make(map[int]*models.Article)
	handler := &ArticleHandler{}
	e.GET("/articles", handler.GetAll)
	e.POST("/articles", handler.Create)
	e.GET("/articles/:id", handler.GetByID)
}

func (h *ArticleHandler) Create(c echo.Context) error {
	var article models.Article
	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	articleStoreSeq++
	article.ID = articleStoreSeq
	article.CreatedAt = time.Now().UTC()
	article.UpdatedAt = time.Now().UTC()
	articleStore[article.ID] = &article

	return c.JSON(http.StatusCreated, article)
}

func (h *ArticleHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	article, ok := articleStore[idP]
	if !ok {
		return c.JSON(http.StatusNotFound, errors.New("not article found"))
	}

	return c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) GetAll(c echo.Context) error {
	var res []*models.Article
	for _, a := range articleStore {
		res = append(res, a)
	}

	return c.JSON(http.StatusOK, res)
}
