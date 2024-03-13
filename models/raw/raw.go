package raw

type JSONRequestRegisterCustomer struct {
	NamaCustomer string `gorm:"column:nama_customer;type:varchar(100)" json:"nama_customer"`
	Password     string `gorm:"column:password;type:varchar(100)" json:"password"`
	Email        string `gorm:"column:email;type:varchar(100)" json:"email"`
}
