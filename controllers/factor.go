package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strconv"
	"strings"
)

var GetAlphaFactors = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array.
	var factors []models.AlphaFactor

	// Execute the SQL query to get all factors.
	firstQuery, err := database.DBSQL.Query("SELECT * FROM factors;"); if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function to the result set of the first SQL query.
	defer firstQuery.Close()

	// Parse the result set of the first SQL query.
	for firstQuery.Next() {
		// Variable "factor" has been initialized by assigning it to a "AlphaFactor" struct.
		var factor models.AlphaFactor

		// Call "Scan()" function to the result set of the first SQL query.
		if err := firstQuery.Scan(&factor.ID, &factor.Name); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Execute the SQL query to get information about all questions of the specific factor.
		secondQuery, err := database.DBSQL.Query(`SELECT
       		QUESTIONS.ID,
       		QUESTIONS.TEXT,
			QUESTIONS.WIDGET,
       		QUESTIONS.REQUIRED,
       		QUESTIONS.POSITION
		FROM FACTORS_QUESTIONS_RELATIONSHIP
		INNER JOIN QUESTIONS
		ON FACTORS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
		WHERE FACTORS_QUESTIONS_RELATIONSHIP.FACTOR_ID = $1;`, factor.ID); if err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Parse the result set of the second SQL query.
		for secondQuery.Next() {
			// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
			var question models.AlphaQuestion

			// Call "Scan()" function to the result set of the second SQL query.
			if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Execute SQL query to get information about all options of the specific question.
			thirdQuery, err := database.DBSQL.Query(`SELECT
				OPTIONS.ID,
				OPTIONS.TEXT,
       			OPTIONS.POSITION
			FROM QUESTIONS_OPTIONS_RELATIONSHIP
			INNER JOIN OPTIONS
			ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
			WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Parse the result set of the third SQL query.
			for thirdQuery.Next() {
				// Variable "option" has been initialized by assigning it to a "Option" struct.
				var option models.Option

				// Call "Scan()" function to the result set of the third SQL query.
				if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Append the information about option to the array.
				question.Options = append(question.Options, option)
			}

			// Call "Close" function to the result set of the third SQL query.
			thirdQuery.Close()

			// Append information about question to the array.
			factor.Questions = append(factor.Questions, question)
		}

		// Call "Close" function to the result set of the second SQL query.
		secondQuery.Close()

		// Append information about factor to the array.
		factors = append(factors, factor)
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, factors)
}

var GetBetaFactors = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array.
	var factors []models.BetaFactor

	// CRUD interface of "GORM" ORM library to select all entries.
	if err := database.DBGORM.Find(&factors).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
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

var GetAlphaFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	factorIdentifier := mux.Vars(request)["factor_id"]

	// Variable has been initialized by assigning it to a "AlphaFactor" struct.
	factor := models.AlphaFactor{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", factorIdentifier).Find(&factor).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// Execute SQL query to get information about all questions of the specific factor.
	firstQuery, err := database.DBSQL.Query(`SELECT
       		QUESTIONS.ID,
       		QUESTIONS.TEXT,
			QUESTIONS.WIDGET,
       		QUESTIONS.REQUIRED,
       		QUESTIONS.POSITION
		FROM FACTORS_QUESTIONS_RELATIONSHIP
		INNER JOIN QUESTIONS
		ON FACTORS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
		WHERE FACTORS_QUESTIONS_RELATIONSHIP.FACTOR_ID = $1;`, factorIdentifier); if err != nil {
		log.Println(err)
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
		if err := firstQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Execute SQL query to get information about all options of the specific question.
		secondQuery, err := database.DBSQL.Query(`SELECT
				OPTIONS.ID,
				OPTIONS.TEXT,
       			OPTIONS.POSITION
			FROM QUESTIONS_OPTIONS_RELATIONSHIP
			INNER JOIN OPTIONS
			ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
			WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Parse the result set of the third SQL query.
		for secondQuery.Next() {
			// Variable "option" has been initialized by assigning it to a "Option" struct.
			var option models.Option

			// Call "Scan()" function to the result set of the third SQL query.
			if err := secondQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Append the information about option to the array.
			question.Options = append(question.Options, option)
		}

		// Call "Close" function to the result set of the second SQL query.
		secondQuery.Close()

		// Append information about question to the array.
		factor.Questions = append(factor.Questions, question)
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, factor)
}

var GetBetaFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	factorIdentifier := mux.Vars(request)["factor_id"]

	// Variable has been initialized by assigning it to a "BetaFactor" struct.
	factor := models.BetaFactor{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", factorIdentifier).Find(&factor).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, factor)
}

var CreateFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a "BetaFactor" struct.
	factor := models.BetaFactor{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&factor".
	if err := decoder.Decode(&factor); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&factor).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, strconv.Itoa(factor.ID))
}

var UpdateFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	factorIdentifier := mux.Vars(request)["factor_id"]

	// Variable has been initialized by assigning it to a "BetaFactor" struct.
	factor := models.BetaFactor{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", factorIdentifier).Find(&factor).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&factor".
	if err := decoder.Decode(&factor); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to update information of the entry.
	if err := database.DBGORM.Save(&factor).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var DeleteFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	factorIdentifier := mux.Vars(request)["factor_id"]

	// Variable has been initialized by assigning it to a "BetaFactor" struct.
	factor := models.BetaFactor{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", factorIdentifier).Find(&factor).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// CRUD interface of "GORM" ORM library to delete the entry.
	if err := database.DBGORM.Delete(&factor).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var DeleteInFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	factorIdentifier := mux.Vars(request)["factor_id"]

	// Build second SQL statement.
	var secondStatement strings.Builder
	secondStatement.WriteString("CALL proper(")
	secondStatement.WriteString(factorIdentifier)
	secondStatement.WriteString(")")

	// Execute SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(secondStatement.String()); if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var DeleteOutFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	factorIdentifier := mux.Vars(request)["factor_id"]

	// Build second SQL statement.
	var secondStatement strings.Builder
	secondStatement.WriteString("CALL tide(")
	secondStatement.WriteString(factorIdentifier)
	secondStatement.WriteString(")")

	// Execute SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(secondStatement.String()); if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}