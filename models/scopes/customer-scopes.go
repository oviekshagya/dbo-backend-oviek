package scopes

import "gorm.io/gorm"

func JOINCustomerOrders() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("customer").Joins("INNER JOIN `order` as b ON b.id_customer = customer.id_customer")
	}
}

func FILTERStatusOrdersCustomer(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			return db.Where("b.status_order = ?", status)
		}
		return db
	}
}

func FILTEROrdersCustomer(idCustomer int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if idCustomer != 0 {
			return db.Where("b.id_customer = ?", idCustomer)
		}
		return db
	}
}

func SELECTCustomerOrders() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("customer.nama_customer, b.kode_order, b.tanggal, b.status_order")
	}
}
