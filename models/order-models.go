package models

import "time"

type Order struct {
	IdOrder     int       `gorm:"column:id_order;type:int;primaryKey" json:"id_order"`
	IdCustomer  int       `gorm:"column:id_customer;type:int" json:"id_customer"`
	KodeOrder   string    `gorm:"column:kode_order;type:varchar(10)" json:"kode_order"`
	Tanggal     time.Time `gorm:"column:tanggal;type:datetime" json:"tanggal"`
	StatusOrder string    `gorm:"column:status_order;type:varchar(20)" json:"status_order"`
	Produk      string    `gorm:"column:produk;type:varchar(255)" json:"produk"`
}

func (Order) TableName() string {
	return "order"
}

type OrderCustomerDetail struct {
	IdOrder     int       `gorm:"column:id_order;type:int;primaryKey" json:"id_order"`
	IdCustomer  int       `gorm:"column:id_customer;type:int"`
	KodeOrder   string    `gorm:"column:kode_order;type:varchar(10)" json:"kode_order"`
	Tanggal     time.Time `gorm:"column:tanggal;type:datetime" json:"tanggal"`
	StatusOrder string    `gorm:"column:status_order;type:varchar(20)" json:"status_order"`
	Customer    *Customer `gorm:"foreignKey:id_customer;references:IdCustomer" json:"customer,omitempty"`
}

func (OrderCustomerDetail) TableName() string {
	return "order"
}

type CustomerOrder struct {
	NamaCustomer string    `gorm:"column:nama_customer;type:varchar(100)" json:"nama_customer"`
	KodeOrder    string    `gorm:"column:kode_order;type:varchar(10)" json:"kode_order"`
	Tanggal      time.Time `gorm:"column:tanggal;type:datetime" json:"tanggal"`
	StatusOrder  string    `gorm:"column:status_order;type:varchar(20)" json:"status_order"`
}
