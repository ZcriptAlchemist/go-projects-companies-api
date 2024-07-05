package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-projects-companies-api/controller"
)

func UserRouter(router *gin.Engine) {
	// Creating Company
	router.POST("/company/create", controller.CreateCompany)
	// Getting all Companies
	router.GET("/company/get/all", controller.GetAllCompanies)
	// Getting Company by ID
	router.GET("/company/get/:id", controller.GetCompanyByID)
	// Deleting all Companies
	router.DELETE("/company/delete/all", controller.DeleteAllCompanies)
}