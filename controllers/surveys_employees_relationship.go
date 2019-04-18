package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strings"
)

var CreateSingleSurveyEmployeeRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a "SurveyEmployeeRelationship" struct.
	surveyEmployeeRelationship := models.SurveyEmployeeRelationship{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&surveyFactorRelationship".
	if err := decoder.Decode(&surveyEmployeeRelationship); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&surveyEmployeeRelationship).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var CreateMultipleSurveyEmployeeRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize "RequestBody" struct.
	type RequestBody struct {
		SurveyID string `json:"survey_id"`
		Employees []string `json:"employees"`
	}

	// Variable has been initialized by assigning it a "RequestBody" struct.
	requestBody := RequestBody{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&requestBody".
	if err := decoder.Decode(&requestBody); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Build SQL statement.
	var sqlStatement strings.Builder
	sqlStatement.WriteString("INSERT INTO SURVEYS_EMPLOYEES_RELATIONSHIP (SURVEY_ID, EMPLOYEE) SELECT '")
	sqlStatement.WriteString(requestBody.SurveyID)
	sqlStatement.WriteString("' SURVEY_ID, EMPLOYEE FROM UNNEST(ARRAY[")
	sqlStatement.WriteString(utils.ConvertStringArrayToString(requestBody.Employees))
	sqlStatement.WriteString("]) EMPLOYEE")

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(sqlStatement.String()); if err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var DeleteSingleSurveyEmployeeRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Variable has been initialized by assigning it a unique identifier of factor.
		employee := keys.Get("employee")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 && len(employee) != 0 {
			// Variable has been initialized by assigning it a "SurveyEmployeeRelationship" struct.
			surveyEmployeeRelationship := models.SurveyEmployeeRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("SURVEY_ID = ? AND EMPLOYEE = ?", surveyIdentifier, employee).Delete(&surveyEmployeeRelationship).Error; err != nil {
				logger.Println(err)
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

var DeleteMultipleSurveyEmployeeRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 {
			// Variable has been initialized by assigning it a "SurveyEmployeeRelationship" struct.
			surveyEmployeeRelationship := models.SurveyEmployeeRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("SURVEY_ID = ?", surveyIdentifier).Delete(&surveyEmployeeRelationship).Error; err != nil {
				logger.Println(err)
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

var UpdateSingleSurveyEmployeeRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Variable has been initialized by assigning it a unique email of employee.
		employee := keys.Get("employee")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 && len(employee) != 0 {
			// Create struct called "RequestBody".
			type RequestBody struct {
				Status bool `json:"status"`
			}

			// Variable has been initialized by assigning it a "RequestBody" struct.
			requestBody := RequestBody{}

			// "NewDecoder" returns a new decoder that reads from request body.
			// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
			decoder := json.NewDecoder(request.Body)

			// Decode reads the JSON value from its input and stores it in the value pointed to by "&requestBody".
			if err := decoder.Decode(&requestBody); err != nil {
				logger.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			// Make SQL query by "database/sql" package.
			_, err := database.DBSQL.Exec("UPDATE SURVEYS_EMPLOYEES_RELATIONSHIP SET STATUS = $1 WHERE SURVEY_ID = $2 AND EMPLOYEE = $3", requestBody.Status, surveyIdentifier, employee); if err != nil {
				logger.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Send JSON response with status code "200".
			utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
		} else {
			// Send JSON response with status code "400".
			utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
			return
		}
	} else {
		// Send JSON response with status code "400".
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
		return
	}
}
