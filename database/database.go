package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mangofinance.com/bank-backend/helpers"
	"mangofinance.com/bank-backend/models"
)

func ConnectDB() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5431/bankapp"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	helpers.HandleErr(err)

	db.AutoMigrate(&models.User{}, &models.Account{})

	return db
}
