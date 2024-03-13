package repositories

import (
	"dbo-backend-oviek/models"
	"dbo-backend-oviek/models/raw"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepositories struct {
}

var UserRepositories = userRepositories{}

func (service userRepositories) GetAllCustomer(c *gin.Context) (interface{}, error) {
	//db := c.MustGet("db").(*gorm.DB)

	return nil, nil
}

func (service userRepositories) RegisterCustomer(input raw.JSONRequestRegisterCustomer, c *gin.Context) (*models.Customer, error) {
	db := c.MustGet("db").(*gorm.DB)
	trx := db.Begin()
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		trx.Rollback()
		return nil, fmt.Errorf("password tidak bisa di generate")

	}

	dataCrated := models.Customer{
		NamaCustomer:  input.NamaCustomer,
		Password:      string(hash),
		Email:         input.Email,
		TglRegistrasi: time.Now(),
	}

	if created := trx.Create(&dataCrated); created.Error != nil {
		trx.Rollback()
		return nil, fmt.Errorf("rollback created")
	}

	trx.Commit()
	return &dataCrated, nil
}
