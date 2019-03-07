package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strconv"
)

var GetQuestions = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Execute the SQL query to take all questions and set it's result to the variable called "firstQuery".
	firstQuery, err := database.DBSQL.Query("SELECT * FROM questions;")
	if err != nil {
		log.Println(err)
		return
	}

	// Call "Close" function.
	defer firstQuery.Close()

	// Initialize the variable called "questions" and assign an array to the variable.
	var questions []models.Question

	// Parse the result set of the first SQL query.
	for firstQuery.Next() {
		// Initialize the variable called "question" and assign an "Question" struct to the variable.
		var question models.Question

		// Call "Scan()" function on the result set of the first SQL query.
		if err := firstQuery.Scan(&question.ID, &question.Text); err != nil {
			log.Println(err)
			return
		}

		// Execute the SQL query to take information about the widget of the specific question and set it's result to the variable called "secondQuery".
		secondQuery := database.DBSQL.QueryRow(`SELECT 
			widgets.id AS widget_id,
			widgets.name AS widget_name
		FROM questions_widgets
		INNER JOIN widgets
		ON questions_widgets.widget_id = widgets.id
		WHERE questions_widgets.question_id = $1;`, question.ID)

		// Initialize the variable called "widget" and assign an "Widget" struct to the variable.
		var widget models.Widget

		// Call "Scan()" function on the result set of the second SQL query.
		err = secondQuery.Scan(&widget.ID, &widget.Name)
		if err != nil {
			log.Println(err)
			return
		}

		// Set the information about the widget of specific question to the "Widget" field of the struct "Question".
		question.Widget = widget

		// Make SQL query to take information about all options of the specific question and set it's result to the variable called "thirdQuery".
		thirdQuery, err := database.DBSQL.Query(`SELECT
			options.id AS option_id,
			options.text AS option_text
		FROM questions_options
		INNER JOIN options
		ON questions_options.option_id = options.id
		WHERE questions_options.question_id = $1;`, question.ID)
		if err != nil {
			log.Println(err)
			return
		}


		// Parse the result set of the third SQL query.
		for thirdQuery.Next() {
			// Initialize the variable called "option" and assign an "Option" struct to the variable.
			var option models.Option

			// Call "Scan()" function on the result set of the third SQL query.
			if err := thirdQuery.Scan(&option.ID, &option.Text); err != nil {
				log.Println(err)
				return
			}

			// Set the information about all options of specific question to the "Options" field of the struct "Question".
			question.Options = append(question.Options, option)
		}

		// Call "Close" function.
		thirdQuery.Close()

		// Set all questions to the final array.
		questions = append(questions, question)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, questions)
}

var CreateQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "question" and assign struct to the variable.
	question := models.Question{}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&question".
	if err := decoder.Decode(&question); err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save new record to the "questions" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&question).Error; err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "201" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "The new entry successfully created.")
}

var GetQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	questionID, err := strconv.Atoi(vars["question_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "question" and assign object with information about the specific question to the variable if the record is exist in the database.
	question := GetQuestionOr404(database.DBGORM, questionID, responseWriter, request)
	if question == nil {
		return
	}

	// Initialize the variable called "widget" and assign an "Widget" struct to the variable.
	var widget models.Widget

	// Execute the SQL query to take information about the widget of the specific question and set it's result to the variable called "firstQuery".
	firstQuery := database.DBSQL.QueryRow(`SELECT 
			widgets.id AS widget_id,
			widgets.name AS widget_name
		FROM questions_widgets
		INNER JOIN widgets
		ON questions_widgets.widget_id = widgets.id
		WHERE questions_widgets.question_id = $1;`, questionID)

	// Call "Scan()" function on the result set of the second SQL query.
	err = firstQuery.Scan(&widget.ID, &widget.Name)
	if err != nil {
		log.Println(err)
		return
	}

	// Set the information about the widget of specific question to the "Widget" field of the struct "Question".
	question.Widget = widget

	// Make SQL query to take information about all options of the specific question and set it's result to the variable called "secondQuery".
	secondQuery, _ := database.DBSQL.Query(`SELECT
			options.id AS option_id,
			options.text AS option_text
		FROM questions_options
		INNER JOIN options
		ON questions_options.option_id = options.id
		WHERE questions_options.question_id = $1;`, questionID)

	// Parse the result set of the third SQL query.
	for secondQuery.Next() {
		// Initialize the variable called "option" and assign an "Option" struct to the variable.
		var option models.Option

		// Call "Scan()" function on the result set of the third SQL query.
		if err := secondQuery.Scan(&option.ID, &option.Text); err != nil {
			log.Println(err)
			return
		}

		// Set the information about all options of specific question to the "Options" field of the struct "Question".
		question.Options = append(question.Options, option)
	}

	// Call "Close" function.
	defer secondQuery.Close()

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, question)
}

var UpdateQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	questionID, err := strconv.Atoi(vars["question_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "question" and assign object with information about the specific question to the variable if the record is exist in the database.
	question := GetQuestionOr404(database.DBGORM, questionID, responseWriter, request)
	if question == nil {
		return
	}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&question".
	if err := decoder.Decode(&question); err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save record with new information to the "questions" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&question).Error; err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully updated.")
}

var DeleteQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	questionID, err := strconv.Atoi(vars["question_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable and assign object with information about the specific option to the variable if the record is exist in the database.
	question := GetQuestionOr404(database.DBGORM, questionID, responseWriter, request)
	if question == nil {
		log.Println(err)
		return
	}

	// Delete the record from the "questions" table with the help of "gorm" package.
	if err := database.DBGORM.Delete(&question).Error; err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully deleted.")
}

func GetQuestionOr404(db *gorm.DB, questionID int, responseWriter http.ResponseWriter, request *http.Request) *models.Question {
	// Initialize the variable called "question" and assign an struct to the variable.
	question := models.Question{}

	// Find the question by id number.
	if err := db.First(&question, models.Question{ID: questionID}).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "Record not found.")
		return nil
	}

	// return "Question" struct.
	return &question
}
