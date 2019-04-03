package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
)

var GetBetaSurveys = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array.
	var surveys []models.BetaSurvey

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) == 0 {
		// CRUD interface of "GORM" ORM library to select all entries.
		if err := database.DBGORM.Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		// Variable has been initialized by assigning it a unique identifier of category.
		categoryIdentifier := keys.Get("category_id")

		// Variable has been initialized by assigning it a unique identifier of condition.
		conditionIdentifier := keys.Get("condition_id")

		// Variable has been initialized by assigning it a mark.
		mark := keys.Get("mark")

		// Variable has been initialized by assigning it a control.
		control := keys.Get("control")

		// Variable has been initialized by assigning it a email.
		email := keys.Get("email")

		if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(mark) != 0 && len(control) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, control, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, mark, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, control, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, mark, control, email).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, control, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, mark, control, email).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, control, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, mark, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(mark) != 0 && len(control) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, control.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND CONTROL = ?", categoryIdentifier, conditionIdentifier, mark, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, control.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND CONTROL = ?", categoryIdentifier, conditionIdentifier, mark, control).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, control.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND CONTROL = ?", categoryIdentifier, conditionIdentifier, mark, control).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, control.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND CONTROL = ?", categoryIdentifier, conditionIdentifier, mark, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(mark) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, mark, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, mark, email).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, mark, email).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, mark, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(control) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, control, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, control, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, control, email).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, control, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, control, email).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, control, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(conditionIdentifier) != 0 && len(mark) != 0 && len(control) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, control, email.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? CONTROL = ? AND EMAIL = ?", conditionIdentifier, mark, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, control, email.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? CONTROL = ? AND EMAIL = ?", conditionIdentifier, mark, control, email).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, control, email.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? CONTROL = ? AND EMAIL = ?", conditionIdentifier, mark, control, email).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, control, email.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? CONTROL = ? AND EMAIL = ?", conditionIdentifier, mark, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(categoryIdentifier) != 0 && len(mark) != 0 && len(control) != 0 && len(email) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, mark, control, email.
			if err := database.DBGORM.Where("CATEGORY = ? AND MARK = ? CONTROL = ? AND EMAIL = ?", categoryIdentifier, mark, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(mark) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ?", categoryIdentifier, conditionIdentifier, mark).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ?", categoryIdentifier, conditionIdentifier, mark).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ?", categoryIdentifier, conditionIdentifier, mark).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, mark.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND MARK = ?", categoryIdentifier, conditionIdentifier, mark).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(control) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, control.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND CONTROL = ?", categoryIdentifier, conditionIdentifier, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, control.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND CONTROL = ?", categoryIdentifier, conditionIdentifier, control).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, control.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND CONTROL = ?", categoryIdentifier, conditionIdentifier, control).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, control.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND CONTROL = ?", categoryIdentifier, conditionIdentifier, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, email).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, email).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition, email.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ? AND EMAIL = ?", categoryIdentifier, conditionIdentifier, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(categoryIdentifier) != 0 && len(mark) != 0 && len(control) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, mark, control.
			if err := database.DBGORM.Where("CATEGORY = ? AND MARK = ? AND CONTROL = ?", categoryIdentifier, mark, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(categoryIdentifier) != 0 && len(mark) != 0 && len(email) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, mark, email.
			if err := database.DBGORM.Where("CATEGORY = ? AND MARK = ? AND EMAIL = ?", categoryIdentifier, mark, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(categoryIdentifier) != 0 && len(control) != 0 && len(email) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, control, email.
			if err := database.DBGORM.Where("CATEGORY = ? AND CONTROL = ? AND EMAIL = ?", categoryIdentifier, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(conditionIdentifier) != 0 && len(mark) != 0 && len(control) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, control.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? AND CONTROL = ?", conditionIdentifier, mark, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, control.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? AND CONTROL = ?", conditionIdentifier, mark, control).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, control.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? AND CONTROL = ?", conditionIdentifier, mark, control).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, control.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? AND CONTROL = ?", conditionIdentifier, mark, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(conditionIdentifier) != 0 && len(mark) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, email.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? AND EMAIL = ?", conditionIdentifier, mark, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, email.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? AND EMAIL = ?", conditionIdentifier, mark, email).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, email.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? AND EMAIL = ?", conditionIdentifier, mark, email).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark, email.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ? AND EMAIL = ?", conditionIdentifier, mark, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(conditionIdentifier) != 0 && len(control) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, control, email.
				if err := database.DBGORM.Where("CONDITION = ? AND CONTROL = ? AND EMAIL = ?", conditionIdentifier, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, control, email.
				if err := database.DBGORM.Where("CONDITION = ? AND CONTROL = ? AND EMAIL = ?", conditionIdentifier, control, email).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, control, email.
				if err := database.DBGORM.Where("CONDITION = ? AND CONTROL = ? AND EMAIL = ?", conditionIdentifier, control, email).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, control, email.
				if err := database.DBGORM.Where("CONDITION = ? AND CONTROL = ? AND EMAIL = ?", conditionIdentifier, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(mark) != 0 && len(control) != 0 && len(email)!= 0 {
			// CRUD interface of "GORM" ORM library to find entries by mark, control, email.
			if err := database.DBGORM.Where("MARK = ? AND CONTROL = ? AND EMAIL = ?", mark, control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ?", categoryIdentifier, conditionIdentifier).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ?", categoryIdentifier, conditionIdentifier).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ?", categoryIdentifier, conditionIdentifier).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, unique identifier of condition.
				if err := database.DBGORM.Where("CATEGORY = ? AND CONDITION = ?", categoryIdentifier, conditionIdentifier).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(categoryIdentifier) != 0 && len(mark) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, mark.
			if err := database.DBGORM.Where("CATEGORY = ? AND MARK = ?", categoryIdentifier, mark).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(categoryIdentifier) != 0 && len(control) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, control.
			if err := database.DBGORM.Where("CATEGORY = ? AND CONTROL = ?", categoryIdentifier, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(categoryIdentifier) != 0 && len(email) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by unique identifier of category, email.
			if err := database.DBGORM.Where("CATEGORY = ? AND EMAIL = ?", categoryIdentifier, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(conditionIdentifier) != 0 && len(mark) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ?", conditionIdentifier, mark).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ?", conditionIdentifier, mark).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ?", conditionIdentifier, mark).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, mark.
				if err := database.DBGORM.Where("CONDITION = ? AND MARK = ?", conditionIdentifier, mark).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(conditionIdentifier) != 0 && len(control) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, control.
				if err := database.DBGORM.Where("CONDITION = ? AND CONTROL = ?", conditionIdentifier, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, control.
				if err := database.DBGORM.Where("CONDITION = ? AND CONTROL = ?", conditionIdentifier, control).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, control.
				if err := database.DBGORM.Where("CONDITION = ? AND CONTROL = ?", conditionIdentifier, control).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, control.
				if err := database.DBGORM.Where("CONDITION = ? AND CONTROL = ?", conditionIdentifier, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(conditionIdentifier) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, email.
				if err := database.DBGORM.Where("CONDITION = ? AND EMAIL = ?", conditionIdentifier, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, email.
				if err := database.DBGORM.Where("CONDITION = ? AND EMAIL = ?", conditionIdentifier, email).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, email.
				if err := database.DBGORM.Where("CONDITION = ? AND EMAIL = ?", conditionIdentifier, email).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entries by unique identifier of condition, email.
				if err := database.DBGORM.Where("CONDITION = ? AND EMAIL = ?", conditionIdentifier, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(mark) != 0 && len(control) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by mark, control.
			if err := database.DBGORM.Where("MARK = ? AND CONTROL = ?", mark, control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(mark) != 0 && len(email) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by mark, email.
			if err := database.DBGORM.Where("MARK = ? AND EMAIL = ?", mark, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(control) != 0 && len(email) != 0 {
			// CRUD interface of "GORM" ORM library to find entries by control, email.
			if err := database.DBGORM.Where("CONTROL = ? AND EMAIL = ?", control, email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(categoryIdentifier) != 0 {
			// CRUD interface of "GORM" ORM library to find entry by unique identifier of category.
			if err := database.DBGORM.Where("CATEGORY = ?", categoryIdentifier).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(conditionIdentifier) != 0 {
			if conditionIdentifier == "1" {
				// CRUD interface of "GORM" ORM library to find entry by unique identifier of condition.
				if err := database.DBGORM.Where("CONDITION = ?", conditionIdentifier).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "2" {
				// CRUD interface of "GORM" ORM library to find entry by unique identifier of condition.
				if err := database.DBGORM.Where("CONDITION = ?", conditionIdentifier).Order("START_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else if conditionIdentifier == "3" {
				// CRUD interface of "GORM" ORM library to find entry by unique identifier of condition.
				if err := database.DBGORM.Where("CONDITION = ?", conditionIdentifier).Order("END_PERIOD DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				// CRUD interface of "GORM" ORM library to find entry by unique identifier of condition.
				if err := database.DBGORM.Where("CONDITION = ?", conditionIdentifier).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}
			}
		} else if len(mark) != 0 {
			// CRUD interface of "GORM" ORM library to find entry by mark.
			if err := database.DBGORM.Where("MARK = ?", mark).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(control) != 0 {
			// CRUD interface of "GORM" ORM library to find entry by control.
			if err := database.DBGORM.Where("CONTROL = ?", control).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else if len(email) != 0 {
			// CRUD interface of "GORM" ORM library to find entry by EMAIL.
			if err := database.DBGORM.Where("EMAIL = ?", email).Order("CREATED_AT DESC").Find(&surveys).Error; err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
			return
		}
	}

	// Check the length of the array.
	if len(surveys) == 0 {
		utils.Response(responseWriter, http.StatusOK, nil)
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, surveys)
}

var GetBetaSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	surveyIdentifier := mux.Vars(request)["survey_id"]

	// Variable has been initialized by assigning it to a "BetaSurvey" struct.
	survey := models.BetaSurvey{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", surveyIdentifier).Find(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, survey)
}

var CreateSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it to a "BetaSurvey" struct.
	survey := models.BetaSurvey{}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from request body beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&survey".
	if err := decoder.Decode(&survey); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to create new entry.
	if err := database.DBGORM.Save(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "201".
	utils.ResponseWithSuccess(responseWriter, http.StatusCreated, survey.ID)
}

var UpdateSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	surveyIdentifier := mux.Vars(request)["survey_id"]

	// Variable has been initialized by assigning it to a "BetaSurvey" struct.
	survey := models.BetaSurvey{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", surveyIdentifier).Find(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// "NewDecoder" returns a new decoder that reads from request body.
	// The decoder introduces its own buffering and may read data from r beyond the JSON values requested.
	decoder := json.NewDecoder(request.Body)

	// Decode reads the JSON value from its input and stores it in the value pointed to by "&survey".
	if err := decoder.Decode(&survey); err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	// Close the HTTP request body.
	defer request.Body.Close()

	// CRUD interface of "GORM" ORM library to update information of the entry.
	if err := database.DBGORM.Save(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send JSON response with status code "200".
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}

var DeleteSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	surveyIdentifier := mux.Vars(request)["survey_id"]

	// Variable has been initialized by assigning it to a "BetaSurvey" struct.
	survey := models.BetaSurvey{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", surveyIdentifier).Find(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// CRUD interface of "GORM" ORM library to delete information of the entry.
	if err := database.DBGORM.Delete(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Send successful response with status code "200" and message.
	utils.ResponseWithSuccess(responseWriter, http.StatusOK, "http.StatusOK")
}
