package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-projects-companies-api/config"
	"github.com/suhailmshaik/go-projects-companies-api/routes"
)

// ==============================================================================================================================
// Execution starts from the main function, this application can create `Projects` and `Companies` and store them in a PostgreSQL
// ==============================================================================================================================

func main() {

	// Initializing a GIN router or a Web Server
	router:=gin.New()

	// Connecting to the database
	config.Connect()	

	// Passing router or web server into user defined package of routes
	routes.CompanyRouter(router)
	routes.ProjectRouter(router)

	// Running the server on port 5000
	router.Run(":5000")
}