package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-projects-companies-api/dao"
	"github.com/suhailmshaik/go-projects-companies-api/models"
	"github.com/suhailmshaik/go-projects-companies-api/service"
)

// create a company
func CreateCompany(context *gin.Context) {
	var company models.Company
	context.BindJSON(&company)
	// passing the company address to the service layer to generate unique CID
	service.AddCompany(&company)
	context.JSON(201,&company)
}

// get all companies
func GetAllCompanies(context *gin.Context) {
	var companies []models.Company
	context.BindJSON(&companies)
	dao.FetchCompanies(&companies)
	context.JSON(200,&companies)
}

// get company by ID
func GetCompanyByID(context *gin.Context) {
	var company models.Company
	id := context.Param("id")
	dao.FetchCompanyByID(&company, id)
	context.JSON(200,&company)
}

// Delete all companies
func DeleteAllCompanies(context *gin.Context) {
	dao.ResetCompanies()
	context.JSON(204,"successfully deleted all companies")
}