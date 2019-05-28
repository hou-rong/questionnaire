package controllers

import (
	"github.com/lib/pq"
	"log"
	"os"
	"questionnaire/database"
	"questionnaire/models"
	"strconv"
	"strings"
)

var History = func() {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("CRONTAB scheduler job/script (\"HISTORY\") has been started.")

	// Initialize the variable.
	var surveyIdentifiers pq.StringArray

	// Make SQL query to remote PostgreSQL database.
	if err := database.DBSQL.QueryRow("SELECT ARRAY_AGG(ID) AS SURVEY_IDENTIFIERS FROM SURVEYS WHERE CONDITION = 2 AND START_PERIOD IS NOT NULL AND END_PERIOD IS NOT NULL AND CURRENT_TIMESTAMP BETWEEN START_PERIOD AND END_PERIOD").Scan(&surveyIdentifiers); err != nil {
		logger.Println(err)
		return
	}

	// Check the length of the array.
	if len(surveyIdentifiers) != 0 {
		// Make SQL query to remote Oracle database.
		structures, err := database.OracleDB.Query(`SELECT DISTINCT 
        	O.ORGANIZATION_ID,
            O.ORGANIZATION_NAME,
            O.ORGANIZATION_RANG,
            O.PARENT_ORGANIZATION_ID
		FROM
			ONLREP.NFS_DIM_ORG_STR O
		JOIN (
			SELECT
				DISTINCT O.ORGANIZATION_ID,
				O.TREE_ORGANIZATION_ID || '\' TREE_ORGANIZATION_ID
			FROM
				ONLREP.NFS_DIM_ORG_STR O
			JOIN ONLREP.NFS_DIM_ORG_PER P ON
				P.ORGANIZATION_ID = O.ORGANIZATION_ID
			WHERE
				1 = 1
		) FO ON
		FO.TREE_ORGANIZATION_ID LIKE '%\' || O.ORGANIZATION_ID || '\%'
		WHERE
			ORGANIZATION_NAME IS NOT NULL
		AND 
			RANG1_ORGANIZATION_ID NOT IN(28825, 27624, 27626, 28833, 29033)
		ORDER BY
			ORGANIZATION_RANG,
			ORGANIZATION_ID`); if err != nil {
			logger.Println(err)
			return
		}

		// Call "Close" function to the result set of the SQL query to remote Oracle database.
		defer structures.Close()

		// Parse the result set of the SQL query to remote Oracle database.
		for structures.Next() {
			// Iterate over the array.
			for _, surveyIdentifier := range surveyIdentifiers {
				// Variable "organizationalStructureHistory" has been initialized by assigning it to a "OrganizationalStructureHistory" struct.
				var organizationalStructure models.OrganizationalStructure

				// Call "Scan()" function to the result set of the SQL query to remote Oracle database.
				if err := structures.Scan(
					&organizationalStructure.OrgStructureVersionID,
					&organizationalStructure.VerDFrom,
					&organizationalStructure.VerPrevDFrom,
					&organizationalStructure.VerPrevDTo,
					&organizationalStructure.OrganizationID,
					&organizationalStructure.ParentOrganizationID,
					&organizationalStructure.OrganizationName,
					&organizationalStructure.OrganizationRang,
					&organizationalStructure.TreeOrganizationID,
					&organizationalStructure.TreeOrganizationName,
					&organizationalStructure.Rang1OrganizationID,
					&organizationalStructure.Rang1OrganizationName,
					&organizationalStructure.CreationDate); err != nil {
					logger.Println(err)
					return
				}

				// Create "INSERT" query statement to the "ORGANIZATIONAL_STRUCTURE_HISTORY" table of remote PostgreSQL database.
				var insertQuery strings.Builder
				insertQuery.WriteString("INSERT INTO ORGANIZATIONAL_STRUCTURE_HISTORY VALUES(")
				insertQuery.WriteString(strconv.Itoa(organizationalStructure.OrgStructureVersionID))
				insertQuery.WriteString(", '")
				insertQuery.WriteString(organizationalStructure.VerDFrom.Format("2006-01-02 15:04:05"))
				insertQuery.WriteString("', '")
				insertQuery.WriteString(organizationalStructure.VerPrevDFrom.Format("2006-01-02 15:04:05"))
				insertQuery.WriteString("', '")
				insertQuery.WriteString(organizationalStructure.VerPrevDTo.Format("2006-01-02 15:04:05"))
				insertQuery.WriteString("', ")
				insertQuery.WriteString(strconv.Itoa(organizationalStructure.OrganizationID))
				insertQuery.WriteString(", ")
				insertQuery.WriteString(strconv.Itoa(organizationalStructure.ParentOrganizationID))
				insertQuery.WriteString(", '")
				insertQuery.WriteString(organizationalStructure.OrganizationName)
				insertQuery.WriteString("', ")
				insertQuery.WriteString(strconv.Itoa(organizationalStructure.OrganizationRang))
				insertQuery.WriteString(", '")
				insertQuery.WriteString(organizationalStructure.TreeOrganizationID)
				insertQuery.WriteString("', '")
				insertQuery.WriteString(organizationalStructure.TreeOrganizationName)
				insertQuery.WriteString("', ")
				insertQuery.WriteString(strconv.Itoa(organizationalStructure.Rang1OrganizationID))
				insertQuery.WriteString(", '")
				insertQuery.WriteString(organizationalStructure.Rang1OrganizationName)
				insertQuery.WriteString("', '")
				insertQuery.WriteString(organizationalStructure.CreationDate.Format("2006-01-02 15:04:05"))
				insertQuery.WriteString("', '")
				insertQuery.WriteString(surveyIdentifier)
				insertQuery.WriteString("') ON CONFLICT ON CONSTRAINT ORGANIZATIONAL_STRUCTURE_HISTORY_UNIQUE_KEY DO NOTHING")

				// Make SQL query by "database/sql" package.
				_, err := database.DBSQL.Exec(insertQuery.String()); if err != nil {
					logger.Println(err)
					return
				}
			}
		}
	}
}
