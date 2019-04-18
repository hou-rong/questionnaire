package controllers

import (
	"github.com/lib/pq"
	"log"
	"os"
	"questionnaire/database"
	"strconv"
	"strings"
)

// Create struct called "Entry".
type Entry struct {
	SurveyIdentifier  string
	Organizations pq.Int64Array
}

var Creator = func() {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("CRONTAB scheduler job/script (\"CREATOR\") has been started.")

	/*
	Make SQL query by "database/sql" package.

	| survey_id                            | organizations                               |
	|--------------------------------------|---------------------------------------------|
	| 66c89a34-fff2-4cbc-a542-b1e956a352f3 | {27623,27734,27737,27777,27778,27781,27741} |
	| 99c89a24-fff2-4cbc-a542-b1e956a352f9 | {27623,27734}                               |
	*/
	rows, err := database.DBSQL.Query(`SELECT 
    	SURVEY_ID,
    	ARRAY_AGG (ORGANIZATION_ID) AS ORGANIZATIONS
	FROM
	     SURVEYS_ORGANIZATIONS_RELATIONSHIP 
	WHERE 
	      SURVEY_ID IN (SELECT ID FROM SURVEYS WHERE CONDITION = 2 AND BLOCKED = TRUE)
	GROUP BY
		SURVEY_ID`); if err != nil {
		logger.Println(err)
		return
	}

	// Call "Close" function.
	defer rows.Close()

	// Initialize multiple threads channel.
	channel := make(chan Entry)

	// Building worker pool.
	for worker := 1; worker <= 10; worker++ {
		go Worker(channel)
	}

	// Parse the result set of the SQL query.
	for rows.Next() {
		// Create struct called "Entry".
		var entry Entry

		// Call "Scan()" function on the result set of the SQL query.
		if err := rows.Scan(&entry.SurveyIdentifier, &entry.Organizations); err != nil {
			logger.Println(err)
			return
		}

		// Put all entries to channel.
		channel <- entry
	}
	close(channel)
}

func Worker(channel <- chan Entry) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	for entry := range channel {
		// Build first SQL statement.
		var firstStatement strings.Builder
		firstStatement.WriteString(`SELECT 
			REPLACE('''' || RTRIM(XMLAGG(XMLELEMENT(e, P_EMAIL, ',').EXTRACT('//text()')).GetClobVal(), ',') || '''', ',', '''' || ',' || '''') AS EMPLOYEES
		FROM
			NFS_DIM_ORG_PER 
		WHERE 
			P_EMAIL IS NOT NULL AND ORGANIZATION_ID IN (`)
		for i := 1; i <= len(entry.Organizations); i++ {
			firstStatement.WriteString(":value")
			firstStatement.WriteString(strconv.Itoa(i))
			if i < len(entry.Organizations) {
				firstStatement.WriteString(",")
			}
		}
		firstStatement.WriteString(")")

		// Build array of arguments for first SQL statement.
		arguments  := make([]interface{}, len(entry.Organizations))
		for i, identifier := range entry.Organizations {
			arguments [i] = identifier
		}

		// Create struct called "Employees".
		type Employees struct {
			Emails string
		}

		// Variable has been initialized by assigning it a "Employees" struct.
		var employees Employees

		/*
		Make SQL query by "go-goracle/goracle" package.

		| Employees                                                                |
		|--------------------------------------------------------------------------|
		| 'SKorzhavykh@beeline.kz','YKulikpayev@beeline.kz','SChebykin@beeline.kz' |
		*/
		if err := database.OracleDB.QueryRow(firstStatement.String(), arguments...).Scan(&employees.Emails); err != nil {
			logger.Println(err)
			return
		}

		// Build second SQL statement.
		var secondStatement strings.Builder
		secondStatement.WriteString("CALL creator('")
		secondStatement.WriteString(entry.SurveyIdentifier)
		secondStatement.WriteString("', ARRAY[")
		secondStatement.WriteString(employees.Emails)
		secondStatement.WriteString("])")

		// Make SQL query by "database/sql" package which insert multiple rows by one query.
		_, err := database.DBSQL.Exec(secondStatement.String()); if err != nil {
			logger.Println(err)
			return
		}
	}
}
