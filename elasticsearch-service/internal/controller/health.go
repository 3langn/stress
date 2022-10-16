package controller

import (
	"net/http"
	"search-service/internal/dto"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func InitHealthController(c *gin.RouterGroup) {
	controller := &HealthController{}
	c.GET("/health", controller.Health)
}

func (index *HealthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, dto.SimpleResponse{
		Message: "success",
		Data:    time.Now().Format("2006-01-02 15:04:05"),
	})
	return
}