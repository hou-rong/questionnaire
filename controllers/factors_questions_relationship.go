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

var CreateSingleIntFactorQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a "FactorQuestionRelationship" struct.
	factorQuestionRelationship := models.FactorQuestionRelationship{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&factorQuestionRelationship".
	if err := decoder.Decode(&factorQuestionRelationship); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&factorQuestionRelationship).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var CreateMultipleIntFactorQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize "RequestBody" struct.
	type RequestBody struct {
		FactorID int `json:"factor_id"`
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
	sqlStatement.WriteString("INSERT INTO FACTORS_QUESTIONS_RELATIONSHIP (FACTOR_ID, QUESTION_ID) SELECT ")
	sqlStatement.WriteString(strconv.Itoa(requestBody.FactorID))
	sqlStatement.WriteString(" FACTOR_ID, QUESTION_ID FROM UNNEST(ARRAY[")
	sqlStatement.WriteString(utils.ConvertIntArrayToString(requestBody.Questions))
	sqlStatement.WriteString("]) QUESTION_ID")
	sqlStatement.WriteString(" ON CONFLICT ON CONSTRAINT FACTORS_QUESTIONS_RELATIONSHIP_UNIQUE_KEY DO NOTHING")

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(sqlStatement.String()); if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var CreateMultipleTextFactorQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize "RequestBody" struct.
	type RequestBody struct {
		FactorID int `json:"factor_id"`
		Questions [] struct {
			Text string `json:"question_text"`
			Widget int `json:"widget"`
			Required bool `json:"required"`
		} `json:"questions"`
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

	// Initialize several arrays which would be used in string builder.
	var textArray []string
	var widgetArray []int
	var requiredArray []bool
	var positionArray []int

	// Parse content of the request body.
	for i := 0; i < len(requestBody.Questions); i++ {
		textArray = append(textArray, requestBody.Questions[i].Text)
		widgetArray = append(widgetArray, requestBody.Questions[i].Widget)
		requiredArray = append(requiredArray, requestBody.Questions[i].Required)
		positionArray = append(positionArray, i)
	}

	// Build SQL statement.
	var sqlStatement strings.Builder
	sqlStatement.WriteString("SELECT ARRAY_AGG(RESULTS.ID) FROM factorio (")
	sqlStatement.WriteString(strconv.Itoa(requestBody.FactorID))
	sqlStatement.WriteString(", ARRAY[")
	sqlStatement.WriteString(utils.ConvertStringArrayToString(textArray))
	sqlStatement.WriteString("], ARRAY[")
	sqlStatement.WriteString(utils.ConvertIntArrayToString(widgetArray))
	sqlStatement.WriteString("], ARRAY[")
	sqlStatement.WriteString(utils.ConvertBooleanArrayToString(requiredArray))
	sqlStatement.WriteString("], ARRAY[")
	sqlStatement.WriteString(utils.ConvertIntArrayToString(positionArray))
	sqlStatement.WriteString("]) RESULTS(ID);")

	// Variable has been initialized by assigning it a array.
	var questionIdentifiers pq.Int64Array

	// Execute SQL query by "database/sql" package.
	if err := database.DBSQL.QueryRow(sqlStatement.String()).Scan(&questionIdentifiers); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and array of ids.
	utils.Response(responseWriter, http.StatusOK, map[string]pq.Int64Array{"questions": questionIdentifiers})
}

var DeleteSingleIntFactorQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of factor.
		factorIdentifier := keys.Get("factor_id")

		// Variable has been initialized by assigning it a unique identifier of question.
		questionIdentifier := keys.Get("question_id")

		// Check the value of the variables.
		if  len(factorIdentifier) != 0 && len(questionIdentifier) != 0 {
			// Variable has been initialized by assigning it a "FactorQuestionRelationship" struct.
			factorQuestionRelationship := models.FactorQuestionRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("FACTOR_ID = ? AND QUESTION_ID = ?", factorIdentifier, questionIdentifier).Delete(&factorQuestionRelationship).Error; err != nil {
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

var DeleteMultipleIntFactorQuestionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of factor.
		factorIdentifier := keys.Get("factor_id")

		// Check the value of the variables.
		if len(factorIdentifier) != 0 {
			// Variable has been initialized by assigning it a "SurveyFactorRelationship" struct.
			factorQuestionRelationship := models.FactorQuestionRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("FACTOR_ID = ?", factorIdentifier).Delete(&factorQuestionRelationship).Error; err != nil {
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

var GetBetaFactorsQuestionsRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Variable has been initialized by assigning it a array.
	var questions []models.BetaQuestion

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of factor.
		factorIdentifier := keys.Get("factor_id")

		// Check the value of the variables.
		if len(factorIdentifier) != 0 {
			// Make SQL query by "database/sql" package.
			rows, err := database.DBSQL.Query("SELECT ID, TEXT, WIDGET, REQUIRED, POSITION, CATEGORY FROM FACTORS_QUESTIONS_RELATIONSHIP INNER JOIN QUESTIONS ON FACTORS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID WHERE FACTOR_ID = $1 ORDER BY POSITION ASC",  factorIdentifier); if err != nil {
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
				if err := rows.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position, &question.Category); err != nil {
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
