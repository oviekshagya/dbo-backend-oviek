package models

import "time"

type Order struct {
	IdOrder     int       `gorm:"column:id_order;type:int;primaryKey" json:"id_order"`
	IdCustomer  int       `gorm:"column:id_customer;type:int"`
	KodeOrder   string    `gorm:"column:kode_order;type:varchar(10)" json:"kode_order"`
	Tanggal     time.Time `gorm:"column:tanggal;type:datetime" json:"tanggal"`
	StatusOrder string    `gorm:"column:status_order;type:varchar(20)" json:"status_order"`
}

func (Order) TableName() string {
	return "order"
}
