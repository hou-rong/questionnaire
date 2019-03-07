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

var GetOptions = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "options" and assign an array to the variable.
	var options []models.Option

	// Take all records from "options" table with the help of "gorm" package.
	database.DBGORM.Find(&options)

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, options)
}

var CreateOption = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "option" and assign struct to the variable.
	option := models.Option{}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&option".
	if err := decoder.Decode(&option); err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save new record to the "options" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&option).Error; err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "201" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "The new entry successfully created.")
}

var GetOption = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	optionID, err := strconv.Atoi(vars["option_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "option" and assign object with information about the specific option to the variable.
	option := GetOptionOr404(database.DBGORM, optionID, responseWriter, request)
	if option == nil {
		return
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, option)
}

var UpdateOption = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	optionID, err := strconv.Atoi(vars["option_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "option" and assign object with information about the specific option to the variable if the record is exist in the database.
	option := GetOptionOr404(database.DBGORM, optionID, responseWriter, request)
	if option == nil {
		log.Println(err)
		return
	}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&option".
	if err := decoder.Decode(&option); err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save record with new information to the "options" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&option).Error; err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully updated.")
}

var DeleteOption = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	optionID, err := strconv.Atoi(vars["option_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "option" and assign object with information about the specific option to the variable if the record is exist in the database.
	option := GetOptionOr404(database.DBGORM, optionID, responseWriter, request)
	if option == nil {
		return
	}

	// Delete the record from the "options" table with the help of "gorm" package.
	if err := database.DBGORM.Delete(&option).Error; err != nil {
		// Send response with detailed information about the error if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully deleted.")
}

func GetOptionOr404(db *gorm.DB, optionID int, responseWriter http.ResponseWriter, request *http.Request) *models.Option {
	// Initialize the variable called "option" and assign an struct to the variable.
	option := models.Option{}

	// Find the option by id number.
	if err := db.First(&option, models.Option{ID: optionID}).Error; err != nil {
		// Send response with error message if the above process was unsuccessful.
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "Record not found.")
		return nil
	}

	// return "Option" struct.
	return &option
}
