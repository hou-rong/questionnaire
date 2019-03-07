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

var GetWidgets = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "widgets" and assign an array to the variable.
	var widgets []models.Widget

	// Take all records from "widgets" table with the help of "gorm" package.
	database.DBGORM.Find(&widgets)

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, widgets)
}

var CreateWidget = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Initialize the variable called "widget" and assign struct to the variable.
	widget := models.Widget{}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&widget".
	if err := decoder.Decode(&widget); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save new record to the "widgets" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&widget).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "201" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "The new entry successfully created.")
}

var GetWidget = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	widgetID, err := strconv.Atoi(vars["widget_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "widget" and assign object with information about the specific widget to the variable.
	widget := GetWidgetOr404(database.DBGORM, widgetID, responseWriter, request)
	if widget == nil {
		return
	}

	// Send successful response with status code "200" and JSON.
	utils.Response(responseWriter, http.StatusOK, widget)
}

var UpdateWidget = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	widgetID, err := strconv.Atoi(vars["widget_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "widget" and assign object with information about the specific widget to the variable if the record is exist in the database.
	widget := GetWidgetOr404(database.DBGORM, widgetID, responseWriter, request)
	if widget == nil {
		return
	}

	// NewDecoder returns a new decoder that reads from request body.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by "&widget".
	if err := decoder.Decode(&widget); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// Save record with new information to the "widgets" table with the help of "gorm" package.
	if err := database.DBGORM.Save(&widget).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully updated.")
}

var DeleteWidget = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take query parameters with the help of "mux.Vars" function.
	vars := mux.Vars(request)

	// Convert "string" to "int".
	widgetID, err := strconv.Atoi(vars["widget_id"])
	// Send response with detailed information about the error if the above process was unsuccessful.
	if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}

	// Initialize the variable called "widget" and assign object with information about the specific widget to the variable if the record is exist in the database.
	widget := GetWidgetOr404(database.DBGORM, widgetID, responseWriter, request)
	if widget == nil {
		return
	}

	// Delete the record from the "widget" table with the help of "gorm" package.
	if err := database.DBGORM.Delete(&widget).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully deleted.")
}

func GetWidgetOr404(db *gorm.DB, widgetID int, responseWriter http.ResponseWriter, request *http.Request) *models.Widget {
	// Initialize the variable called "widget" and assign an struct to the variable.
	widget := models.Widget{}

	// Find the widget by id number.
	if err := db.First(&widget, models.Widget{ID: widgetID}).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "Record not found.")
		return nil
	}

	// return "Widget" struct.
	return &widget
}
