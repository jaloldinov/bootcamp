package api

import (
	_ "market/api/docs"
	"market/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetByIDBranch)
	r.GET("/branch", h.GetListBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/category", h.CreateCategory)
	r.GET("/category/:id", h.GetByIDCategory)
	r.GET("/category", h.GetListCategory)
	r.PUT("/category/:id", h.UpdateCategory)
	r.DELETE("/category/:id", h.DeleteCategory)

	r.POST("/product", h.CreateProduct)
	r.GET("/product/:id", h.GetByIDProduct)
	r.GET("/product", h.GetListProduct)
	r.PUT("/product/:id", h.UpdateProduct)
	r.DELETE("/product/:id", h.DeleteProduct)

	r.POST("/coming_table", h.CreateComingTable)
	r.GET("/coming_table/:id", h.GetByIDComingTable)
	r.GET("/coming_table", h.GetListComingTable)
	r.PUT("/coming_table/:id", h.UpdateComingTable)
	r.PUT("/doincome/:id", h.UpdateComingTableStatus) // doincome change status 'in_process' to 'finished' by coming table ID
	r.DELETE("/coming_table/:id", h.DeleteComingTable)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
