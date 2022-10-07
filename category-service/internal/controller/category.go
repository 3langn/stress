package controller

import (
	"category-service/internal/dto"
	"category-service/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryController struct {
	CategoryService service.CategoryService
	BaseController
	logger *zap.Logger
}

func InitCategoryController(c *gin.RouterGroup, s service.CategoryService, logger *zap.Logger) {
	controller := &CategoryController{
		CategoryService: s,
		logger:          logger,
	}
	g := c.Group("/category")
	g.GET("/:id", controller.GetByID)
	g.POST("/create", controller.CreateCategory)
	g.GET("/search", controller.SearchCategory)
}

func (b *CategoryController) GetByID(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	id := int64(idParam)
	bg, statusCode, err := b.CategoryService.GetByID(c, id)
	if err != nil {
		b.ResponseError(c, statusCode, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", bg)
	return
}

func (b *CategoryController) CreateCategory(c *gin.Context) {
	var req *dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	err := b.CategoryService.CreateCategory(c, req)
	if err != nil {
		b.ResponseError(c, http.StatusInternalServerError, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *CategoryController) SearchCategory(c *gin.Context) {
	var req dto.SearchCategoryRequest
	if err := c.Bind(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	res, total, err := b.CategoryService.SearchCategory(c, &req)
	if err != nil {
		b.ResponseError(c, http.StatusInternalServerError, []error{err})
		return
	}
	b.ResponseList(c, "success", total, res)
}
