package scopes

import "gorm.io/gorm"

func JOINCustomerOrders() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("customer").Joins("INNER JOIN order ON order.id_customer = customer.id_customer")
	}
}

func FILTERStatusOrdersCustomer(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			return db.Where("order.status_order = ?", status)
		}
		return db
	}
}

func SELECTCustomerOrders() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("customer.nama_customer, order.kode_order, order.tanggal, order.status_order")
	}
}
