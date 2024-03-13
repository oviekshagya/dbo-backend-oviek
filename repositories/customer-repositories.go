package repositories

import (
	"dbo-backend-oviek/models"
	"dbo-backend-oviek/models/raw"
	"dbo-backend-oviek/models/scopes"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepositories struct {
}

var UserRepositories = userRepositories{}

func (service userRepositories) GetCustomer(c *gin.Context) (interface{}, error) {
	db := c.MustGet("db").(*gorm.DB)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	idCustomer, _ := strconv.Atoi(c.Query("idCustomer"))

	var data []models.DataCustomer

	if idCustomer != 0 {
		var dataDetail models.DataCustomer
		db.Where("id_customer = ?", idCustomer).Take(&dataDetail)
		return dataDetail, nil
	}

	if page == 0 {
		db.Find(&data)
		return data, nil
	}

	var count int64
	var totalPage int

	switch {
	case pageSize > 100:
		pageSize = pageSize
	case pageSize <= 0:
		pageSize = 10
	}

	db.Scopes(scopes.Paginate(pageSize, page)).Find(&data)
	db.Table("customer").Count(&count)

	if int(count) < pageSize {
		totalPage = 1
	} else {
		totalPage = int(count) / pageSize
		if (int(count) % pageSize) != 0 {
			totalPage = totalPage + 1
		}
	}

	if page == 0 {
		page = 1
	}

	return map[string]interface{}{
		"Data":        data,
		"totalRecord": count,
		"page":        page,
		"pageSize":    pageSize,
		"totalPage":   totalPage,
	}, nil
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

	if input.IdCustomer != 0 {
		hashUpdate, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			trx.Rollback()
			return nil, fmt.Errorf("password tidak bisa di generate")
		}

		if updated := trx.Table("customer").Where("id_customer = ?", input.IdCustomer).Updates(map[string]interface{}{
			"nama_customer": input.NamaCustomer,
			"password":      string(hashUpdate),
			"email":         input.Email,
		}); updated.Error != nil {
			trx.Rollback()
			return nil, fmt.Errorf("update tidak bisa dilakukan")
		}
		trx.Commit()
		return &dataCrated, nil
	}

	if created := trx.Create(&dataCrated); created.Error != nil {
		trx.Rollback()
		return nil, fmt.Errorf("rollback created")
	}

	trx.Commit()
	return &dataCrated, nil
}

func (service userRepositories) DeleteCustomer(input raw.JSONRequestRegisterCustomer, c *gin.Context) (interface{}, error) {
	db := c.MustGet("db").(*gorm.DB)
	trx := db.Begin()

	var data models.Customer
	trx.Where("id_customer = ?", input.IdCustomer).Take(&data)

	if created := trx.Delete(&data); created.Error != nil {
		trx.Rollback()
		return nil, fmt.Errorf("rollback deleted")
	}

	trx.Commit()
	return map[string]interface{}{
		"data": data,
	}, nil
}
