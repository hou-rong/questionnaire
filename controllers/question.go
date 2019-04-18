package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strconv"
)

var GetAlphaQuestions = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) != 0 {
		// Variable has been initialized by assigning it a unique identifier of category.
		categoryIdentifier := keys.Get("category_id")

		// Variable "questions" has been initialized by assigning it to array of structures.
		var questions []models.AlphaQuestion

		// Execute the SQL query to get all questions by unique identifier of category.
		firstQuery, err := database.DBSQL.Query("SELECT ID, TEXT, WIDGET, REQUIRED, POSITION, CATEGORY FROM QUESTIONS WHERE CATEGORY = $1;", categoryIdentifier); if err != nil {
			logger.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Call "Close" function to the result set of the first SQL query.
		defer firstQuery.Close()

		// Parse the result set of the first SQL query.
		for firstQuery.Next() {
			// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
			var question models.AlphaQuestion

			// Call "Scan()" function to the result set of the first SQL query.
			if err := firstQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position, &question.Category); err != nil {
				logger.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Execute the SQL query to get information about all options of the specific question.
			secondQuery, err := database.DBSQL.Query(`SELECT
				OPTIONS.ID,
				OPTIONS.TEXT,
       			OPTIONS.POSITION
			FROM QUESTIONS_OPTIONS_RELATIONSHIP
			INNER JOIN OPTIONS
			ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
			WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
				logger.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Parse the result set of the second SQL query.
			for secondQuery.Next() {
				// Variable "option" has been initialized by assigning it to a "Option" struct.
				var option models.Option

				// Call "Scan()" function to the result set of the second SQL query.
				if err := secondQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
					logger.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Append information about option of the array.
				question.Options = append(question.Options, option)
			}

			// Call "Close" function to the result set of the second SQL query.
			secondQuery.Close()

			// Append information about question to the array.
			questions = append(questions, question)
		}

		// Send JSON response with status code "200".
		utils.Response(responseWriter, http.StatusOK, questions)
	} else {
		// Variable "questions" has been initialized by assigning it to array of structures.
		var questions []models.AlphaQuestion

		// Execute the SQL query to get all questions.
		firstQuery, err := database.DBSQL.Query("SELECT ID, TEXT, WIDGET, REQUIRED, POSITION, CATEGORY FROM QUESTIONS;"); if err != nil {
			logger.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Call "Close" function to the result set of the first SQL query.
		defer firstQuery.Close()

		// Parse the result set of the first SQL query.
		for firstQuery.Next() {
			// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
			var question models.AlphaQuestion

			// Call "Scan()" function to the result set of the first SQL query.
			if err := firstQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position, &question.Category); err != nil {
				logger.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Execute the SQL query to get information about all options of the specific question.
			secondQuery, err := database.DBSQL.Query(`SELECT
				OPTIONS.ID,
				OPTIONS.TEXT,
       			OPTIONS.POSITION
			FROM QUESTIONS_OPTIONS_RELATIONSHIP
			INNER JOIN OPTIONS
			ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
			WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
				logger.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Parse the result set of the second SQL query.
			for secondQuery.Next() {
				// Variable "option" has been initialized by assigning it to a "Option" struct.
				var option models.Option

				// Call "Scan()" function to the result set of the second SQL query.
				if err := secondQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
					logger.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Append information about option of the array.
				question.Options = append(question.Options, option)
			}

			// Call "Close" function to the result set of the second SQL query.
			secondQuery.Close()

			// Append information about question to the array.
			questions = append(questions, question)
		}

		// Send JSON response with status code "200".
		utils.Response(responseWriter, http.StatusOK, questions)
	}
}

var GetBetaQuestions = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) != 0 {
		// Variable has been initialized by assigning it a unique identifier of category.
		categoryIdentifier := keys.Get("category_id")

		// Variable has been initialized by assigning it a array.
		var questions []models.BetaQuestion

		// CRUD interface of "GORM" ORM library to select all entries by unique identifier of category.
		if err := database.DBGORM.Where("CATEGORY = ?", categoryIdentifier).Find(&questions).Error; err != nil {
			logger.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Check the length of the array.
		if len(questions) == 0 {
			utils.Response(responseWriter, http.StatusOK, nil)
			return
		}

		// Send JSON response with status code "200".
		utils.Response(responseWriter, http.StatusOK, questions)
	} else {
		// Variable has been initialized by assigning it a array.
		var questions []models.BetaQuestion

		// CRUD interface of "GORM" ORM library to select all entries.
		if err := database.DBGORM.Find(&questions).Error; err != nil {
			logger.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
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
}

var GetAlphaQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	questionIdentifier := mux.Vars(request)["question_id"]

	// Variable has been initialized by assigning it to a "AlphaQuestion" struct.
	question := models.AlphaQuestion{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", questionIdentifier).Find(&question).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// Execute the SQL query to get information about all options of the specific question.
	firstQuery, err := database.DBSQL.Query(`SELECT
			OPTIONS.ID,
			OPTIONS.TEXT,
       		OPTIONS.POSITION
		FROM QUESTIONS_OPTIONS_RELATIONSHIP
		INNER JOIN OPTIONS
		ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
		WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, questionIdentifier); if err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Parse the result set of the first SQL query.
	for firstQuery.Next() {
		// Variable "option" has been initialized by assigning it to a "Option" struct.
		var option models.Option

		// Call "Scan()" function on the result set of the first SQL query.
		if err := firstQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
			logger.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Append information about option of the array.
		question.Options = append(question.Options, option)
	}

	// Call "Close" function to the result set of the first SQL query.
	defer firstQuery.Close()

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, question)
}

var GetBetaQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	questionIdentifier := mux.Vars(request)["question_id"]

	// Variable has been initialized by assigning it to a "BetaQuestion" struct.
	question := models.BetaQuestion{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", questionIdentifier).Find(&question).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, question)
}

var CreateQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a "BetaQuestion" struct.
	question := models.BetaQuestion{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&question".
	if err := decoder.Decode(&question); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&question).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, strconv.Itoa(question.ID))
}

var UpdateQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	questionIdentifier := mux.Vars(request)["question_id"]

	// Variable has been initialized by assigning it a "BetaQuestion" struct.
	question := models.BetaQuestion{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", questionIdentifier).Find(&question).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&question".
	if err := decoder.Decode(&question); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to update information of the entry.
	if err := database.DBGORM.Save(&question).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var DeleteQuestion = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	questionIdentifier := mux.Vars(request)["question_id"]

	// Variable has been initialized by assigning it to a "BetaQuestion" struct.
	question := models.BetaQuestion{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", questionIdentifier).Find(&question).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// CRUD interface of "GORM" ORM library to delete the entry.
	if err := database.DBGORM.Delete(&question).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}
