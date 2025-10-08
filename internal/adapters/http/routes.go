package http

import (
	"fsm-modulo-three/internal/adapters/handler"
	"fsm-modulo-three/internal/fsm"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes wires handlers to routes
func RegisterRoutes(r *gin.Engine) {
	service := &fsm.ModuloService{}
	h := handler.NewModuloHandler(service)

	api := r.Group("/api")
	{
		api.POST("/check", h.Check)

	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
