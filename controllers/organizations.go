package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"sort"
)

var GetOrganizations = func(responseWriter http.ResponseWriter, request *http.Request) {
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	rows, err := database.OracleDB.Query("SELECT ORGANIZATION_ID, ORGANIZATION_NAME, ORGANIZATION_RANG, PARENT_ORGANIZATION_ID FROM NFS_DIM_ORG_STR WHERE ORGANIZATION_NAME IS NOT NULL AND RANG1_ORGANIZATION_ID NOT IN (28825, 27624, 27626, 28833, 29033) ORDER BY ORGANIZATION_RANG, ORGANIZATION_ID"); if err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	structure := map[int]*models.Organization{}
	for rows.Next() {
		organization := &models.Organization{}
		var parentOrganizationIdentifier sql.NullInt64
		if err = rows.Scan(&organization.ID, &organization.Name, &organization.Rang, &parentOrganizationIdentifier); err != nil {
			logger.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}
		if parentOrganizationIdentifier.Valid {
			if parentOrganization, ok := structure[int(parentOrganizationIdentifier.Int64)]; ok {
				parentOrganization.Children = append(parentOrganization.Children, organization)
			} else {
				structure[int(parentOrganizationIdentifier.Int64)] = &models.Organization{ID: int(parentOrganizationIdentifier.Int64)}
				structure[int(parentOrganizationIdentifier.Int64)].Children = append(structure[int(parentOrganizationIdentifier.Int64)].Children, organization)
			}
		}
		if _, ok := structure[organization.ID]; ok {
			structure[organization.ID].Name = organization.Name
			structure[organization.ID].Rang = organization.Rang
			continue
		}
		structure[organization.ID] = organization
	}
	var identifiers []int
	for item := range structure {
		identifiers = append(identifiers, item)
	}
	sort.Ints(identifiers)
	var organizations []models.Organization
	for _, identifier := range identifiers {
		if len(structure[identifier].Children) > 0 && structure[identifier].Rang == 1 {
			organizations = append(organizations, *structure[identifier])
		}
	}
	if len(organizations) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}
	utils.Response(responseWriter, http.StatusOK, organizations)
}
