package models

import (
	"gorm.io/gorm"

	"ilmudata/task1/database"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type DbModels struct {
	// declare variables
	Db *gorm.DB
}

func InitDbModels() *DbModels {
	db := database.InitDb()
	// gorm
	// db.AutoMigrate(&User{})
	db.AutoMigrate(&Video{})
	// db.AutoMigrate(&Advert{})

	return &DbModels{Db: db}
}
