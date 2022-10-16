package controller

import (
	"search-service/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SearchController struct {
	SearchService service.SearchService
	logger        *zap.Logger
	BaseController
}

func InitSearchController(c *gin.RouterGroup, groupService service.SearchService, logger *zap.Logger) {
	controller := &SearchController{
		SearchService: groupService,
		logger:        logger,
	}
	g := c.Group("/search")
	g.GET("/", controller.Search)
}

func (s *SearchController) Search(c *gin.Context) {

}
