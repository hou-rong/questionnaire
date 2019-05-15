package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mileusna/crontab"
	"log"
	"net/http"
	"os"
	"questionnaire/controllers"
	"questionnaire/database"
	"questionnaire/routes"
	"questionnaire/utils"
)

func main()  {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// The application load environment variables from the ".env" file.
	err := godotenv.Load(".env")
	// If the ".env" file is not available the application will show an error message.
	if err != nil {
		panic(err)
	}

	// The application check the connection to remote PostgreSQL database with the help of "gorm" and "database/sql" packages.
	database.ConnectPostgreSQL()
	defer database.SQLDisconnectPostgreSQL()
	defer database.GORMDisconnectPostgreSQL()

	// The application check the connection to remote Oracle database with the help of "gopkg.in/goracle.v2" packages.
	database.ConnectOracle()
	defer database.DisconnectOracle()

	// The application create the "Gorilla Mux" router.
	router := mux.NewRouter()

	// The application define CORS (cross-origin resource sharing) settings by "Gorilla Handlers" library.
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// "StrictSlash" is a function of the "Gorilla Mux" library by which the application can order the router to redirect URL routes with trailing slashes to those without them.
	router.StrictSlash(true)

	// The application define the list of all available URLs.
	routes.Handle(router)

	// Defining the application port for listening the HTTP requests.
	port := utils.CheckEnvironmentVariable("APPLICATION_PORT")

	logger.Printf("Web service is running on %s port.", port)

	err = crontab.New().AddJob("* * * * *", controllers.Tracker); if err != nil {
		logger.Fatal(err)
		return
	}

	err = crontab.New().AddJob("* * * * *", controllers.Creator); if err != nil {
		logger.Fatal(err)
		return
	}

	err = crontab.New().AddJob("* * * * *", controllers.History); if err != nil {
		logger.Fatal(err)
		return
	}

	// The application is starting to listen and serving the web service with CORS options.
	logger.Fatal(http.ListenAndServe(":" + port, handlers.CORS(headers, methods, origins)(router)))
}