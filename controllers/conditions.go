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

var GetConditions = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array.
	var conditions []models.Condition

	// CRUD interface of "GORM" ORM library to select all entries.
	if err := database.DBGORM.Find(&conditions).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Check the length of the array.
	if len(conditions) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, conditions)
}

var CreateCondition = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a "Condition" struct.
	condition := models.Condition{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&condition".
	if err := decoder.Decode(&condition); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&condition).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, strconv.Itoa(condition.ID))
}

var GetCondition = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	conditionIdentifier := mux.Vars(request)["condition_id"]

	// Variable has been initialized by assigning it to a "Condition" struct.
	condition := models.Condition{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", conditionIdentifier).Find(&condition).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, condition)
}

var UpdateCondition = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	conditionIdentifier := mux.Vars(request)["condition_id"]

	// Variable has been initialized by assigning it to a "Condition" struct.
	condition:= models.Condition{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", conditionIdentifier).Find(&condition).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&condition".
	if err := decoder.Decode(&condition); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to update information of the entry.
	if err := database.DBGORM.Save(&condition).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var DeleteCondition = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	conditionIdentifier := mux.Vars(request)["condition_id"]

	// Variable has been initialized by assigning it to a "Condition" struct.
	condition := models.Condition{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", conditionIdentifier).Find(&condition).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// CRUD interface of "GORM" ORM library to delete the entry.
	if err := database.DBGORM.Delete(&condition).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}
