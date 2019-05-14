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

var CreateMultipleAnswer = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a "MultipleAnswer" struct.
	requestBody := models.MultipleAnswer{}

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

	// Check the length of the value in the request body element.
	if len(requestBody.OptionID) != 0 {
		sqlStatement.WriteString("INSERT INTO ANSWERS (SURVEY_ID, EMPLOYEE, QUESTION_ID, QUESTION_TEXT, OPTION_ID, OPTION_TEXT) SELECT '")
		sqlStatement.WriteString(requestBody.SurveyID)
		sqlStatement.WriteString("' SURVEY_ID, '")
		sqlStatement.WriteString(requestBody.Employee)
		sqlStatement.WriteString("' EMPLOYEE, UNNEST(ARRAY[")
		sqlStatement.WriteString(utils.ConvertIntArrayToString(requestBody.QuestionID))
		sqlStatement.WriteString("]) AS QUESTION_ID, UNNEST(ARRAY[")
		sqlStatement.WriteString(utils.ConvertStringArrayToString(requestBody.QuestionText))
		sqlStatement.WriteString("]) AS QUESTION_TEXT, UNNEST(ARRAY[")
		sqlStatement.WriteString(utils.ConvertIntArrayToString(requestBody.OptionID))
		sqlStatement.WriteString("]) AS OPTION_ID, UNNEST(ARRAY[")
		sqlStatement.WriteString(utils.ConvertStringArrayToString(requestBody.OptionText))
		sqlStatement.WriteString("]) AS OPTION_TEXT")
	} else {
		sqlStatement.WriteString("INSERT INTO ANSWERS (SURVEY_ID, EMPLOYEE, QUESTION_ID, QUESTION_TEXT, OPTION_TEXT) SELECT '")
		sqlStatement.WriteString(requestBody.SurveyID)
		sqlStatement.WriteString("' SURVEY_ID, '")
		sqlStatement.WriteString(requestBody.Employee)
		sqlStatement.WriteString("' EMPLOYEE, UNNEST(ARRAY[")
		sqlStatement.WriteString(utils.ConvertIntArrayToString(requestBody.QuestionID))
		sqlStatement.WriteString("]) AS QUESTION_ID, UNNEST(ARRAY[")
		sqlStatement.WriteString(utils.ConvertStringArrayToString(requestBody.QuestionText))
		sqlStatement.WriteString("]) AS QUESTION_TEXT, UNNEST(ARRAY[")
		sqlStatement.WriteString(utils.ConvertStringArrayToString(requestBody.OptionText))
		sqlStatement.WriteString("]) AS OPTION_TEXT")
	}

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(sqlStatement.String()); if err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}
