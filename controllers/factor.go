package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"strconv"
)

var GetFactors = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "factors" and assign an array to the variable.
	var factors []models.Factor

	// Take all records from "factors" table with the help of "gorm" package.
	database.DBGORM.Find(&factors)

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
