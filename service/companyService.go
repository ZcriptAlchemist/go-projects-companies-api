package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/suhailmshaik/go-projects-companies-api/dao"
	"github.com/suhailmshaik/go-projects-companies-api/models"
	"github.com/suhailmshaik/go-projects-companies-api/utils"
)

// ==============================================================
// Creates a Company and adds unique CID or Company ID to it
// ==============================================================

func CreateCompany(company *models.Company) error {

	// generate a 4 digit unique CID - Customer ID
	cid := utils.GenerateCID()

	company.ID = fmt.Sprintf("CID%v", cid)

	err := dao.SaveCompany(company)

	if err != nil {
		log.Printf("error in service layer %v \n", err)
		return err
	}

	return nil
}

// ========================================
// Fetches all company details from the DB
// ========================================

func FetchCompanyDetails() ([]models.Company, error) {

	companyList, err := dao.FindCompanies()

	if err != nil {
		log.Println("error in service layer: ",err)
		return companyList, errors.New("companies not found, error in service layer",)
	}

	return companyList, nil
}

// =============================================
// Fetches company details from the DB using ID
// =============================================

func FetchCompanyDetailsByID(id string) (models.Company, error) {
	
	company, err := dao.FindCompanyByID(id)

	if err != nil {
		return company, err
	}
	return company, nil
}

// ========================================
// Deletes all company details from the DB
// ========================================

func DeleteAllCompanies( ) error {

	err := dao.RemoveAllCompanies()

	if err != nil {
		log.Printf("error in service layer: %v \n", err)
		return err
	}

	return nil
}