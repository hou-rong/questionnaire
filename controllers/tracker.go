package controllers

import (
	"log"
	"os"
	"questionnaire/database"
	"time"
)

var Tracker = func() {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("CRONTAB scheduler job/script (\"TRACKER\") has been started.")

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec("CALL tracker($1)", time.Now().Format("2006-01-02 15:04:05")); if err != nil {
		logger.Fatal(err)
		return
	}
}
