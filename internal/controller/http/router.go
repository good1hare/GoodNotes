package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter -.
// Swagger spec:
// @title       MateMind API
// @description Using a service for quick notes
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
