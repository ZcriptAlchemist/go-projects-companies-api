package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-projects-companies-api/models"
	"github.com/suhailmshaik/go-projects-companies-api/service"
)

// =================================
// Function to create a new project
// =================================

func CreateProject(context *gin.Context) {

	var project models.Project
	
	err := context.BindJSON(&project)

	if err != nil{
		context.JSON(400, gin.H{"message:": "Invalid request"})
	}

	err = service.CreateProject(&project)

	if err != nil{
		context.JSON(400, gin.H{"message:": err})
	} else {
		context.JSON(201, &project)
	}
	
}

// =================
// Get all projects
// =================

func GetAllProjects(context *gin.Context) {
	
	response, err := service.FetchProjectsDetails()
	
	if err != nil {
		context.JSON(400, gin.H{"message:": err})
		log.Printf("error from controller layer %v", err)
	} else {
		context.JSON(200 ,response)
	}

}

// ==================
// Get project by ID
// ==================


func GetProjectByID(context *gin.Context) {
	
	id := context.Param("id")

	response, err :=  service.FetchProjectDetailsByID(id)

	if err != nil {
		context.JSON(400, gin.H{"message:": err})
		log.Printf("error from controller layer %v", err)
	} else {
		context.JSON(200,response)
	}

}

// ====================
// Delete all projects
// ====================

func DeleteAllProjects(context *gin.Context) {

	err := service.DeleteAllProjects()

	if err != nil {
		context.JSON(400, gin.H{"message:": err})
		log.Printf("error from controller layer %v", err)
	} else {
		context.JSON(204,"successfully deleted all projects from the DB")
	}

}

// ==============================================
// Get projects under a company using company ID
// ==============================================

func GetProjectsByCompanyID(context *gin.Context) {
	
	companyID := context.Param("CID")
	
	response, err := service.FetchProjectsUnderCompany(companyID)

	if err != nil {
		context.JSON(400, gin.H{"message:": err})
		log.Printf("error from controller layer %v", err)
	} else {
		context.JSON(http.StatusOK, response)
	}

}