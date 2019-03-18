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
)

var GetExtendedSurveys = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "surveys" and assign an array to the variable.
	var surveys []models.ExtendedSurvey

	// Execute the SQL query to take all surveys and set it's result to the variable called "firstQuery".
	firstQuery, err := database.DBSQL.Query("SELECT * FROM surveys;")
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function.
	defer firstQuery.Close()

	// Parse the result set of the first SQL query.
	for firstQuery.Next() {
		// Initialize the variable called "survey" and assign an "Survey" struct to the variable.
		var survey models.ExtendedSurvey

		// Call "Scan()" function on the result set of the first SQL query.
		if err := firstQuery.Scan(&survey.ID, &survey.Name, &survey.Description, &survey.Status, &survey.StartPeriod, &survey.EndPeriod); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Make SQL query to take information about all questions of the specific survey and set it's result to the variable called "secondQuery".
		secondQuery, err := database.DBSQL.Query(`SELECT
       		questions.id AS question_id,
       		questions.text AS question_text
		FROM surveys_questions
		INNER JOIN questions
		ON surveys_questions.question_id = questions.id
		WHERE surveys_questions.survey_id = $1;`, survey.ID)
		if err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Parse the result set of the second SQL query.
		for secondQuery.Next() {
			// Initialize the variable called "question" and assign an "Question" struct to the variable.
			var question models.Question

			// Call "Scan()" function on the result set of the second SQL query.
			if err := secondQuery.Scan(&question.ID, &question.Text); err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
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
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
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
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Parse the result set of the fourth SQL query.
			for fourthSQL.Next() {
				// Initialize the variable called "option" and assign an "Option" struct to the variable.
				var option models.Option

				// Call "Scan()" function on the result set of the fourth SQL query.
				if err := fourthSQL.Scan(&option.ID, &option.Text); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Set the information about all options of specific question to the "Options" field of the struct "Question".
				question.Options = append(question.Options, option)
			}

			// Call "Close" function.
			fourthSQL.Close()

			// Set all questions to the field "Questions" of the struct "Survey".
			survey.Questions = append(survey.Questions, question)
		}

		// Call "Close" function.
		secondQuery.Close()

		// Set all surveys to the final array.
		surveys = append(surveys, survey)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, surveys)
}

var GetAbridgedSurveys = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "surveys" and assign an array to the variable.
	var surveys []models.AbridgedSurvey

	// Execute the SQL query to take all surveys and set it's result to the variable called "rows".
	rows, err := database.DBSQL.Query("SELECT ID, NAME FROM surveys;")
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function.
	defer rows.Close()

	// Parse the result set of the SQL query.
	for rows.Next() {
		// Initialize the variable called "survey" and assign an "Survey" struct to the variable.
		var survey models.AbridgedSurvey

		// Call "Scan()" function on the result set of the SQL query.
		if err := rows.Scan(&survey.ID, &survey.Name); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Set all surveys to the final array.
		surveys = append(surveys, survey)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, surveys)
}

var GetAbridgedActiveSurveys = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "surveys" and assign an array to the variable.
	var surveys []models.AbridgedSurvey

	// Execute the SQL query to take all surveys and set it's result to the variable called "rows".
	rows, err := database.DBSQL.Query("SELECT ID, NAME FROM surveys WHERE STATUS = TRUE;")
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function.
	defer rows.Close()

	// Parse the result set of the SQL query.
	for rows.Next() {
		// Initialize the variable called "survey" and assign an "Survey" struct to the variable.
		var survey models.AbridgedSurvey

		// Call "Scan()" function on the result set of the SQL query.
		if err := rows.Scan(&survey.ID, &survey.Name); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Set all surveys to the final array.
		surveys = append(surveys, survey)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, surveys)
}

var GetAbridgedInactiveSurveys = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "surveys" and assign an array to the variable.
	var surveys []models.AbridgedSurvey

	// Execute the SQL query to take all surveys and set it's result to the variable called "rows".
	rows, err := database.DBSQL.Query("SELECT ID, NAME FROM surveys WHERE STATUS = FALSE;")
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function.
	defer rows.Close()

	// Parse the result set of the SQL query.
	for rows.Next() {
		// Initialize the variable called "survey" and assign an "Survey" struct to the variable.
		var survey models.AbridgedSurvey

		// Call "Scan()" function on the result set of the SQL query.
		if err := rows.Scan(&survey.ID, &survey.Name); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Set all surveys to the final array.
		surveys = append(surveys, survey)
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, surveys)
}

var CreateSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "survey" and assign struct to the variable.
	survey := models.ExtendedSurvey{}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&survey".
	if err := decoder.Decode(&survey); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save new record to the "surveys" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "201" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, survey.ID)
}

var GetSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	surveyID := vars["survey_id"]

	// Initialize the variable called "survey" and assign object with information about the specific survey to the variable.
	survey := GetSurveyOr404(database.DBGORM, surveyID, responseWriter, request)
	if survey == nil {
		return
	}

	// Make SQL query to take information about all questions of the specific survey and set it's result to the variable called "firstQuery".
	firstQuery, err := database.DBSQL.Query(`SELECT
       		questions.id AS question_id,
       		questions.text AS question_text
		FROM surveys_questions
		INNER JOIN questions
		ON surveys_questions.question_id = questions.id
		WHERE surveys_questions.survey_id = $1;`, surveyID)
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Parse the result set of the first SQL query.
	for firstQuery.Next() {
		// Initialize the variable called "question" and assign an "Question" struct to the variable.
		var question models.Question

		// Call "Scan()" function on the result set of the second SQL query.
		if err := firstQuery.Scan(&question.ID, &question.Text); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
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
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
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
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Parse the result set of the third SQL query.
		for thirdSQL.Next() {
			// Initialize the variable called "option" and assign an "Option" struct to the variable.
			var option models.Option

			// Call "Scan()" function on the result set of the third SQL query.
			if err := thirdSQL.Scan(&option.ID, &option.Text); err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Set the information about all options of specific question to the "Options" field of the struct "Question".
			question.Options = append(question.Options, option)
		}

		// Call "Close" function.
		thirdSQL.Close()

		// Set all questions to the field "Questions" of the struct "Survey".
		survey.Questions = append(survey.Questions, question)
	}

	// Call "Close" function.
	firstQuery.Close()

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, survey)
}

var UpdateSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)
	surveyID := vars["survey_id"]

	// Initialize the variable called "survey" and assign object with information about the specific widget to the variable if the record is exist in the database.
	survey := GetSurveyOr404(database.DBGORM, surveyID, responseWriter, request)
	if survey == nil {
		return
	}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&survey".
	if err := decoder.Decode(&survey); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save record with new information to the "widgets" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully updated.")
}

var DeleteSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)
	surveyID := vars["survey_id"]

	// Initialize the variable called "survey" and assign object with information about the specific widget to the variable if the record is exist in the database.
	survey := GetSurveyOr404(database.DBGORM, surveyID, responseWriter, request)
	if survey == nil {
		return
	}

	// Delete the record from the "widget" table with the help of "gorm" package.
	if err := database.DBGORM.Delete(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully deleted.")
}

func GetSurveyOr404(db *gorm.DB, surveyID string, responseWriter http.ResponseWriter, request *http.Request) *models.ExtendedSurvey {
	// Initialize the variable called "survey" and assign an struct to the variable.
	survey := models.ExtendedSurvey{}

	// Find the survey by id number which is uuid4.
	if err := db.First(&survey, models.ExtendedSurvey{ID: surveyID}).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "Record not found.")
		return nil
	}

	// return "Survey" struct.
	return &survey
}
