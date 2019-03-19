package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/utils"
)

var DeleteSurveysQuestionsRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)
	surveyID := vars["survey_id"]

	// Delete all records from "surveys_factors" table for the specific survey.
	if _, err := database.DBSQL.Exec("DELETE FROM surveys_questions WHERE survey_id = $1;", surveyID); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The records successfully deleted.")
}

var CreateSurveysQuestionsRelationship  = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)
	surveyID := vars["survey_id"]

	type RequestBody struct {
		Questions []int `json:"questions"`
	}

	// Initialize the variable called "requestBody" and assign "RequestBody" struct to the variable.
	requestBody := RequestBody{}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&result".
	if err := decoder.Decode(&requestBody); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Parse the array which was send in JSON body from frontend side.
	for i := 0; i < len(requestBody.Questions); i++ {
		// Execute the SQL query to create new record in the "surveys_factors" table.
		if _, err := database.DBSQL.Exec("INSERT INTO surveys_questions (survey_id, question_id) VALUES ($1, $2);", surveyID, requestBody.Questions[i]); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "All new records successfully created.")
}
