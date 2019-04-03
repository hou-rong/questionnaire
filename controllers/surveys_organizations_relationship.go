package controllers

import (
	"encoding/json"
	"github.com/lib/pq"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strconv"
	"strings"
)

var CreateSingleSurveyOrganizationRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a "SurveyOrganizationRelationship" struct.
	surveyOrganizationRelationship := models.SurveyOrganizationRelationship{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&surveyOrganizationRelationship".
	if err := decoder.Decode(&surveyOrganizationRelationship); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&surveyOrganizationRelationship).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var CreateMultipleSurveyOrganizationRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	type RequestBody struct {
		SurveyID string `json:"survey_id"`
		Organizations []int `json:"organizations"`
	}

	// Variable has been initialized by assigning it a "RequestBody" struct.
	requestBody := RequestBody{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&requestBody".
	if err := decoder.Decode(&requestBody); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Build SQL statement.
	var sqlStatement strings.Builder
	sqlStatement.WriteString("INSERT INTO SURVEYS_ORGANIZATIONS_RELATIONSHIP (SURVEY_ID, ORGANIZATION_ID) SELECT '")
	sqlStatement.WriteString(requestBody.SurveyID)
	sqlStatement.WriteString("' SURVEY_ID, ORGANIZATION_ID FROM UNNEST(ARRAY[")
	sqlStatement.WriteString(utils.ConvertIntArrayToString(requestBody.Organizations))
	sqlStatement.WriteString("]) ORGANIZATION_ID")

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(sqlStatement.String()); if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var DeleteSingleSurveyOrganizationRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Variable has been initialized by assigning it a unique identifier of organization.
		organizationIdentifier := keys.Get("organization_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 && len(organizationIdentifier) != 0 {
			// Variable has been initialized by assigning it a "SurveyOrganizationRelationship" struct.
			surveyOrganizationRelationship := models.SurveyOrganizationRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("SURVEY_ID = ? AND ORGANIZATION_ID = ?", surveyIdentifier, organizationIdentifier).Delete(&surveyOrganizationRelationship).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
			return
		}
	} else {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var DeleteMultipleSurveyOrganizationRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 {
			// Variable has been initialized by assigning it a "SurveyOrganizationRelationship" struct.
			surveyOrganizationRelationship := models.SurveyOrganizationRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("SURVEY_ID = ?", surveyIdentifier).Delete(&surveyOrganizationRelationship).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
			return
		}
	} else {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
		return
	}

	// Send JSON response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var GetBetaSurveysOrganizationsRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	keys := request.URL.Query()
	type Organization struct {
		ID int `json:"organization_id"`
		Name *string `json:"organization_name"`
	}
	var organizations []Organization
	if len(keys) > 0 {
		surveyIdentifier := keys.Get("survey_id")
		if len(surveyIdentifier) != 0 {
			var identifiers pq.Int64Array
			if err := database.DBSQL.QueryRow("SELECT ARRAY_AGG (ORGANIZATION_ID) FROM SURVEYS_ORGANIZATIONS_RELATIONSHIP WHERE SURVEY_ID = $1", surveyIdentifier).Scan(&identifiers); err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
			var stmt strings.Builder
			args := make([]interface{}, len(identifiers))
			for i, id := range identifiers {
				args[i] = id
			}
			stmt.WriteString("SELECT ORGANIZATION_ID, ORGANIZATION_NAME FROM NFS_DIM_ORG_STR WHERE ORGANIZATION_ID IN (")
			for i := 1; i <= len(identifiers); i++ {
				stmt.WriteString(":value")
				stmt.WriteString(strconv.Itoa(i))
				if i < len(identifiers) {
					stmt.WriteString(",")
				}
			}
			stmt.WriteString(")")
			if len(args) == 0 {
				utils.Response(responseWriter, http.StatusOK, nil)
				return
			}
			rows, err := database.OracleDB.Query(stmt.String(), args...); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
			defer rows.Close()
			for rows.Next() {
				var organization Organization
				if err := rows.Scan(&organization.ID, &organization.Name); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
				organizations = append(organizations, organization)
			}
		} else {
			utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
			return
		}
	} else {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
		return
	}
	if len(organizations) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}
	utils.Response(responseWriter, http.StatusOK, organizations)
}
