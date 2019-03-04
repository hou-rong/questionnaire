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
	var factors []models.Factor
	database.DBGORM.Find(&factors)
	utils.Response(responseWriter, http.StatusOK, factors)
}

var CreateFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	factor := models.Factor{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&factor); err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()
	if err := database.DBGORM.Save(&factor).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "The new entry successfully created.")
}

var GetFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	factorID, err := strconv.Atoi(vars["factor_id"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	factor := GetFactorOr404(database.DBGORM, factorID, responseWriter, request)
	if factor == nil {
		return
	}
	utils.Response(responseWriter, http.StatusOK, factor)
}

var UpdateFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	factorID, err := strconv.Atoi(vars["factor_id"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	factor := GetFactorOr404(database.DBGORM, factorID, responseWriter, request)
	if factor == nil {
		return
	}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&factor); err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()
	if err := database.DBGORM.Save(&factor).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully updated.")
}

var DeleteFactor = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	factorID, err := strconv.Atoi(vars["factor_id"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	factor := GetFactorOr404(database.DBGORM, factorID, responseWriter, request)
	if factor == nil {
		return
	}
	if err := database.DBGORM.Delete(&factor).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully deleted.")
}

func GetFactorOr404(db *gorm.DB, factorID int, responseWriter http.ResponseWriter, request *http.Request) *models.Factor {
	factor := models.Factor{}
	if err := db.First(&factor, models.Factor{FactorID: factorID}).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "Record not found.")
		return nil
	}
	return &factor
}
