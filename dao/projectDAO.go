package dao

import (
	"errors"
	"log"

	"github.com/suhailmshaik/go-projects-companies-api/config"
	"github.com/suhailmshaik/go-projects-companies-api/models"
)

// ===================
// Adds Project to DB
// ===================

func SaveProject(project *models.Project) error {
	
	err := config.DB.Create(project)
	
	if err != nil {
		return errors.New("unable to create project")
	}

	return nil
}

// ================================================================
// FetchCompanies function to fetch all projects from the database
// ================================================================

func FindProjects() ([]models.Project, error) {

	var projects []models.Project 
	
	result := config.DB.Find(&projects)

	if result.Error != nil {
		log.Printf("error in DAO layer %v \n", result.Error)
		return projects, errors.New("error fetching projects in dao layer")
	}
	
	return projects, nil
}

// =========================================================================
// FetchCompanyByID function to fetch a project by its ID from the database
// =========================================================================

func FindProjectByID(id string) (models.Project, error) {
	
	var project models.Project
	
	result := config.DB.Where("id = ?", id).Find(&project)
	
	if result.Error != nil {
		log.Printf("error in DAO layer %v \n", result.Error)
		return project, errors.New("error fetching project with ID in dao layer")
	}

	return project, nil
}

// ===========================================================================================================================
// ResetProject function to delete all projects from the database - reset or delete all companies options is only for testing
// ===========================================================================================================================

func RemoveAllProjects() error {
	// DB.Exec is used to execute raw SQL queries
	err := config.DB.Exec("DELETE FROM projects")
	
	if err != nil {
		return errors.New("unable to delete all projects - DAO")
	}

	return nil
}

// ========================================================================================
// FetchProjectsByCompanyID function to fetch all projects by company ID from the database
// ========================================================================================

func FindProjectsByCompanyID(companyID string) ([]models.ProjectWithCompany, error) {

	var projectsUnderCompany []models.ProjectWithCompany

	result := config.DB.Table("projects").
		Select("companies.name, companies.id, companies.email, projects.id, projects.name, projects.domain, projects.employees, projects.duration").Joins("JOIN companies ON companies.id = projects.c_id").Where("projects.c_id = ?", companyID).Find(&projectsUnderCompany)

	if result.Error != nil {
		log.Printf("error in DAO layer %v \n", result.Error)
		return projectsUnderCompany, errors.New("error fetching project with ID in dao layer")
	}

	return projectsUnderCompany, nil

}