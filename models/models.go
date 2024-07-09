package models

import (
	"gorm.io/gorm"
)

// ==============================================
// Company struct defines company table's schema
// ==============================================
type Company struct {
	gorm.Model
	ID 	  string `gorm:"primaryKey"`
	Name      string
	Email     string
	Employees int
	Projects  []Project `gorm:"foreignKey:CID"`
}

// ==============================================
// Project struct defines project table's schema
// ==============================================
type Project struct {
	gorm.Model
	ID 	  string `gorm:"primaryKey"`
	Name      string
	Domain    string
	Employees int
	Duration  string
	CID       string
}
// ====================================================
// custom struct to fetch projects with company details
// ====================================================
type ProjectWithCompany struct {
    CompanyName      string `gorm:"column:name"`
    CompanyID        string `gorm:"column:id"`
    CompanyEmail     string `gorm:"column:email"`
    ProjectID        string `gorm:"column:id"`
    ProjectName      string `gorm:"column:name"`
    ProjectDomain    string `gorm:"column:domain"`
    ProjectEmployees int    `gorm:"column:employees"`
    ProjectDuration  string `gorm:"column:duration"`
}