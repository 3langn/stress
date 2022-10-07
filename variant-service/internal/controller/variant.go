package controller

import (
	"net/http"
	"strconv"
	"variant-service/internal/dto"
	"variant-service/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VariantController struct {
	VariantService service.VariantService
	BaseController
	logger *zap.Logger
}

func InitVariantController(c *gin.RouterGroup, s service.VariantService, logger *zap.Logger) {
	controller := &VariantController{
		VariantService: s,
		logger:         logger,
	}
	g := c.Group("/variant")
	g.GET("/:id", controller.GetByID)
	g.POST("/ids", controller.FindVariantsByIDs)
	g.POST("/create", controller.CreateVariant)
	g.GET("/search", controller.SearchVariant)
}

func (b *VariantController) FindVariantsByIDs(c *gin.Context) {
	var req dto.FindVariantsByIDsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	variants, total, code, err := b.VariantService.FindVariantsByIDs(c, &req)
	if err != nil {
		b.ResponseError(c, code, []error{err})
		return
	}
	b.ResponseList(c, "success", total, variants)
	return
}

func (b *VariantController) GetByID(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	id := int64(idParam)
	bg, code, err := b.VariantService.GetByID(c, id)
	if err != nil {
		b.ResponseError(c, code, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", bg)
	return
}

func (b *VariantController) CreateVariant(c *gin.Context) {
	var req *dto.CreateVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	statusCode, err := b.VariantService.CreateVariant(c, req)
	if err != nil {
		b.ResponseError(c, statusCode, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *VariantController) SearchVariant(c *gin.Context) {
	var req dto.SearchVariantRequest
	if err := c.Bind(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	res, total, err := b.VariantService.SearchVariant(c, &req)
	if err != nil {
		b.ResponseError(c, http.StatusInternalServerError, []error{err})
		return
	}
	b.ResponseList(c, "success", total, res)
}
