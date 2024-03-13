package controllers

import (
	"dbo-backend-oviek/models"
	"dbo-backend-oviek/pkg"
	"dbo-backend-oviek/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
}

var OrderController = orderController{}

func (controller orderController) InsertUpdateOrders(c *gin.Context) {
	_, errMetaData := pkg.ExtractToken(c.Request)
	if errMetaData != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"errorMessage": errMetaData.Error(),
		})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}

	result, err := repositories.OrderRepositories.InsertUpdateOrders(input, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, result)
	return

}

func (controller orderController) GetOrder(c *gin.Context) {
	//metaData, _ := pkg.ExtractToken(c.Request)

	result, err := repositories.UserRepositories.GetOrder(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func (controller orderController) DeleteOrder(c *gin.Context) {
	metaData, errMetaData := pkg.ExtractToken(c.Request)
	if errMetaData != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"errorMessage": errMetaData.Error(),
		})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}

	result, err := repositories.OrderRepositories.DeleteOrder(metaData.UserId, input, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, result)
	return
}
