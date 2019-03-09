package database

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"questionnaire/utils"
)

var DBSQL *sql.DB

var DBGORM *gorm.DB

/*
Function name:
"ConnectPostgreSQL"

Function description:
The main task of the function is to check the connection to remote PostgreSQL database with the help of "gorm" and "database/sql" package.
*/
func ConnectPostgreSQL() {
	// The application load environment variables from the ".env" file.
	err := godotenv.Load(".env")
	// If the ".env" file is not available the application will show an error message.
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// The application initialize PostgreSQL database related variables.
	databaseUser := utils.CheckEnvironmentVariable("PostgreSQL_USER")
	databasePassword := utils.CheckEnvironmentVariable("PostgreSQL_PASSWORD")
	databaseHost := utils.CheckEnvironmentVariable("PostgreSQL_HOST")
	databaseName := utils.CheckEnvironmentVariable("PostgreSQL_DATABASE_NAME")

	// The application defining the connection string for the remote PostgreSQL database with the help of the "gorm" package.
	GORMDatabaseURL:= fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", databaseHost, databaseUser, databaseName, databasePassword)

	// The application create connection pool to remote PostgreSQL database with the help of the "gorm" package.
	DBGORM, err = gorm.Open("postgres", GORMDatabaseURL)
	// If connection pool creation process was unsuccessful the application show an error message.
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// The application ping the remote PostgreSQL database with the help of "gorm" package.
	err = DBGORM.DB().Ping()
	// If ping process to the remote PostgreSQL database raise error the application show an error message.
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Enable logging mode of "gorm" package.
	DBGORM.LogMode(true)

	log.Println("RESTful web service successfully connected to remote PostgreSQL database with the help of \"gorm\" package.")

	// The application defining the connection string for the remote PostgreSQL database with the help of the "database/sql" package.
	SQLDatabaseURL:= fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", databaseUser, databasePassword, databaseHost, databaseName)

	// The application create connection pool to remote PostgreSQL database with the help of the "database/sql" package.
	DBSQL, err = sql.Open("postgres", SQLDatabaseURL)
	// If connection pool creation process was unsuccessful the application show an error message.
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// The application ping the remote PostgreSQL database with the help of "database/sql" package.
	err = DBSQL.Ping()
	// If ping process to the remote PostgreSQL database raise error the application show an error message.
	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("RESTful web service successfully connected to remote PostgreSQL database with the help of \"database/sql\" package.")
}

/*
Function name:
"GORMDisconnectPostgreSQL"

Description:
The main task of the function is to disconnect connection to remote PostgreSQL database with the help of "gorm" package.
*/
func GORMDisconnectPostgreSQL() error {
	return DBGORM.Close()
}

/*
Function name:
"SQLDisconnectPostgreSQL"

Description:
The main task of the function is to disconnect connection to remote PostgreSQL database with the help of "database/sql" package.
*/
func SQLDisconnectPostgreSQL() error {
	return DBSQL.Close()
}
