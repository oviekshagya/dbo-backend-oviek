package raw

type JSONRequestRegisterCustomer struct {
	IdCustomer   int    `json:"id_customer"`
	NamaCustomer string `gorm:"column:nama_customer;type:varchar(100)" json:"nama_customer"`
	Password     string `gorm:"column:password;type:varchar(100)" json:"password"`
	Email        string `gorm:"column:email;type:varchar(100)" json:"email"`
}

type JSONRequestLogin struct {
	Password string `gorm:"column:password;type:varchar(100)" json:"password"`
	Email    string `gorm:"column:email;type:varchar(100)" json:"email"`
}
