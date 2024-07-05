package config

import (
	"fmt"

	"github.com/suhailmshaik/go-projects-companies-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	
	// gorm.Open: This function is used to initialize a new GORM DB instance. It requires two arguments: a dialect and a configuration. The dialect specifies the type of database you are connecting to (in this case, PostgreSQL), and the configuration allows for custom settings for the database connection.


	// postgres.Open: This is a function provided by the GORM PostgreSQL driver. It constructs a DSN (Data Source Name) string for connecting to a PostgreSQL database. The DSN string contains the credentials, host, port, and database name required to establish the connection.

	// &gorm.Config{}: This is the configuration object for GORM. By passing an empty gorm.Config{}, you're using the default configuration. However, this object can be customized to change GORM's behavior, such as enabling logging, specifying naming strategies, and more.

	// use your database credentials
	dsn := "host=localhost user=postgres password=password dbname=go-projects-companies-api port=5432 sslmode=disable"


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	// Automated Schema Migration: AutoMigrate automatically updates the database table schema to match the structure of the specified model(s). In this case, it's adjusting the schema for the `Company` and `Projects` models defined in the models package.
	db.AutoMigrate(&models.Company{}, &models.Project{})
	fmt.Println("Database connected")

	DB = db
}