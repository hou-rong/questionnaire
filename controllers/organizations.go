package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"sort"
)

var GetOrganizations = func(responseWriter http.ResponseWriter, request *http.Request) {
	rows, err := database.OracleDB.Query(`SELECT
	  	ORGANIZATION_ID,
		ORGANIZATION_NAME,
       	ORGANIZATION_RANG,
       	PARENT_ORGANIZATION_ID
	FROM
	    DMP_ORG_STR
	WHERE
	    ORGANIZATION_NAME IS NOT NULL
	ORDER BY
		ORGANIZATION_RANG, ORGANIZATION_ID`)

	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	defer rows.Close()

	structure := map[int]*models.Organization{}

	for rows.Next() {
		organization := &models.Organization{}

		var parentOrganizationID sql.NullInt64

		if err = rows.Scan(&organization.ID, &organization.Name, &organization.Rang, &parentOrganizationID); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		if parentOrganizationID.Valid {
			if parentOrg, ok := structure[int(parentOrganizationID.Int64)]; ok {
				parentOrg.Children = append(parentOrg.Children, organization)
			} else {
				structure[int(parentOrganizationID.Int64)] = &models.Organization{ID: int(parentOrganizationID.Int64)}
				structure[int(parentOrganizationID.Int64)].Children = append(structure[int(parentOrganizationID.Int64)].Children, organization)
			}
		}

		if _, ok := structure[organization.ID]; ok {
			structure[organization.ID].Name = organization.Name
			structure[organization.ID].Rang = organization.Rang
			continue
		}

		structure[organization.ID] = organization
	}

	var IDs []int

	for i := range structure {
		IDs = append(IDs, i)
	}

	sort.Ints(IDs)

	var organizations []models.Organization

	for _, ID := range IDs {
		if len(structure[ID].Children) > 0 && structure[ID].Rang == 1 {
			organizations = append(organizations, *structure[ID])
		}
	}

	utils.Response(responseWriter, http.StatusOK, organizations)
}
