package api

import (
	"api_gateway/config"
	"api_gateway/pkg/logger"
	"api_gateway/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api_gateway/api/docs"
	v1 "api_gateway/api/handlers/v1"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))
	// router.Use(MaxAllowed(100))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	apiV1 := router.Group("/v1")

	// user
	apiV1.POST("/branch/create", handlerV1.CreateBranch)
	apiV1.GET("/branch/list", handlerV1.GetAllBranch)
	apiV1.GET("/branch/get/:branch_id", handlerV1.GetBranch)
	apiV1.PUT("/branch/update/:branch_id", handlerV1.UpdateBranch)
	apiV1.DELETE("/branch/delete/:branch_id", handlerV1.DeleteBranch)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}
