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

var GetWidgets = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a array.
	var widgets []models.Widget

	// CRUD interface of "GORM" ORM library to select all entries.
	if err := database.DBGORM.Find(&widgets).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Check the length of the array.
	if len(widgets) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, widgets)
}

var CreateWidget = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Variable has been initialized by assigning it a "Widget" struct.
	widget := models.Widget{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&widget".
	if err := decoder.Decode(&widget); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&widget).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, strconv.Itoa(widget.ID))
}

var GetWidget = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	widgetIdentifier := mux.Vars(request)["widget_id"]

	// Variable has been initialized by assigning it to a "Widget" struct.
	widget := models.Widget{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", widgetIdentifier).Find(&widget).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, widget)
}

var UpdateWidget = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	widgetIdentifier := mux.Vars(request)["widget_id"]

	// Variable has been initialized by assigning it to a "Widget" struct.
	widget := models.Widget{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", widgetIdentifier).Find(&widget).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from r beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&widget".
	if err := decoder.Decode(&widget); err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to update information of the entry.
	if err := database.DBGORM.Save(&widget).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var DeleteWidget = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Create and customize logger.
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	widgetIdentifier := mux.Vars(request)["widget_id"]

	// Variable has been initialized by assigning it to a "Widget" struct.
	widget := models.Widget{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", widgetIdentifier).Find(&widget).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// CRUD interface of "GORM" ORM library to delete the entry.
	if err := database.DBGORM.Delete(&widget).Error; err != nil {
		logger.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}
