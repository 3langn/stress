package controller

import (
	"net/http"
	"time"
	"variant-service/internal/dto"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func InitIndexController(c *gin.RouterGroup) {
	controller := &HealthController{}
	c.GET("/health", controller.Health)
}

func (index *HealthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, dto.SimpleResponse{
		Message: "ok",
		Data:    time.Now().Format("2006-01-02 15:04:05"),
	})
	return
}
