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
)

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
