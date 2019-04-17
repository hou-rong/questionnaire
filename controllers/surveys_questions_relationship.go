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

var CreateSingleSurveyQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a "SurveyQuestionRelationship" struct.
	surveyQuestionRelationship := models.SurveyQuestionRelationship{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&surveyQuestionRelationship".
	if err := decoder.Decode(&surveyQuestionRelationship); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&surveyQuestionRelationship).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var CreateMultipleSurveyQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	type RequestBody struct {
		SurveyID string `json:"survey_id"`
		Questions []int `json:"questions"`
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
	sqlStatement.WriteString("INSERT INTO SURVEYS_QUESTIONS_RELATIONSHIP (SURVEY_ID, QUESTION_ID) SELECT '")
	sqlStatement.WriteString(requestBody.SurveyID)
	sqlStatement.WriteString("' SURVEY_ID, QUESTION_ID FROM UNNEST(ARRAY[")
	sqlStatement.WriteString(utils.ConvertIntArrayToString(requestBody.Questions))
	sqlStatement.WriteString("]) QUESTION_ID")
	sqlStatement.WriteString(" ON CONFLICT ON CONSTRAINT SURVEYS_QUESTIONS_RELATIONSHIP_UNIQUE_KEY DO NOTHING")

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(sqlStatement.String()); if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var DeleteSingleSurveyQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Variable has been initialized by assigning it a unique identifier of question.
		questionIdentifier := keys.Get("question_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 && len(questionIdentifier) != 0 {
			// Variable has been initialized by assigning it a "SurveyQuestionRelationship" struct.
			surveyQuestionRelationship := models.SurveyQuestionRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("SURVEY_ID = ? AND QUESTION_ID = ?", surveyIdentifier, questionIdentifier).Delete(&surveyQuestionRelationship).Error; err != nil {
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

var DeleteMultipleSurveyQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 {
			// Variable has been initialized by assigning it a "SurveyQuestionRelationship" struct.
			surveyQuestionRelationship := models.SurveyQuestionRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("SURVEY_ID = ?", surveyIdentifier).Delete(&surveyQuestionRelationship).Error; err != nil {
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

var GetBetaSurveysQuestionsRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Variable has been initialized by assigning it a array.
	var questions []models.BetaQuestion

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Check the value of the variables.
		if len(surveyIdentifier) != 0 {
			// Make SQL query by "database/sql" package.
			rows, err := database.DBSQL.Query("SELECT ID, NAME, WIDGET FROM SURVEYS_QUESTIONS_RELATIONSHIP INNER JOIN QUESTIONS ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID WHERE SURVEY_ID = $1", surveyIdentifier); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function.
			defer rows.Close()

			// Parse the result set of the first SQL query.
			for rows.Next() {
				// Variable has been initialized by assigning it to a "BetaQuestion" struct.
				var question models.BetaQuestion

				// Call "Scan()" function on the result set of the SQL query.
				if err := rows.Scan(&question.ID, &question.Text, &question.Widget); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Set all questions to the final array.
				questions = append(questions, question)
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
	if len(questions) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, questions)
}

var CheckSurveyQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize "Result" struct.
	type Result struct {
		ID string `gorm:"primary_key" json:"survey_id"`
		Name *string `json:"survey_name"`
		Email *string `json:"email"`
	}

	// Variable has been initialized by assigning it a array.
	var results []Result

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of question.
		questionIdentifier := keys.Get("question_id")

		// Check the value of the variables.
		if len(questionIdentifier) != 0 {
			// Execute SQL query by "database/sql" package.
			rows, err := database.DBSQL.Query(`SELECT
       			ID,
       			NAME,
       			EMAIL
			FROM SURVEYS
			INNER JOIN SURVEYS_QUESTIONS_RELATIONSHIP
			ON SURVEYS.ID = SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID
			WHERE SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = $1
			AND SURVEYS.CONDITION = 2
			AND SURVEYS.BLOCKED = true
			GROUP BY ID;`, questionIdentifier); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function.
			defer rows.Close()

			// Parse the result set of the SQL query.
			for rows.Next() {
				var result Result

				// Call "Scan()" function to the result set of the second SQL query.
				if err := rows.Scan(&result.ID, &result.Name, &result.Email); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Append result to the final array.
				results = append(results, result)
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
	if len(results) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, results)
}
