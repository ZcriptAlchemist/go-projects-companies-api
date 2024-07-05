package dao

import (
	"fmt"

	"github.com/suhailmshaik/go-projects-companies-api/config"
	"github.com/suhailmshaik/go-projects-companies-api/models"
)

// AddCompany function to add a company into the database
func AddCompany(company *models.Company) {
	config.DB.Create(company)
	fmt.Printf("successfully added %v into the DB", company.Name)
}

// FetchCompanies function to fetch all companies from the database
func FetchCompanies(companies *[]models.Company) {
	config.DB.Find(&companies)
	fmt.Printf("successfully fetched all companies from the DB")
}

// FetchCompanyByID function to fetch a company by its ID from the database
func FetchCompanyByID(company *models.Company, id string) {
	config.DB.Where("id = ?", id).Find(&company)
	fmt.Printf("successfully fetched company with id %v from the DB", id)
}

// ResetCompanies function to delete all companies from the database - reset or delete all companies options is only for testing
func ResetCompanies() {
	// DB.Exec is used to execute raw SQL queries
	config.DB.Exec("DELETE FROM companies")
	fmt.Println("successfully deleted all companies from the DB")
}