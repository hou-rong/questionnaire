package controllers

import (
	"encoding/json"
	"github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strconv"
	"strings"
)

var CreateSingleIntQuestionOptionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a "QuestionOptionRelationship" struct.
	questionOptionRelationship := models.QuestionOptionRelationship{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&questionOptionRelationship".
	if err := decoder.Decode(&questionOptionRelationship); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&questionOptionRelationship).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var CreateMultipleIntQuestionOptionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize "RequestBody" struct.
	type RequestBody struct {
		QuestionID int `json:"question_id"`
		Options []int `json:"options"`
	}

	// Variable has been initialized by assigning it a "RequestBody" struct.
	requestBody := RequestBody{}

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
	sqlStatement.WriteString("INSERT INTO QUESTIONS_OPTIONS_RELATIONSHIP (QUESTION_ID, OPTION_ID) SELECT ")
	sqlStatement.WriteString(strconv.Itoa(requestBody.QuestionID))
	sqlStatement.WriteString(" QUESTION_ID, OPTION_ID FROM UNNEST(ARRAY[")
	sqlStatement.WriteString(utils.ConvertIntArrayToString(requestBody.Options))
	sqlStatement.WriteString("]) OPTION_ID")
	sqlStatement.WriteString(" ON CONFLICT ON CONSTRAINT QUESTIONS_OPTIONS_RELATIONSHIP_UNIQUE_KEY DO NOTHING")

	// Make SQL query by "database/sql" package.
	_, err := database.DBSQL.Exec(sqlStatement.String()); if err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "http.StatusCreated")
}

var CreateMultipleTextQuestionOptionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize "RequestBody" struct.
	type RequestBody struct {
		QuestionID int `json:"question_id"`
		Options [] struct {
			Text string `json:"option_text"`
		} `json:"options"`
	}

	// Variable has been initialized by assigning it a "RequestBody" struct.
	requestBody := RequestBody{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&requestBody".
	if err := decoder.Decode(&requestBody); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Initialize several arrays which would be used in string builder.
	var textArray []string
	var positionArray []int

	// Parse content of the request body.
	for i := 0; i < len(requestBody.Options); i++ {
		textArray = append(textArray, requestBody.Options[i].Text)
		positionArray = append(positionArray, i)
	}

	// Build SQL statement.
	var sqlStatement strings.Builder
	sqlStatement.WriteString("SELECT ARRAY_AGG(RESULTS.ID) FROM alexa (")
	sqlStatement.WriteString(strconv.Itoa(requestBody.QuestionID))
	sqlStatement.WriteString(", ARRAY[")
	sqlStatement.WriteString(utils.ConvertStringArrayToString(textArray))
	sqlStatement.WriteString("], ARRAY[")
	sqlStatement.WriteString(utils.ConvertIntArrayToString(positionArray))
	sqlStatement.WriteString("]) RESULTS(ID);")

	// Variable has been initialized by assigning it a array.
	var optionsIdentifiers pq.Int64Array

	// Execute SQL query by "database/sql" package.
	if err := database.DBSQL.QueryRow(sqlStatement.String()).Scan(&optionsIdentifiers); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and array of ids.
	utils.Response(responseWriter, http.StatusOK, map[string]pq.Int64Array{"options": optionsIdentifiers})
}

var DeleteSingleIntQuestionOptionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of question.
		questionIdentifier := keys.Get("question_id")

		// Variable has been initialized by assigning it a unique identifier of option.
		optionIdentifier := keys.Get("option_id")

		// Check the value of the variables.
		if len(questionIdentifier) != 0 && len(optionIdentifier) != 0 {
			// Variable has been initialized by assigning it a "QuestionOptionRelationship" struct.
			questionOptionRelationship := models.QuestionOptionRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("QUESTION_ID = ? AND OPTION_ID = ?", questionIdentifier, optionIdentifier).Delete(&questionOptionRelationship).Error; err != nil {
				logger.Println(err)
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

var DeleteMultipleIntQuestionOptionRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of question.
		questionIdentifier := keys.Get("question_id")

		// Check the value of the variables.
		if len(questionIdentifier) != 0 {
			// Variable has been initialized by assigning it a "QuestionOptionRelationship" struct.
			questionOptionRelationship := models.QuestionOptionRelationship{}

			// CRUD interface of "GORM" ORM library to delete information of the entry.
			if err := database.DBGORM.Where("QUESTION_ID = ?", questionIdentifier).Delete(&questionOptionRelationship).Error; err != nil {
				logger.Println(err)
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

var GetQuestionsOptionsRelationship = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Variable has been initialized by assigning it a array.
	var options []models.Option

	// Check if an array contains any element.
	if len(keys) > 0 {
		// Variable has been initialized by assigning it a unique identifier of question.
		questionIdentifier := keys.Get("question_id")

		// Check the value of the variables.
		if len(questionIdentifier) != 0 {
			// Make SQL query by "database/sql" package.
			rows, err := database.DBSQL.Query("SELECT ID, TEXT FROM QUESTIONS_OPTIONS_RELATIONSHIP INNER JOIN OPTIONS ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID WHERE QUESTION_ID = $1", questionIdentifier); if err != nil {
				logger.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function.
			defer rows.Close()

			// Parse the result set of the first SQL query.
			for rows.Next() {
				// Variable has been initialized by assigning it to a "Option" struct.
				var option models.Option

				// Call "Scan()" function on the result set of the SQL query.
				if err := rows.Scan(&option.ID, &option.Text); err != nil {
					logger.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Set all factors to the final array.
				options = append(options, option)
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
	if len(options) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, options)
}
