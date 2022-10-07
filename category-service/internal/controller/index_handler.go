package controller

import (
	"category-service/internal/dto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func InitIndexController(c *gin.RouterGroup) {
	controller := &IndexController{}
	c.GET("/health", controller.Health)
}

func (index *IndexController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, dto.SimpleResponse{
		Message: "ok",
		Data:    time.Now().Format("2006-01-02 15:04:05"),
	})
	return
}
