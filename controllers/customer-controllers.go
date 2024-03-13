package controllers

import (
	"dbo-backend-oviek/models/raw"
	"dbo-backend-oviek/pkg"
	"dbo-backend-oviek/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerController struct {
}

var CustomerController = customerController{}

func (controller customerController) InsertUpdateCustomer(c *gin.Context) {
	var input raw.JSONRequestRegisterCustomer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}

	result, err := repositories.UserRepositories.RegisterCustomer(input, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}

	if c.Request.Method == "POST" {
		createdToken, errToken := pkg.CreateToken(result.IdCustomer)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"errorMessage": errToken.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, map[string]interface{}{
			"message":      "success",
			"accessToken":  createdToken.AccessToken,
			"refreshToken": createdToken.RefreshToken,
		})
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    result,
	})
	return

}

func (controller customerController) GetCustomer(c *gin.Context) {
	_, errMetaData := pkg.ExtractToken(c.Request)
	if errMetaData != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"errorMessage": errMetaData.Error(),
		})
		return
	}

	result, err := repositories.UserRepositories.GetCustomer(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
	return
}

func (controller customerController) DeleteCustomer(c *gin.Context) {
	var input raw.JSONRequestRegisterCustomer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}

	result, err := repositories.UserRepositories.DeleteCustomer(input, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, result)
	return
}
