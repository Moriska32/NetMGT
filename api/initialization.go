package api

import (
	"ivan_search/api/services"
	"ivan_search/configuration"

	"github.com/gin-gonic/gin"
)

// InitAPI Инициализация API
func InitAPI(gr *gin.RouterGroup, dbConnections *configuration.Connections) {
	// Инициализация логгера
	gr.Use(gin.Logger())

	// Инициализация сервисов
	gr.GET("/ping", services.SayHello())
	gr.GET("/search/:search_field", services.SearchData())
	gr.GET("/search_v2", services.SearchDataV2())
}
