package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-projects-companies-api/config"
	"github.com/suhailmshaik/go-projects-companies-api/routes"
)

// execution starts from the main function, this application can create `Projects` and `Companies` and store them in a PostgreSQL database.
func main() {
	fmt.Println("welcome to go-projects-comapnies-api")
	// initializing a GIN router or a Web Server
	router:=gin.New()
	// connecting to the database
	config.Connect()	
	// passing router or web server into user defined package of routes
	routes.UserRouter(router)
	// running the server on port 5000
	router.Run(":5000")
}


// #file:companyController.go #file:user.go I should create a projectsController.go file and user.go under dao package, build one to many relation between company and project where a company can have multiple projects and I can fetch the data of a company along with its projects and vice versa #file:db.go #file:main.go #file:models.go

// How can I update a specific project associated with a company in the database?