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

var GetFactors = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "factors" and assign an array to the variable.
	var factors []models.Factor

	// Execute the SQL query to take all surveys and set it's result to the variable called "firstQuery".
	firstQuery, err := database.DBSQL.Query("SELECT * FROM factors;")
	if err != nil {
		log.Println(err)
		return
	}

	// Call "Close" function.
	defer firstQuery.Close()

	// Parse the result set of the first SQL query.
	for firstQuery.Next() {
		// Initialize the variable called "factor" and assign an "Factor" struct to the variable.
		var factor models.Factor

		// Call "Scan()" function on the result set of the first SQL query.
		if err := firstQuery.Scan(&factor.ID, &factor.Name); err != nil {
			log.Println(err)
			return
		}

		// Make SQL query to take information about all questions of the specific survey and set it's result to the variable called "secondQuery".
		secondQuery, err := database.DBSQL.Query(`SELECT
       		questions.id AS question_id,
       		questions.text AS question_text
		FROM factors_questions
		INNER JOIN questions
		ON factors_questions.question_id = questions.id
		WHERE factors_questions.factor_id = $1;`, factor.ID)
		if err != nil {
			log.Println(err)
			return
		}

		// Parse the result set of the second SQL query.
		for secondQuery.Next() {
			// Initialize the variable called "question" and assign an "Question" struct to the variable.
			var question models.Question

			// Call "Scan()" function on the result set of the second SQL query.
			if err := secondQuery.Scan(&question.ID, &question.Text); err != nil {
				log.Println(err)
				return
			}

			// Initialize the variable called "widget" and assign an "Widget" struct to the variable.
			var widget models.Widget

			// Execute the SQL query to take information about the widget of the specific question and set it's result to the variable called "thirdQuery".
			thirdQuery := database.DBSQL.QueryRow(`SELECT 
				widgets.id AS widget_id,
				widgets.name AS widget_name
			FROM questions_widgets
			INNER JOIN widgets
			ON questions_widgets.widget_id = widgets.id
			WHERE questions_widgets.question_id = $1;`, question.ID)

			// Call "Scan()" function on the result of the third SQL query.
			err = thirdQuery.Scan(&widget.ID, &widget.Name)
			if err != nil {
				log.Println(err)
				return
			}

			// Set the information about the widget of specific question to the "Widget" field of the struct "Question".
			question.Widget = widget

			// Make SQL query to take information about all options of the specific question and set it's result to the variable called "fourthSQL".
			fourthSQL, err := database.DBSQL.Query(`SELECT
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

			// Parse the result set of the fourth SQL query.
			for fourthSQL.Next() {
				// Initialize the variable called "option" and assign an "Option" struct to the variable.
				var option models.Option

				// Call "Scan()" function on the result set of the fourth SQL query.
				if err := fourthSQL.Scan(&option.ID, &option.Text); err != nil {
					log.Println(err)
					return
				}

				// Set the information about all options of specific question to the "Options" field of the struct "Question".
				question.Options = append(question.Options, option)
			}

			// Call "Close" function.
			fourthSQL.Close()

			// Set all questions to the field "Questions" of the struct "Survey".
			factor.Questions = append(factor.Questions, question)
		}

		// Call "Close" function.
		secondQuery.Close()

		// Set all surveys to the final array.
		factors = append(factors, factor)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, factors)
}

var CreateFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "factor" and assign struct to the variable.
	factor := models.Factor{}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&factor".
	if err := decoder.Decode(&factor); err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save new record to the "factors" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&factor).Error; err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "201" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "The new entry successfully created.")
}

var GetFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	factorID, err := strconv.Atoi(vars["factor_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "factor" and assign object with information about the specific factor to the variable.
	factor := GetFactorOr404(database.DBGORM, factorID, responseWriter, request)
	if factor == nil {
		return
	}

	// Make SQL query to take information about all questions of the specific survey and set it's result to the variable called "firstQuery".
	firstQuery, err := database.DBSQL.Query(`SELECT
       		questions.id AS question_id,
       		questions.text AS question_text
		FROM factors_questions
		INNER JOIN questions
		ON factors_questions.question_id = questions.id
		WHERE factors_questions.factor_id = $1;`, factorID)
	if err != nil {
		log.Println(err)
		return
	}

	// Parse the result set of the first SQL query.
	for firstQuery.Next() {
		// Initialize the variable called "question" and assign an "Question" struct to the variable.
		var question models.Question

		// Call "Scan()" function on the result set of the second SQL query.
		if err := firstQuery.Scan(&question.ID, &question.Text); err != nil {
			log.Println(err)
			return
		}

		// Initialize the variable called "widget" and assign an "Widget" struct to the variable.
		var widget models.Widget

		// Execute the SQL query to take information about the widget of the specific question and set it's result to the variable called "secondQuery".
		secondQuery := database.DBSQL.QueryRow(`SELECT 
			widgets.id AS widget_id,
			widgets.name AS widget_name
		FROM questions_widgets
		INNER JOIN widgets
		ON questions_widgets.widget_id = widgets.id
		WHERE questions_widgets.question_id = $1;`, question.ID)

		// Call "Scan()" function on the result of the second SQL query.
		err = secondQuery.Scan(&widget.ID, &widget.Name)
		if err != nil {
			log.Println(err)
			return
		}

		// Set the information about the widget of specific question to the "Widget" field of the struct "Question".
		question.Widget = widget

		// Make SQL query to take information about all options of the specific question and set it's result to the variable called "thirdSQL".
		thirdSQL, err := database.DBSQL.Query(`SELECT
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
		for thirdSQL.Next() {
			// Initialize the variable called "option" and assign an "Option" struct to the variable.
			var option models.Option

			// Call "Scan()" function on the result set of the third SQL query.
			if err := thirdSQL.Scan(&option.ID, &option.Text); err != nil {
				log.Println(err)
				return
			}

			// Set the information about all options of specific question to the "Options" field of the struct "Question".
			question.Options = append(question.Options, option)
		}

		// Call "Close" function.
		thirdSQL.Close()

		// Set all questions to the field "Questions" of the struct "Survey".
		factor.Questions = append(factor.Questions, question)
	}

	// Call "Close" function.
	firstQuery.Close()

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, factor)
}

var UpdateFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	factorID, err := strconv.Atoi(vars["factor_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "factor" and assign object with information about the specific factor to the variable if the record is exist in the database.
	factor := GetFactorOr404(database.DBGORM, factorID, responseWriter, request)
	if factor == nil {
		return
	}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&factor".
	if err := decoder.Decode(&factor); err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save record with new information to the "factors" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&factor).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully updated.")
}

var DeleteFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	factorID, err := strconv.Atoi(vars["factor_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "factor" and assign object with information about the specific factor to the variable if the record is exist in the database.
	factor := GetFactorOr404(database.DBGORM, factorID, responseWriter, request)
	if factor == nil {
		return
	}

	// Delete the record from the "factor" table with the help of "gorm" package.
	if err := database.DBGORM.Delete(&factor).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully deleted.")
}

func GetFactorOr404(db *gorm.DB, factorID int, responseWriter http.ResponseWriter, request *http.Request) *models.Factor {
	// Initialize the variable called "factor" and assign an struct to the variable.
	factor := models.Factor{}

	// Find the factor by id number.
	if err := db.First(&factor, models.Factor{ID: factorID}).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "Record not found.")
		return nil
	}

	// return "Factor" struct.
	return &factor
}
