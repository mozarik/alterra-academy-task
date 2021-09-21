package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost  user=postgres password=mysecretpassword port=5432 sslmode=disable"
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_ = DB
}
