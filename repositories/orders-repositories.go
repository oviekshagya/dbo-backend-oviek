package repositories

import (
	"dbo-backend-oviek/models"
	"dbo-backend-oviek/models/scopes"
	"dbo-backend-oviek/pkg"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

type orderRepositories struct {
}

var OrderRepositories = orderRepositories{}

func (service orderRepositories) InsertUpdateOrders(input models.Order, c *gin.Context) (*models.Order, error) {
	db := c.MustGet("db").(*gorm.DB)
	trx := db.Begin()

	dataCrated := models.Order{
		IdCustomer:  input.IdCustomer,
		KodeOrder:   pkg.KodeRand(10),
		Tanggal:     time.Now(),
		StatusOrder: "waiting",
		Produk:      input.Produk,
	}

	if input.IdOrder != 0 {

		if updated := trx.Table("order").Where("id_order = ?", input.IdOrder).Updates(map[string]interface{}{
			"id_customer":  input.IdCustomer,
			"status_order": input.StatusOrder,
			"produk":       input.Produk,
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

func (service userRepositories) GetOrder(c *gin.Context) (interface{}, error) {
	db := c.MustGet("db").(*gorm.DB)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	idOrder, _ := strconv.Atoi(c.Query("idOrder"))
	idCustomer, _ := strconv.Atoi(c.Query("idCustomer"))
	search := c.Query("search")
	status := c.Query("status")

	var data []models.CustomerOrder

	if idOrder != 0 {
		var dataDetail models.OrderCustomerDetail
		db.Where("id_order = ?", idOrder).Preload("Customer").Take(&dataDetail)
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

	db.Scopes(scopes.Paginate(pageSize, page), scopes.JOINCustomerOrders(), scopes.FILTERStatusOrdersCustomer(status), scopes.FILTEROrdersCustomer(idCustomer), scopes.SELECTCustomerOrders()).Scopes(func(d *gorm.DB) *gorm.DB {
		if search != "" {
			return d.Where("customer.nama_customer LIKE ?", fmt.Sprintf("%%%s%%", search))
		}
		return d
	}).Find(&data)

	db.Scopes(scopes.Paginate(pageSize, page), scopes.JOINCustomerOrders(), scopes.FILTERStatusOrdersCustomer(status), scopes.FILTEROrdersCustomer(idCustomer)).Scopes(func(d *gorm.DB) *gorm.DB {
		if search != "" {
			return d.Where("customer.nama_customer LIKE ?", fmt.Sprintf("%%%s%%", search))
		}
		return d
	}).Count(&count)

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

func (service orderRepositories) DeleteOrder(id int, input models.Order, c *gin.Context) (interface{}, error) {
	db := c.MustGet("db").(*gorm.DB)
	trx := db.Begin()

	var data models.Order
	if id != 0 || input.IdOrder != 0 {
		trx.Where("id_customer = ? OR id_order = ?", id, input.IdOrder).Take(&data)
	} else if id != 0 {
		trx.Where("id_customer = ?", id).Take(&data)
	}
	//trx.Where("id_order = ?", input.IdOrder).Take(&data)

	if created := trx.Delete(&data); created.Error != nil {
		trx.Rollback()
		return nil, fmt.Errorf("rollback deleted")
	}

	trx.Commit()
	return map[string]interface{}{
		"data": data,
	}, nil
}
