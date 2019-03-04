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

var GetWidgetsTypes = func(responseWriter http.ResponseWriter, request *http.Request) {
	var widgetsTypes []models.WidgetType
	database.DBGORM.Find(&widgetsTypes)
	utils.Response(responseWriter, http.StatusOK, widgetsTypes)
}

var CreateWidgetType = func(responseWriter http.ResponseWriter, request *http.Request) {
	widgetType := models.WidgetType{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&widgetType); err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()
	if err := database.DBGORM.Save(&widgetType).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, "The new entry successfully created.")
}

var GetWidgetType = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	widgetTypeID, err := strconv.Atoi(vars["widget_type"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	widgetType := GetWidgetTypeOr404(database.DBGORM, widgetTypeID, responseWriter, request)
	if widgetType == nil {
		return
	}
	utils.Response(responseWriter, http.StatusOK, widgetType)
}

var UpdateWidgetType = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	widgetTypeID, err := strconv.Atoi(vars["widget_type"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	widgetType := GetWidgetTypeOr404(database.DBGORM, widgetTypeID, responseWriter, request)
	if widgetType == nil {
		return
	}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&widgetType); err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()
	if err := database.DBGORM.Save(&widgetType).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully updated.")
}

var DeleteWidgetType = func(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	widgetTypeID, err := strconv.Atoi(vars["widget_type"])
	if err != nil {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "The request parameter is invalid.")
		return
	}
	widgetType := GetWidgetTypeOr404(database.DBGORM, widgetTypeID, responseWriter, request)
	if widgetType == nil {
		return
	}
	if err := database.DBGORM.Delete(&widgetType).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The entry successfully deleted.")
}

func GetWidgetTypeOr404(db *gorm.DB, widgetTypeID int, responseWriter http.ResponseWriter, request *http.Request) *models.WidgetType {
	widgetType := models.WidgetType{}
	if err := db.First(&widgetType, models.WidgetType{WidgetTypeID: widgetTypeID}).Error; err != nil {
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "Record not found.")
		return nil
	}
	return &widgetType
}
