package routes

import (
	"github.com/gorilla/mux"
	"questionnaire/controllers"
)

/*
Function name:
"Handle"

Function description:
The function provides a list of all available URLs which would use RESTful web service.
*/

func Handle(router *mux.Router) {
	router.HandleFunc("/api/extended_factors", controllers.GetExtendedFactors).Methods("GET")
	router.HandleFunc("/api/abridged_factors", controllers.GetAbridgedFactors).Methods("GET")
	router.HandleFunc("/api/factor", controllers.CreateFactor).Methods("POST")
	router.HandleFunc("/api/factor/{factor_id:[0-9]+}", controllers.GetFactor).Methods("GET")
	router.HandleFunc("/api/factor/{factor_id:[0-9]+}", controllers.DeleteFactor).Methods("DELETE")
	router.HandleFunc("/api/factor/{factor_id:[0-9]+}", controllers.UpdateFactor).Methods("PUT")

	router.HandleFunc("/api/options", controllers.GetOptions).Methods("GET")
	router.HandleFunc("/api/option", controllers.CreateOption).Methods("POST")
	router.HandleFunc("/api/option/{option_id:[0-9]+}", controllers.GetOption).Methods("GET")
	router.HandleFunc("/api/option/{option_id:[0-9]+}", controllers.DeleteOption).Methods("DELETE")
	router.HandleFunc("/api/option/{option_id:[0-9]+}", controllers.UpdateOption).Methods("PUT")

	router.HandleFunc("/api/widgets", controllers.GetWidgets).Methods("GET")
	router.HandleFunc("/api/widget", controllers.CreateWidget).Methods("POST")
	router.HandleFunc("/api/widget/{widget_id:[0-9]+}", controllers.GetWidget).Methods("GET")
	router.HandleFunc("/api/widget/{widget_id:[0-9]+}", controllers.DeleteWidget).Methods("DELETE")
	router.HandleFunc("/api/widget/{widget_id:[0-9]+}", controllers.UpdateWidget).Methods("PUT")

	router.HandleFunc("/api/questions", controllers.GetQuestions).Methods("GET")
	router.HandleFunc("/api/question", controllers.CreateQuestion).Methods("POST")
	router.HandleFunc("/api/question/{question_id:[0-9]+}", controllers.GetQuestion).Methods("GET")
	router.HandleFunc("/api/question/{question_id:[0-9]+}", controllers.DeleteQuestion).Methods("DELETE")
	router.HandleFunc("/api/question/{question_id:[0-9]+}", controllers.UpdateQuestion).Methods("PUT")

	router.HandleFunc("/api/extended_surveys", controllers.GetExtendedSurveys).Methods("GET")
	router.HandleFunc("/api/abridged_surveys", controllers.GetAbridgedSurveys).Methods("GET")
	router.HandleFunc("/api/abridged_surveys/active", controllers.GetAbridgedActiveSurveys).Methods("GET")
	router.HandleFunc("/api/abridged_surveys/inactive", controllers.GetAbridgedInactiveSurveys).Methods("GET")
	router.HandleFunc("/api/survey", controllers.CreateSurvey).Methods("POST")
	router.HandleFunc("/api/survey/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.GetSurvey).Methods("GET")
	router.HandleFunc("/api/survey/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.DeleteSurvey).Methods("DELETE")
	router.HandleFunc("/api/survey/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.UpdateSurvey).Methods("PUT")

	router.HandleFunc("/api/surveys_factors_relationship/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.DeleteSurveysFactorsRelationship).Methods("DELETE")
	router.HandleFunc("/api/surveys_factors_relationship/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.CreateSurveysFactorsRelationship).Methods("POST")

	router.HandleFunc("/api/surveys_questions_relationship/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.DeleteSurveysQuestionsRelationship).Methods("DELETE")
	router.HandleFunc("/api/surveys_questions_relationship/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.CreateSurveysQuestionsRelationship).Methods("POST")

	router.HandleFunc("/api/employees", controllers.GetEmployees).Methods("GET")
	router.HandleFunc("/api/employee/{email}", controllers.GetEmployee).Methods("GET")
	router.HandleFunc("/api/organization_employees", controllers.GetOrganizationEmployees).Methods("GET")
	router.HandleFunc("/api/organizations", controllers.GetOrganizations).Methods("GET")
}
