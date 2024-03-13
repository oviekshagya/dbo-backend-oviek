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

func (controller customerController) RegisterCustomer(c *gin.Context) {
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
