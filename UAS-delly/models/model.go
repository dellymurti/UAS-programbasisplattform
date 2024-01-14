package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)


func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	
	DB.AutoMigrate(&User{}, &Produk{})
}


type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Produk struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
