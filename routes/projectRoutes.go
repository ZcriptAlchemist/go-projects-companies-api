package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-projects-companies-api/controller"
)

// ===========================
// defining the project routes
// ===========================

func ProjectRouter(router *gin.Engine) {

	// Creating project
	router.POST("/project/create", controller.CreateProject)
	
	// Getting all Projects
	router.GET("/project/get/all", controller.GetAllProjects)
	
	// Getting Project by ID
	router.GET("/project/get/:id", controller.GetProjectByID)
	
	// Deleting all Projects
	router.DELETE("/project/delete/all", controller.DeleteAllProjects)
	
	// Getting Projects by Company ID
	router.GET("/project/get/company/:CID", controller.GetProjectsByCompanyID)
}