package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strings"
)

var CreateSingleSurveyFactorRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a "SurveyFactorRelationship" struct.
	surveyFactorRelationship := models.SurveyFactorRelationship{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&surveyFactorRelationship".
	if err := decoder.Decode(&surveyFactorRelationship); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&surveyFactorRelationship).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var CreateMultipleSurveyFactorRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize "RequestBody" struct.
	type RequestBody struct {
		SurveyID string `json:"survey_id"`
		Factors []int `json:"factors"`
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
	sqlStatement.WriteString("INSERT INTO SURVEYS_FACTORS_RELATIONSHIP (SURVEY_ID, FACTOR_ID) SELECT '")
	sqlStatement.WriteString(requestBody.SurveyID)
	sqlStatement.WriteString("' SURVEY_ID, FACTOR_ID FROM UNNEST(ARRAY[")
	sqlStatement.WriteString(utils.ConvertIntArrayToString(requestBody.Factors))
	sqlStatement.WriteString("]) FACTOR_ID")

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(sqlStatement.String()); if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var DeleteSingleSurveyFactorRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Variable has been initialized by assigning it a unique identifier of factor.
		factorIdentifier := keys.Get("factor_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 && len(factorIdentifier) != 0 {
			// Variable has been initialized by assigning it a "SurveyFactorRelationship" struct.
			surveyFactorRelationship := models.SurveyFactorRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("SURVEY_ID = ? AND FACTOR_ID = ?", surveyIdentifier, factorIdentifier).Delete(&surveyFactorRelationship).Error; err != nil {
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

var DeleteMultipleSurveyFactorRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 {
			// Variable has been initialized by assigning it a "SurveyFactorRelationship" struct.
			surveyFactorRelationship := models.SurveyFactorRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("SURVEY_ID = ?", surveyIdentifier).Delete(&surveyFactorRelationship).Error; err != nil {
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

var GetBetaSurveysFactorsRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Variable has been initialized by assigning it a array.
	var factors []models.BetaFactor

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 {
			// Make SQL query by "database/sql" package.
			rows, err := database.DBSQL.Query("SELECT ID, NAME FROM SURVEYS_FACTORS_RELATIONSHIP INNER JOIN FACTORS ON SURVEYS_FACTORS_RELATIONSHIP.FACTOR_ID = FACTORS.ID WHERE SURVEY_ID = $1", surveyIdentifier); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function.
			defer rows.Close()

			// Parse the result set of the first SQL query.
			for rows.Next() {
				// Variable has been initialized by assigning it to a "BetaFactor" struct.
				var factor models.BetaFactor

				// Call "Scan()" function on the result set of the SQL query.
				if err := rows.Scan(&factor.ID, &factor.Name); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Set all factors to the final array.
				factors = append(factors, factor)
			}
		} else {
			utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
			return
		}
	} else {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
		return
	}

	// Check the length of the array.
	if len(factors) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, factors)
}
