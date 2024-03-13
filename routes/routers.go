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
	}
	customerAuth := v1.Group("/customer/auth")
	customerAuth.Use(middleware.AuthorizeBaererJWT(), middleware.AuthHeader())
	{
		customerAuth.GET("/data", controllers.CustomerController.GetCustomer)
	}

}
