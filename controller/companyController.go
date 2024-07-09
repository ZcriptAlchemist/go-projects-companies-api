package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-projects-companies-api/models"
	"github.com/suhailmshaik/go-projects-companies-api/service"
)

// =================
// Create a company
// =================

func CreateCompany(context *gin.Context) {

	var company models.Company

	err := context.BindJSON(&company)

	if err != nil{
		context.JSON(400, gin.H{"message:": "Invalid request"})
	}

	// Passing the company address to the service layer to generate unique CID

	err = service.CreateCompany(&company)

	if err != nil{
		context.JSON(400, gin.H{"message:": err})
		return
	} 

	context.JSON(201,"successfully created a company")
}

// ==================
// Get all companies
// ==================

func GetAllCompanies(context *gin.Context) {
	
	response, err := service.FetchCompanyDetails()

	if err != nil {
		context.JSON(400, gin.H{"message:": err})
		log.Printf("error from controller layer %v", err)
	} else {
		context.JSON(200 ,response)
	}

}

// ==================
// Get company by ID
// ==================

func GetCompanyByID(context *gin.Context) {

	id := context.Param("id")

	response, err := service.FetchCompanyDetailsByID(id)

	if err != nil {
		context.JSON(400, gin.H{"message:": err})
	} else {
		context.JSON(200,response)
	}

}

// =====================
// Delete all companies
// =====================

func DeleteAllCompanies(context *gin.Context) {

	err := service.DeleteAllCompanies()

	if err != nil {
		context.JSON(400, gin.H{"message:": err})
		log.Printf("error from controller layer %v", err)
	} else {
		context.JSON(204,"successfully deleted all companies")
	}

}