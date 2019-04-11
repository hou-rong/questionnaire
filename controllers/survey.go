package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"questionnaire/database"
	"questionnaire/models"
	"questionnaire/utils"
	"time"
)

var GetAlphaSurveys = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it a array.
	var surveys []models.AlphaSurvey

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) == 0 {
		// Execute the SQL query to get all surveys.
		firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS ORDER BY CREATED_AT DESC;"); if err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Call "Close" function to the result set of the first SQL query.
		defer firstQuery.Close()

		// Parse the result set of the first SQL query.
		for firstQuery.Next() {
			// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
			var survey models.AlphaSurvey

			// Call "Scan()" function to the result set of the first SQL query.
			if err := firstQuery.Scan(&survey.ID,
				&survey.Name,
				&survey.Description,
				&survey.Category,
				&survey.Condition,
				&survey.Mark,
				&survey.Control,
				&survey.StartPeriod,
				&survey.EndPeriod,
				&survey.CreatedAt,
				&survey.UpdatedAt,
				&survey.Email,
				&survey.Blocked,
				&survey.TotalRespondents,
				&survey.PastRespondents); err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Execute the SQL query to get all question.
			secondQuery, err := database.DBSQL.Query(`SELECT
				QUESTIONS.ID,
				QUESTIONS.TEXT,
				QUESTIONS.WIDGET,
				QUESTIONS.REQUIRED,
				QUESTIONS.POSITION
			FROM SURVEYS_QUESTIONS_RELATIONSHIP
			INNER JOIN QUESTIONS
			ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
			WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Parse the result set of the second SQL query.
			for secondQuery.Next() {
				// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
				var question models.AlphaQuestion

				// Call "Scan()" function to the result set of the second SQL query.
				if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute SQL query to get information about all options of the specific question.
				thirdQuery, err := database.DBSQL.Query(`SELECT
					OPTIONS.ID,
					OPTIONS.TEXT,
       				OPTIONS.POSITION
				FROM QUESTIONS_OPTIONS_RELATIONSHIP
				INNER JOIN OPTIONS
				ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
				WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the third SQL query.
				for thirdQuery.Next() {
					// Variable "option" has been initialized by assigning it to a "Option" struct.
					var option models.Option

					// Call "Scan()" function to the result set of the third SQL query.
					if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Append the information about option to the array.
					question.Options = append(question.Options, option)
				}

				// Call "Close" function to the result set of the second SQL query.
				thirdQuery.Close()

				// Append information about question to the array.
				survey.Questions = append(survey.Questions, question)
			}

			// Call "Close" function to the result set of the second SQL query.
			secondQuery.Close()

			// Append information about survey to the array.
			surveys = append(surveys, survey)
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
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND CONTROL = $4 AND EMAIL = $5 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, mark, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND CONTROL = $4 AND EMAIL = $5 ORDER BY START_PERIOD DESC;", categoryIdentifier, conditionIdentifier, mark, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND CONTROL = $4 AND EMAIL = $5 ORDER BY END_PERIOD DESC;", categoryIdentifier, conditionIdentifier, mark, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND CONTROL = $4 AND EMAIL = $5 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, mark, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(mark) != 0 && len(control) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND CONTROL = $4 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, mark, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND CONTROL = $4 ORDER BY START_PERIOD DESC;", categoryIdentifier, conditionIdentifier, mark, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND CONTROL = $4 ORDER BY END_PERIOD DESC;", categoryIdentifier, conditionIdentifier, mark, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND CONTROL = $4 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, mark, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(mark) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND EMAIL = $4 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, mark, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND EMAIL = $4 ORDER BY START_PERIOD DESC;", categoryIdentifier, conditionIdentifier, mark, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND EMAIL = $4 ORDER BY END_PERIOD DESC;", categoryIdentifier, conditionIdentifier, mark, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 AND EMAIL = $4 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, mark, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(control) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY START_PERIOD DESC;", categoryIdentifier, conditionIdentifier, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY END_PERIOD DESC;", categoryIdentifier, conditionIdentifier, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
				QUESTIONS.ID,
				QUESTIONS.TEXT,
				QUESTIONS.WIDGET,
				QUESTIONS.REQUIRED,
				QUESTIONS.POSITION
			FROM SURVEYS_QUESTIONS_RELATIONSHIP
			INNER JOIN QUESTIONS
			ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
			WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
					OPTIONS.ID,
					OPTIONS.TEXT,
       				OPTIONS.POSITION
				FROM QUESTIONS_OPTIONS_RELATIONSHIP
				INNER JOIN OPTIONS
				ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
				WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(conditionIdentifier) != 0 && len(mark) != 0 && len(control) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY CREATED_AT DESC;", conditionIdentifier, mark, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY START_PERIOD DESC;", conditionIdentifier, mark, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY END_PERIOD DESC;", conditionIdentifier, mark, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY CREATED_AT DESC;", conditionIdentifier, mark, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(categoryIdentifier) != 0 && len(mark) != 0 && len(control) != 0 && len(email) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND MARK = $2 AND CONTROL = $3 AND EMAIL = $4 ORDER BY CREATED_AT DESC;", categoryIdentifier, mark, control, email); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(mark) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, mark); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 ORDER BY START_PERIOD DESC;", categoryIdentifier, conditionIdentifier, mark); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 ORDER BY END_PERIOD DESC;", categoryIdentifier, conditionIdentifier, mark); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND MARK = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, mark); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(control) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND CONTROL = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND CONTROL = $3 ORDER BY START_PERIOD DESC;", categoryIdentifier, conditionIdentifier, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND CONTROL = $3 ORDER BY END_PERIOD DESC;", categoryIdentifier, conditionIdentifier, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND CONTROL = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND EMAIL = $3 ORDER BY START_PERIOD DESC;", categoryIdentifier, conditionIdentifier, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND EMAIL = $3 ORDER BY END_PERIOD DESC;", categoryIdentifier, conditionIdentifier, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(categoryIdentifier) != 0 && len(mark) != 0 && len(control) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND MARK = $2 AND CONTROL = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, mark, control); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(categoryIdentifier) != 0 && len(mark) != 0 && len(email) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND MARK = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, mark, email); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(categoryIdentifier) != 0 && len(control) != 0 && len(email) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONTROL = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", categoryIdentifier, control, email); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(conditionIdentifier) != 0 && len(mark) != 0 && len(control) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND CONTROL = $3 ORDER BY CREATED_AT DESC;", conditionIdentifier, mark, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND CONTROL = $3 ORDER BY START_PERIOD DESC;", conditionIdentifier, mark, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND CONTROL = $3 ORDER BY END_PERIOD DESC;", conditionIdentifier, mark, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND CONTROL = $3 ORDER BY CREATED_AT DESC;", conditionIdentifier, mark, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(conditionIdentifier) != 0 && len(mark) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", conditionIdentifier, mark, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND EMAIL = $3 ORDER BY START_PERIOD DESC;", conditionIdentifier, mark, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND EMAIL = $3 ORDER BY END_PERIOD DESC;", conditionIdentifier, mark, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", conditionIdentifier, mark, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(conditionIdentifier) != 0 && len(control) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND CONTROL = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", conditionIdentifier, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
			  		FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND CONTROL = $2 AND EMAIL = $3 ORDER BY START_PERIOD DESC;", conditionIdentifier, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND CONTROL = $2 AND EMAIL = $3 ORDER BY END_PERIOD DESC;", conditionIdentifier, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND CONTROL = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", conditionIdentifier, control, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(mark) != 0 && len(control) != 0 && len(email)!= 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE MARK = $1 AND CONTROL = $2 AND EMAIL = $3 ORDER BY CREATED_AT DESC;", mark, control, email); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(categoryIdentifier) != 0 && len(conditionIdentifier) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 ORDER BY START_PERIOD DESC;", categoryIdentifier, conditionIdentifier); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONDITION = $2 ORDER BY END_PERIOD DESC;", categoryIdentifier, conditionIdentifier); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = ? AND CONDITION = ? ORDER BY CREATED_AT DESC;", categoryIdentifier, conditionIdentifier); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(categoryIdentifier) != 0 && len(mark) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND MARK = $2 ORDER BY CREATED_AT DESC;", categoryIdentifier, mark); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(categoryIdentifier) != 0 && len(control) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND CONTROL = $2 ORDER BY CREATED_AT DESC;", categoryIdentifier, control); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(categoryIdentifier) != 0 && len(email) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 AND EMAIL = $2 ORDER BY CREATED_AT DESC;", categoryIdentifier, email); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(conditionIdentifier) != 0 && len(mark) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 ORDER BY CREATED_AT DESC;", conditionIdentifier, mark); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 ORDER BY START_PERIOD DESC;", conditionIdentifier, mark); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 ORDER BY END_PERIOD DESC;", conditionIdentifier, mark); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND MARK = $2 ORDER BY CREATED_AT DESC;", conditionIdentifier, mark); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(conditionIdentifier) != 0 && len(control) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND CONTROL = $2 ORDER BY CREATED_AT DESC;", conditionIdentifier, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND CONTROL = $2 ORDER BY START_PERIOD DESC;", conditionIdentifier, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND CONTROL = $2 ORDER BY END_PERIOD DESC;", conditionIdentifier, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND CONTROL = $2 ORDER BY CREATED_AT DESC;", conditionIdentifier, control); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(conditionIdentifier) != 0 && len(email) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND EMAIL = $2 ORDER BY CREATED_AT DESC;", conditionIdentifier, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND EMAIL = $2 ORDER BY START_PERIOD DESC;", conditionIdentifier, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND EMAIL = $2 ORDER BY END_PERIOD DESC;", conditionIdentifier, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 AND EMAIL = $2 ORDER BY CREATED_AT DESC;", conditionIdentifier, email); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(mark) != 0 && len(control) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE MARK = $1 AND CONTROL = $2 ORDER BY CREATED_AT DESC;", mark, control); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(mark) != 0 && len(email) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE MARK = $1 AND EMAIL = $2 ORDER BY CREATED_AT DESC;", mark, email); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(control) != 0 && len(email) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONTROL = $1 AND EMAIL = $2 ORDER BY CREATED_AT DESC;", control, email); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(categoryIdentifier) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CATEGORY = $1 ORDER BY CREATED_AT DESC;", categoryIdentifier); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(conditionIdentifier) != 0 {
			if conditionIdentifier == "1" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 ORDER BY CREATED_AT DESC;", conditionIdentifier); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "2" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 ORDER BY START_PERIOD DESC;", conditionIdentifier); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else if conditionIdentifier == "3" {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 ORDER BY END_PERIOD DESC;", conditionIdentifier); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			} else {
				// Execute the SQL query to get all surveys.
				firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONDITION = $1 ORDER BY CREATED_AT DESC;", conditionIdentifier); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Call "Close" function to the result set of the first SQL query.
				defer firstQuery.Close()

				// Parse the result set of the first SQL query.
				for firstQuery.Next() {
					// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
					var survey models.AlphaSurvey

					// Call "Scan()" function to the result set of the first SQL query.
					if err := firstQuery.Scan(&survey.ID,
						&survey.Name,
						&survey.Description,
						&survey.Category,
						&survey.Condition,
						&survey.Mark,
						&survey.Control,
						&survey.StartPeriod,
						&survey.EndPeriod,
						&survey.CreatedAt,
						&survey.UpdatedAt,
						&survey.Email,
						&survey.Blocked,
						&survey.TotalRespondents,
						&survey.PastRespondents); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute the SQL query to get all question.
					secondQuery, err := database.DBSQL.Query(`SELECT
						QUESTIONS.ID,
						QUESTIONS.TEXT,
						QUESTIONS.WIDGET,
						QUESTIONS.REQUIRED,
						QUESTIONS.POSITION
					FROM SURVEYS_QUESTIONS_RELATIONSHIP
					INNER JOIN QUESTIONS
					ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
					WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the second SQL query.
					for secondQuery.Next() {
						// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
						var question models.AlphaQuestion

						// Call "Scan()" function to the result set of the second SQL query.
						if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Execute SQL query to get information about all options of the specific question.
						thirdQuery, err := database.DBSQL.Query(`SELECT
							OPTIONS.ID,
							OPTIONS.TEXT,
       						OPTIONS.POSITION
						FROM QUESTIONS_OPTIONS_RELATIONSHIP
						INNER JOIN OPTIONS
						ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
						WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Parse the result set of the third SQL query.
						for thirdQuery.Next() {
							// Variable "option" has been initialized by assigning it to a "Option" struct.
							var option models.Option

							// Call "Scan()" function to the result set of the third SQL query.
							if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
								log.Println(err)
								utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
								return
							}

							// Append the information about option to the array.
							question.Options = append(question.Options, option)
						}

						// Call "Close" function to the result set of the second SQL query.
						thirdQuery.Close()

						// Append information about question to the array.
						survey.Questions = append(survey.Questions, question)
					}

					// Call "Close" function to the result set of the second SQL query.
					secondQuery.Close()

					// Append information about survey to the array.
					surveys = append(surveys, survey)
				}
			}
		} else if len(mark) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE MARK = $1 ORDER BY CREATED_AT DESC;", mark); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(control) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE CONTROL = $1 ORDER BY CREATED_AT DESC;", control); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
			}
		} else if len(email) != 0 {
			// Execute the SQL query to get all surveys.
			firstQuery, err := database.DBSQL.Query("SELECT * FROM SURVEYS WHERE EMAIL = $1 ORDER BY CREATED_AT DESC;", email); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "survey" has been initialized by assigning it to a "AlphaSurvey" struct.
				var survey models.AlphaSurvey

				// Call "Scan()" function to the result set of the first SQL query.
				if err := firstQuery.Scan(&survey.ID,
					&survey.Name,
					&survey.Description,
					&survey.Category,
					&survey.Condition,
					&survey.Mark,
					&survey.Control,
					&survey.StartPeriod,
					&survey.EndPeriod,
					&survey.CreatedAt,
					&survey.UpdatedAt,
					&survey.Email,
					&survey.Blocked,
					&survey.TotalRespondents,
					&survey.PastRespondents); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute the SQL query to get all question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					QUESTIONS.ID,
					QUESTIONS.TEXT,
					QUESTIONS.WIDGET,
					QUESTIONS.REQUIRED,
					QUESTIONS.POSITION
				FROM SURVEYS_QUESTIONS_RELATIONSHIP
				INNER JOIN QUESTIONS
				ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
				WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, survey.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
					var question models.AlphaQuestion

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Execute SQL query to get information about all options of the specific question.
					thirdQuery, err := database.DBSQL.Query(`SELECT
						OPTIONS.ID,
						OPTIONS.TEXT,
       					OPTIONS.POSITION
					FROM QUESTIONS_OPTIONS_RELATIONSHIP
					INNER JOIN OPTIONS
					ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
					WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Parse the result set of the third SQL query.
					for thirdQuery.Next() {
						// Variable "option" has been initialized by assigning it to a "Option" struct.
						var option models.Option

						// Call "Scan()" function to the result set of the third SQL query.
						if err := thirdQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
							log.Println(err)
							utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
							return
						}

						// Append the information about option to the array.
						question.Options = append(question.Options, option)
					}

					// Call "Close" function to the result set of the second SQL query.
					thirdQuery.Close()

					// Append information about question to the array.
					survey.Questions = append(survey.Questions, question)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about survey to the array.
				surveys = append(surveys, survey)
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

var GetAlphaSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Take variable from path with the help of "Gorilla Mux" library.
	// The most common numeric conversions are Atoi (string to int) and Itoa (int to string).
	surveyIdentifier := mux.Vars(request)["survey_id"]

	// Variable has been initialized by assigning it to a "AlphaSurvey" struct.
	survey := models.AlphaSurvey{}

	// CRUD interface of "GORM" ORM library to find entry by unique identifier.
	if err := database.DBGORM.Where("ID = ?", surveyIdentifier).Find(&survey).Error; err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusNotFound, "http.StatusNotFound")
		return
	}

	// Execute the SQL query to get all question.
	firstQuery, err := database.DBSQL.Query(`SELECT
		QUESTIONS.ID,
		QUESTIONS.TEXT,
		QUESTIONS.WIDGET,
		QUESTIONS.REQUIRED,
		QUESTIONS.POSITION
	FROM SURVEYS_QUESTIONS_RELATIONSHIP
	INNER JOIN QUESTIONS
	ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
	WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, surveyIdentifier); if err != nil {
		log.Println(err)
		utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	// Call "Close" function to the result set of the first SQL query.
	defer firstQuery.Close()

	// Parse the result set of the first SQL query.
	for firstQuery.Next() {
		// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
		var question models.AlphaQuestion

		// Call "Scan()" function to the result set of the second SQL query.
		if err := firstQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Execute SQL query to get information about all options of the specific question.
		secondQuery, err := database.DBSQL.Query(`SELECT
				OPTIONS.ID,
				OPTIONS.TEXT,
       			OPTIONS.POSITION
			FROM QUESTIONS_OPTIONS_RELATIONSHIP
			INNER JOIN OPTIONS
			ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
			WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
			log.Println(err)
			utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		// Parse the result set of the second SQL query.
		for secondQuery.Next() {
			// Variable "option" has been initialized by assigning it to a "Option" struct.
			var option models.Option

			// Call "Scan()" function to the result set of the second SQL query.
			if err := secondQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Append the information about option to the array.
			question.Options = append(question.Options, option)
		}

		// Call "Close" function to the result set of the second SQL query.
		secondQuery.Close()

		// Append information about question to the array.
		survey.Questions = append(survey.Questions, question)
	}

	// Send JSON response with status code "200".
	utils.Response(responseWriter, http.StatusOK, survey)
}

var GetAvailableSurvey = func(responseWriter http.ResponseWriter, request *http.Request) {
	// Variable has been initialized by assigning it to a "AlphaSurvey" struct.
	survey := models.AlphaSurvey{}

	// Variable has been initialized by assigning it a "SurveyEmployeeRelationship" struct.
	surveyEmployeeRelationship := models.SurveyEmployeeRelationship{}

	// Variable has been initialized by assigning it a array of URL parameters from the request.
	keys := request.URL.Query()

	// Check if an array contains any element.
	if len(keys) != 0 {
		// Variable has been initialized by assigning it a unique identifier of survey.
		surveyIdentifier := keys.Get("survey_id")

		// Variable has been initialized by assigning it a unique email.
		email := keys.Get("email")

		// Check key availability.
		if len(surveyIdentifier) != 0 && len(email) != 0 {
			// Check availability of the survey in the PostgreSQL database.
			if err := database.DBGORM.Where("ID = ?", surveyIdentifier).Find(&survey).Error; err != nil {
				utils.ResponseWithSuccess(responseWriter, http.StatusOK, "Survey not found.")
				return
			}

			// Проверка срока прохождения опроса.
			if survey.EndPeriod.Equal(time.Now()) || survey.EndPeriod.Before(time.Now()) {
				utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The survey timed out.")
				return
			}

			// Check if the survey is available to the employee.
			if err := database.DBGORM.Where("SURVEY_ID = ? AND EMPLOYEE = ? AND STATUS = FALSE", surveyIdentifier, email).Find(&surveyEmployeeRelationship).Error; err != nil {
				utils.ResponseWithSuccess(responseWriter, http.StatusOK, "The survey is not available to the employee.")
				return
			}

			// Execute the SQL query to get all question.
			firstQuery, err := database.DBSQL.Query(`SELECT
				QUESTIONS.ID,
				QUESTIONS.TEXT,
				QUESTIONS.WIDGET,
				QUESTIONS.REQUIRED,
				QUESTIONS.POSITION
			FROM SURVEYS_QUESTIONS_RELATIONSHIP
			INNER JOIN QUESTIONS
			ON SURVEYS_QUESTIONS_RELATIONSHIP.QUESTION_ID = QUESTIONS.ID
			WHERE SURVEYS_QUESTIONS_RELATIONSHIP.SURVEY_ID = $1;`, surveyIdentifier); if err != nil {
				log.Println(err)
				utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
				return
			}

			// Call "Close" function to the result set of the first SQL query.
			defer firstQuery.Close()

			// Parse the result set of the first SQL query.
			for firstQuery.Next() {
				// Variable "question" has been initialized by assigning it to a "AlphaQuestion" struct.
				var question models.AlphaQuestion

				// Call "Scan()" function to the result set of the second SQL query.
				if err := firstQuery.Scan(&question.ID, &question.Text, &question.Widget, &question.Required, &question.Position); err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Execute SQL query to get information about all options of the specific question.
				secondQuery, err := database.DBSQL.Query(`SELECT
					OPTIONS.ID,
				  	OPTIONS.TEXT,
       				OPTIONS.POSITION
				FROM QUESTIONS_OPTIONS_RELATIONSHIP
				INNER JOIN OPTIONS
				ON QUESTIONS_OPTIONS_RELATIONSHIP.OPTION_ID = OPTIONS.ID
				WHERE QUESTIONS_OPTIONS_RELATIONSHIP.QUESTION_ID = $1;`, question.ID); if err != nil {
					log.Println(err)
					utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
					return
				}

				// Parse the result set of the second SQL query.
				for secondQuery.Next() {
					// Variable "option" has been initialized by assigning it to a "Option" struct.
					var option models.Option

					// Call "Scan()" function to the result set of the second SQL query.
					if err := secondQuery.Scan(&option.ID, &option.Text, &option.Position); err != nil {
						log.Println(err)
						utils.ResponseWithError(responseWriter, http.StatusInternalServerError, err.Error())
						return
					}

					// Append the information about option to the array.
					question.Options = append(question.Options, option)
				}

				// Call "Close" function to the result set of the second SQL query.
				secondQuery.Close()

				// Append information about question to the array.
				survey.Questions = append(survey.Questions, question)
			}

			// Send JSON response with status code "200".
			utils.Response(responseWriter, http.StatusOK, survey)
		} else {
			utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
			return
		}
	} else {
		utils.ResponseWithError(responseWriter, http.StatusBadRequest, "http.StatusBadRequest")
		return
	}
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
