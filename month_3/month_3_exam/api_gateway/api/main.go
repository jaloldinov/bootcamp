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

	apiV1.POST("/auth/sign-in", handlerV1.SignIn)
	apiV1.POST("/change-password", handlerV1.ChangePassword)

	// category
	apiV1.POST("/category/create", handlerV1.CreateCategory)
	apiV1.GET("/category/list", handlerV1.GetAllCategory)
	apiV1.GET("/category/get/:category_id", handlerV1.GetCategory)
	apiV1.PUT("/category/update/:category_id", handlerV1.UpdateCategory)
	apiV1.DELETE("/category/delete/:category_id", handlerV1.DeleteCategory)

	// product
	apiV1.POST("/product/create", handlerV1.CreateProduct)
	apiV1.GET("/product/list", handlerV1.GetAllProduct)
	apiV1.GET("/product/get/:product_id", handlerV1.GetProduct)
	apiV1.PUT("/product/update/:product_id", handlerV1.UpdateProduct)
	apiV1.DELETE("/product/delete/:product_id", handlerV1.DeleteProduct)

	// order
	apiV1.PUT("/order/update/status/:order_id", handlerV1.UpdateOrderStatus)

	apiV1.POST("/order/create", handlerV1.CreateOrder)
	apiV1.GET("/order/list", handlerV1.GetAllOrder)
	apiV1.GET("/order/get/:order_id", handlerV1.GetOrder)
	apiV1.PUT("/order/update/:order_id", handlerV1.UpdateOrder)
	apiV1.DELETE("/order/delete/:order_id", handlerV1.DeleteOrder)

	// delivery_tarif
	apiV1.POST("/delivery/create", handlerV1.CreateDeliveryTariff)
	apiV1.GET("/delivery/list", handlerV1.GetAllDeliveryTariff)
	apiV1.GET("/delivery/get/:delivery_id", handlerV1.GetDeliveryTariff)
	apiV1.PUT("/delivery/update/:delivery_id", handlerV1.UpdateDeliveryTariff)
	apiV1.DELETE("/delivery/delete/:delivery_id", handlerV1.DeleteDeliveryTariff)

	// branch
	apiV1.POST("/branch/create", handlerV1.CreateBranch)
	apiV1.GET("/branch/list", handlerV1.GetAllBranch)
	apiV1.GET("/branch/get/:branch_id", handlerV1.GetBranch)
	apiV1.PUT("/branch/update/:branch_id", handlerV1.UpdateBranch)
	apiV1.DELETE("/branch/delete/:branch_id", handlerV1.DeleteBranch)

	apiV1.GET("/branch/list/active", handlerV1.GetListActiveBranch)

	// user
	apiV1.POST("/user/create", handlerV1.CreateUser)
	apiV1.GET("/user/list", handlerV1.GetListUser)
	apiV1.GET("/user/get/:user_id", handlerV1.GetUser)
	apiV1.PUT("/user/update/:user_id", handlerV1.UpdateUser)
	apiV1.DELETE("/user/delete/:user_id", handlerV1.DeleteUser)

	// client
	apiV1.POST("/client/create", handlerV1.CreateClients)
	apiV1.GET("/client/list", handlerV1.GetListClients)
	apiV1.GET("/client/get/:client_id", handlerV1.GetClients)
	apiV1.PUT("/client/update/:client_id", handlerV1.UpdateClients)
	apiV1.DELETE("/client/delete/:client_id", handlerV1.DeleteClients)

	// courier
	apiV1.POST("/courier/create", handlerV1.CreateCourier)
	apiV1.GET("/courier/list", handlerV1.GetListCourier)
	apiV1.GET("/courier/get/:courier_id", handlerV1.GetCourier)
	apiV1.PUT("/courier/update/:courier_id", handlerV1.UpdateCourier)
	apiV1.DELETE("/courier/delete/:courier_id", handlerV1.DeleteCourier)

	apiV1.GET("/courier/active-orders/list", handlerV1.GetListActiveOrders)
	apiV1.GET("/courier/delete-order/:id", handlerV1.DeleteCourierInOrder)
	apiV1.GET("/courier/get-order/:id", handlerV1.CourierGetOrder)

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
