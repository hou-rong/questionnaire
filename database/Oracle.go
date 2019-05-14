package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "gopkg.in/goracle.v2"
	"log"
	"os"
	"questionnaire/utils"
)

var OracleDB *sql.DB

/*
Function name:
"ConnectOracle"

Function description:
The main task of the function is to check the connection to remote Oracle database with the help of "gopkg.in/goracle.v2" package.
*/
func ConnectOracle() {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// The application load environment variables from the ".env" file.
	err := godotenv.Load(".env")
	// If the ".env" file is not available the application will show an error message.
	if err != nil {
		logger.Println(err)
		panic(err)
	}

	// The application initialize Oracle database related variables.
	databaseUser := utils.CheckEnvironmentVariable("ORACLE_USER")
	databasePassword := utils.CheckEnvironmentVariable("ORACLE_PASSWORD")
	databaseHost := utils.CheckEnvironmentVariable("ORACLE_HOST")
	databasePort := utils.CheckEnvironmentVariable("ORACLE_PORT")
	databaseName := utils.CheckEnvironmentVariable("ORACLE_DATABASE_NAME")

	// The application defining the connection string for the remote Oracle database with the help of the "gopkg.in/goracle.v2" package.
	databaseURL:= fmt.Sprintf("%s/%s@%s:%s/%s", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	// The application create connection pool to remote Oracle database with the help of the "gopkg.in/goracle.v2" package.
	OracleDB, err = sql.Open("goracle", databaseURL)
	// If connection pool creation process was unsuccessful the application show an error message.
	if err != nil {
		logger.Println(err)
		panic(err)
	}

	// The application ping the remote PostgreSQL database with the help of "gopkg.in/goracle.v2" package.
	err = OracleDB.Ping()
	// If ping process to the remote PostgreSQL database raise error the application show an error message.
	if err != nil {
		logger.Println(err)
		panic(err)
	}

	logger.Println("Web service successfully connected to remote ORACLE database with the help of \"gopkg.in/goracle.v2\" package.")
}

/*
Function name:
"DisconnectOracle"

Description:
The main task of the function is to disconnect connection to remote Oracle database with the help of "database/sql" package.
*/
func DisconnectOracle() error {
	return OracleDB.Close()
}
