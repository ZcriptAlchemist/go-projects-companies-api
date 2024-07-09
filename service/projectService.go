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
// Creates a project and adds unique PID or Project ID to it
// ==============================================================

func CreateProject(project *models.Project) error {
	// Generate a 4 digit unique PID - Project ID
	pid := utils.GenerateCID()

	project.ID = fmt.Sprintf("PID%v", pid)

	err := dao.SaveProject(project)

	if err != nil {
		log.Printf("error in service layer %v \n", err)
		return err
	}

	return nil
}

// ======================================================================
// Gets all project details through the database by calling DAO function
// ======================================================================

func FetchProjectsDetails () ([]models.Project, error) {

	projectsList, err := dao.FindProjects()

	if err != nil {
		log.Println("error in service layer: ",err)
		return projectsList, errors.New("projects not found, error in service layer",)
	}

	return projectsList, nil
}

// =========================================================
// Passes id to DAO layer to get project using id attribute
// =========================================================

func FetchProjectDetailsByID (id string) (models.Project, error) {
	
	project, err := dao.FindProjectByID(id)

	if err != nil {
		log.Println("error in service layer: ",err)
		return project, errors.New("project not found, error in service layer")
	}

	return project, nil
}

// ==================================================
// Deletes all companies in the DB through DAO layer
// ==================================================

func DeleteAllProjects () error {
		err := dao.RemoveAllProjects()

	if err != nil {
		log.Printf("error in service layer: %v \n", err)
		return err
	}

	return nil
}

// ===========================================================
// Fetches All projects under a company using CID through DAO
// ===========================================================

func FetchProjectsUnderCompany (id string) ([]models.ProjectWithCompany, error) {

	projectsUnderCompany, err := dao.FindProjectsByCompanyID(id)

	if err != nil {
		log.Println("error in service layer: ",err)
		return projectsUnderCompany, errors.New("invalid CID, error in service layer")
	}

	return projectsUnderCompany, nil
}