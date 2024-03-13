package models

import "time"

type Customer struct {
	IdCustomer    int       `gorm:"column:id_customer;primaryKey" json:"id_customer"`
	NamaCustomer  string    `gorm:"column:nama_customer;type:varchar(100)" json:"nama_customer"`
	Password      string    `gorm:"column:password;type:varchar(100)" json:"password"`
	Email         string    `gorm:"column:email;type:varchar(100)" json:"email"`
	TglRegistrasi time.Time `gorm:"column:tgl_registrasi;type:datetime" json:"tgl_registrasi"`
}

func (Customer) TableName() string {
	return "customer"
}

type DataCustomer struct {
	IdCustomer    int       `gorm:"column:id_customer;primaryKey" json:"id_customer"`
	NamaCustomer  string    `gorm:"column:nama_customer;type:varchar(100)" json:"nama_customer"`
	Email         string    `gorm:"column:email;type:varchar(100)" json:"email"`
	TglRegistrasi time.Time `gorm:"column:tgl_registrasi;type:datetime" json:"tgl_registrasi"`
}

func (DataCustomer) TableName() string {
	return "customer"
}
