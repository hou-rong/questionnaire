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

var GetOptions = func(responseWriter http.ResponseWriter, request *http.Request) {
	var options []models.Option
	database.DBGORM.Find(&options)
	utils.Response(responseWriter, http.StatusOK, options)
}

var CreateOption = func(responseWriter http.ResponseWriter, request *http.Request) {
	option := models.Option{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&option); err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()
	if err := database.DBGORM.Save(&option).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "The new entry successfully created.")
}

var GetOption = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	optionID, err := strconv.Atoi(vars["option_id"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	option := GetOptionOr404(database.DBGORM, optionID, responseWriter, request)
	if option == nil {
		return
	}
	utils.Response(responseWriter, http.StatusOK, option)
}

var UpdateOption = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	optionID, err := strconv.Atoi(vars["option_id"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	option := GetOptionOr404(database.DBGORM, optionID, responseWriter, request)
	if option == nil {
		return
	}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&option); err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()
	if err := database.DBGORM.Save(&option).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully updated.")
}

var DeleteOption = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	optionID, err := strconv.Atoi(vars["option_id"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	option := GetOptionOr404(database.DBGORM, optionID, responseWriter, request)
	if option == nil {
		return
	}
	if err := database.DBGORM.Delete(&option).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully deleted.")
}

func GetOptionOr404(db *gorm.DB, optionID int, responseWriter http.ResponseWriter, request *http.Request) *models.Option {
	option := models.Option{}
	if err := db.First(&option, models.Option{OptionID: optionID}).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "Record not found.")
		return nil
	}
	return &option
}

