package routes

import (
	"dbo-backend-oviek/controllers"
	"dbo-backend-oviek/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

)

var (
	Routers = gin.Default()
)

func Setup() {

	Routers.Use(requestid.New())
	/*Routers.Use(gin.Recovery())*/
	configs := cors.DefaultConfig()
	configs.AllowAllOrigins = true
	configs.AllowCredentials = true
	configs.AddAllowHeaders("Authorization, api-key")
	Routers.Use(cors.New(configs))
	//Routers.NoRoute(middleware.NoRouteHandler())
	Routers.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "GIN GONIC TEMPLATE API BY Oviek Shagya")
	})

	v1 := Routers.Group("/v1")
	v1.Use(middleware.CORS())
	customer := v1.Group("/customer")
	customer.Use(middleware.AuthBasic(), middleware.AuthHeader())
	{
		customer.POST("/", controllers.CustomerController.InsertUpdateCustomer)
		customer.PUT("/", controllers.CustomerController.InsertUpdateCustomer)
		customer.DELETE("/", controllers.CustomerController.DeleteCustomer)
		customer.POST("/login", controllers.CustomerController.Login)
	}
	customerAuth := v1.Group("/customer/auth")
	customerAuth.Use(middleware.AuthorizeBaererJWT(), middleware.AuthHeader())
	{
		customerAuth.GET("/data", controllers.CustomerController.GetCustomer)
	}

	order := v1.Group("/order")
	order.Use(middleware.AuthorizeBaererJWT(), middleware.AuthHeader())
	{
		order.GET("/data", controllers.OrderController.GetOrder)
		order.POST("/", controllers.OrderController.InsertUpdateOrders)
		order.PUT("/", controllers.OrderController.InsertUpdateOrders)
		order.DELETE("/", controllers.OrderController.DeleteOrder)
	}

}
