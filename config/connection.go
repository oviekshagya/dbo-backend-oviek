package config

import (
	"dbo-backend-oviek/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	grmsql "gorm.io/driver/mysql"
	grm "gorm.io/gorm"

)

func SetupMainDBGORM() *grm.DB {
	configuration := GetConfig()

	database := configuration.Database.Dbname
	host := configuration.Database.Host
	port := configuration.Database.Port
	username := configuration.Database.Username
	password := configuration.Database.Password

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := grm.Open(grmsql.Open(addr), &grm.Config{})

	if err != nil {
		log.Println("err connect db main:", err)
	}
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Order{})
	return db
}
