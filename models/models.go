package models

import (
	"gorm.io/gorm"
)

// Company struct defines company table's schema
type Company struct {
	gorm.Model
	ID 	  string `gorm:"primaryKey"`
	Name      string
	Email     string
	Employees int
	Projects  []Project `gorm:"foreignKey:CID"`
}

// Project struct defines project table's schema
type Project struct {
	gorm.Model
	ID 	  string `gorm:"primaryKey"`
	Name      string
	Domain    string
	Employees int
	Duration  string
	CID       string
}