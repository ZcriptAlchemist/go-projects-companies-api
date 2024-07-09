package dao

import (
	"errors"
	"log"

	"github.com/suhailmshaik/go-projects-companies-api/config"
	"github.com/suhailmshaik/go-projects-companies-api/models"
)

// =======================================================
// AddCompany function to add a company into the database
// =======================================================

func SaveCompany(company *models.Company) error {

	err := config.DB.Create(&company).Error

	if err != nil{
		return errors.New("unable to create company")
	}

	return nil
}

// =================================================================
// FetchCompanies function to fetch all companies from the database
// =================================================================

func FindCompanies() ([]models.Company, error) {

	var companies []models.Company

	result := config.DB.Find(&companies)

	if result.Error != nil {
		log.Printf("error in DAO layer %v \n", result.Error)
		return companies, errors.New("error fetching companies in dao layer")
	}
	
	return companies, nil
}

// =========================================================================
// FetchCompanyByID function to fetch a company by its ID from the database
// =========================================================================

func FindCompanyByID(id string) (models.Company, error) {
	
	var company models.Company

	err := config.DB.Where("id = ?", id).Find(&company)
	
	if err != nil {
		log.Printf("error in DAO layer %v \n", err)
		return company, errors.New("company not found, error in dao layer")
	}

	return company, nil
}

// ==============================================================================================================================
// ResetCompanies function to delete all companies from the database - reset or delete all companies options is only for testing
// ==============================================================================================================================

func RemoveAllCompanies() error {
	// DB.Exec is used to execute raw SQL queries
	err := config.DB.Exec("DELETE FROM companies")

	if err != nil {
		return err.Error
	}

	return nil
}