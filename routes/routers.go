package routes

import (
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
	configs.AddAllowHeaders("Authorization, signature")
	Routers.Use(cors.New(configs))
	//Routers.NoRoute(middleware.NoRouteHandler())
	Routers.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "GIN GONIC TEMPLATE API BY Oviek Shagya")
	})

}
