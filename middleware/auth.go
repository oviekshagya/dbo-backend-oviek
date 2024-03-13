package middleware

import (
	"dbo-backend-oviek/pkg"
	"net/http"

	"github.com/gin-gonic/gin"

)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeBaererJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"errorMessage": "unauthorize!",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"errorMessage": "unauthorize!",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		metadata, errs := pkg.ParseToken(pkg.ExtractTokenMiddleware(c.Request))
		if metadata == nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"errorMessage": errs.Error(),
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func AuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("api-key")
		if key != pkg.HEADERKEY {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"errorMessage": "invalid header access",
			})
			c.Abort()
			return
		}
		c.Next()

	}
}

func AuthBasic() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && user == pkg.AUTHUSER && password == pkg.AUTHPASSWORD {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"errorMessage": "invalid access",
			})
			c.Abort()
			return
		}

	}
}
