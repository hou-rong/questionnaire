package controllers

import (
	"log"
	"questionnaire/database"
	"time"
)

var Tracker = func() {
	log.Println("CRONTAB scheduler job/script (\"TRACKER\") has been started.")

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec("CALL tracker($1)", time.Now().Format("2006-01-02 15:04:05")); if err != nil {
		log.Fatal(err)
		return
	}
}
