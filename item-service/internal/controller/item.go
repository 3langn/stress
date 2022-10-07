package controller

import (
	"item-service/internal/dto"
	"item-service/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ItemController struct {
	ItemService service.ItemService
	logger      *zap.Logger
	BaseController
}

func InitItemController(c *gin.RouterGroup, groupService service.ItemService, logger *zap.Logger) {
	controller := &ItemController{
		ItemService: groupService,
		logger:      logger,
	}
	g := c.Group("/item")
	g.GET("/:id", controller.GetByID)
	g.GET("/search", controller.Search)
	g.POST("/create", controller.Create)
}

// GetByID godoc
// @Summary Get Item By ID
// @Description Get Item By ID
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} SimpleResponse{data=models.Item} "Item Info"
// @Failure 400,401,404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /business_groups/{id} [get]
func (b *ItemController) GetByID(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	id := int64(idParam)
	bg, err := b.ItemService.GetByID(c, id)
	if err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", bg)
	return
}

func (b *ItemController) Search(c *gin.Context) {
	var req dto.SearchItemRequest

	if err := c.Bind(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}

	res, total, statusCode, err := b.ItemService.Search(c, req)
	if err != nil {
		b.ResponseError(c, statusCode, []error{err})
		return
	}
	b.ResponseList(c, "success", total, res)
	return
}

func (b *ItemController) Create(c *gin.Context) {
	var req dto.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	res, statusCode, err := b.ItemService.Create(c, req)
	if err != nil {
		b.ResponseError(c, statusCode, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", res)
	return
}
