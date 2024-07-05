package config

import "gorm.io/gorm"

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=password dbname= port=5432 sslmode=disable"
}