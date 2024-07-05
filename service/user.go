package service

import (
	"fmt"

	"github.com/suhailmshaik/go-projects-companies-api/dao"
	"github.com/suhailmshaik/go-projects-companies-api/models"
	"github.com/suhailmshaik/go-projects-companies-api/utils"
)

func AddCompany(company *models.Company) {
	// generate a 4 digit unique CID - Customer ID
	cid := utils.GenerateCID()

	company.ID = fmt.Sprintf("CID%v", cid)

	dao.AddCompany(company)
}