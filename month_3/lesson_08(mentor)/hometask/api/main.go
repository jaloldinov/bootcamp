package api

import (
	_ "market/api/docs"
	"market/api/handler"
	"market/pkg/helper"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// r.Use(helper.StartMiddleware)
	// r.Use(helper.EndMiddleware)

	r.Use(helper.Logger)
	r.POST("/login", h.Login)
	r.POST("/person", h.CreatePerson)
	r.GET("/person", h.GetAllPersons)
	r.GET("/person/:id", h.GetPerson)
	r.PUT("/person/:id", h.UpdatePerson)
	r.DELETE("/person/:id", h.DeletePerson)
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
