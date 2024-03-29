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
	router.HandleFunc("/api/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/api/category/{category_id:[0-9]+}", controllers.GetCategory).Methods("GET")
	router.HandleFunc("/api/category", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/category/{category_id:[0-9]+}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/api/category/{category_id:[0-9]+}", controllers.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/api/conditions", controllers.GetConditions).Methods("GET")
	router.HandleFunc("/api/condition/{condition_id:[0-9]+}", controllers.GetCondition).Methods("GET")
	router.HandleFunc("/api/condition", controllers.CreateCondition).Methods("POST")
	router.HandleFunc("/api/condition/{condition_id:[0-9]+}", controllers.UpdateCondition).Methods("PUT")
	router.HandleFunc("/api/condition/{condition_id:[0-9]+}", controllers.DeleteCondition).Methods("DELETE")

	router.HandleFunc("/api/options", controllers.GetOptions).Methods("GET")
	router.HandleFunc("/api/option/{option_id:[0-9]+}", controllers.GetOption).Methods("GET")
	router.HandleFunc("/api/option", controllers.CreateOption).Methods("POST")
	router.HandleFunc("/api/option/{option_id:[0-9]+}", controllers.UpdateOption).Methods("PUT")
	router.HandleFunc("/api/option/{option_id:[0-9]+}", controllers.DeleteOption).Methods("DELETE")

	router.HandleFunc("/api/widgets", controllers.GetWidgets).Methods("GET")
	router.HandleFunc("/api/widget/{widget_id:[0-9]+}", controllers.GetWidget).Methods("GET")
	router.HandleFunc("/api/widget", controllers.CreateWidget).Methods("POST")
	router.HandleFunc("/api/widget/{widget_id:[0-9]+}", controllers.UpdateWidget).Methods("PUT")
	router.HandleFunc("/api/widget/{widget_id:[0-9]+}", controllers.DeleteWidget).Methods("DELETE")

	router.HandleFunc("/api/beta/surveys", controllers.GetBetaSurveys).Methods("GET")
	router.HandleFunc("/api/alpha/surveys", controllers.GetAlphaSurveys).Methods("GET")
	router.HandleFunc("/api/beta/survey/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.GetBetaSurvey).Methods("GET")
	router.HandleFunc("/api/alpha/survey/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.GetAlphaSurvey).Methods("GET")
	router.HandleFunc("/api/survey", controllers.CreateSurvey).Methods("POST")
	router.HandleFunc("/api/survey/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.UpdateSurvey).Methods("PUT")
	router.HandleFunc("/api/survey/{survey_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", controllers.DeleteSurvey).Methods("DELETE")
	router.HandleFunc("/api/available/survey", controllers.GetAvailableSurvey).Methods("GET")

	router.HandleFunc("/api/beta/questions", controllers.GetBetaQuestions).Methods("GET")
	router.HandleFunc("/api/alpha/questions", controllers.GetAlphaQuestions).Methods("GET")
	router.HandleFunc("/api/beta/question/{question_id:[0-9]+}", controllers.GetBetaQuestion).Methods("GET")
	router.HandleFunc("/api/alpha/question/{question_id:[0-9]+}", controllers.GetAlphaQuestion).Methods("GET")
	router.HandleFunc("/api/question", controllers.CreateQuestion).Methods("POST")
	router.HandleFunc("/api/question/{question_id:[0-9]+}", controllers.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/api/question/{question_id:[0-9]+}", controllers.DeleteQuestion).Methods("DELETE")

	router.HandleFunc("/api/beta/factors", controllers.GetBetaFactors).Methods("GET")
	router.HandleFunc("/api/alpha/factors", controllers.GetAlphaFactors).Methods("GET")
	router.HandleFunc("/api/beta/factor/{factor_id:[0-9]+}", controllers.GetBetaFactor).Methods("GET")
	router.HandleFunc("/api/alpha/factor/{factor_id:[0-9]+}", controllers.GetAlphaFactor).Methods("GET")
	router.HandleFunc("/api/factor", controllers.CreateFactor).Methods("POST")
	router.HandleFunc("/api/factor/{factor_id:[0-9]+}", controllers.UpdateFactor).Methods("PUT")
	router.HandleFunc("/api/factor/{factor_id:[0-9]+}", controllers.DeleteFactor).Methods("DELETE")
	router.HandleFunc("/api/factor/in/{factor_id:[0-9]+}", controllers.DeleteInFactor).Methods("DELETE")
	router.HandleFunc("/api/factor/out/{factor_id:[0-9]+}", controllers.DeleteOutFactor).Methods("DELETE")

	router.HandleFunc("/api/organizations", controllers.GetOrganizations).Methods("GET")

	router.HandleFunc("/api/beta/surveys_factors_relationship", controllers.GetBetaSurveysFactorsRelationship).Methods("GET")
	router.HandleFunc("/api/single/survey_factor_relationship", controllers.CreateSingleSurveyFactorRelationship).Methods("POST")
	router.HandleFunc("/api/multiple/survey_factor_relationship", controllers.CreateMultipleSurveyFactorRelationship).Methods("POST")
	router.HandleFunc("/api/single/survey_factor_relationship", controllers.DeleteSingleSurveyFactorRelationship).Methods("DELETE")
	router.HandleFunc("/api/multiple/survey_factor_relationship", controllers.DeleteMultipleSurveyFactorRelationship).Methods("DELETE")
	router.HandleFunc("/api/check/survey_factor_relationship", controllers.CheckSurveyFactorRelationship).Methods("GET")

	router.HandleFunc("/api/beta/surveys_organizations_relationship", controllers.GetBetaSurveysOrganizationsRelationship).Methods("GET")
	router.HandleFunc("/api/single/survey_organization_relationship", controllers.CreateSingleSurveyOrganizationRelationship).Methods("POST")
	router.HandleFunc("/api/multiple/survey_organization_relationship", controllers.CreateMultipleSurveyOrganizationRelationship).Methods("POST")
	router.HandleFunc("/api/single/survey_organization_relationship", controllers.DeleteSingleSurveyOrganizationRelationship).Methods("DELETE")
	router.HandleFunc("/api/multiple/survey_organization_relationship", controllers.DeleteMultipleSurveyOrganizationRelationship).Methods("DELETE")

	router.HandleFunc("/api/beta/factors_questions_relationship", controllers.GetBetaFactorsQuestionsRelationship).Methods("GET")
	router.HandleFunc("/api/single/int/factor_question_relationship", controllers.CreateSingleIntFactorQuestionRelationship).Methods("POST")
	router.HandleFunc("/api/multiple/int/factor_question_relationship", controllers.CreateMultipleIntFactorQuestionRelationship).Methods("POST")
	router.HandleFunc("/api/multiple/text/factor_question_relationship", controllers.CreateMultipleTextFactorQuestionRelationship).Methods("POST")
	router.HandleFunc("/api/single/int/factor_question_relationship", controllers.DeleteSingleIntFactorQuestionRelationship).Methods("DELETE")
	router.HandleFunc("/api/multiple/int/factor_question_relationship", controllers.DeleteMultipleIntFactorQuestionRelationship).Methods("DELETE")

	router.HandleFunc("/api/questions_options_relationship", controllers.GetQuestionsOptionsRelationship).Methods("GET")
	router.HandleFunc("/api/single/int/question_option_relationship", controllers.CreateSingleIntQuestionOptionRelationship).Methods("POST")
	router.HandleFunc("/api/multiple/int/question_option_relationship", controllers.CreateMultipleIntQuestionOptionRelationship).Methods("POST")
	router.HandleFunc("/api/multiple/text/question_option_relationship", controllers.CreateMultipleTextQuestionOptionRelationship).Methods("POST")
	router.HandleFunc("/api/single/int/question_option_relationship", controllers.DeleteSingleIntQuestionOptionRelationship).Methods("DELETE")
	router.HandleFunc("/api/multiple/int/question_option_relationship", controllers.DeleteMultipleIntQuestionOptionRelationship).Methods("DELETE")

	router.HandleFunc("/api/beta/surveys_questions_relationship", controllers.GetBetaSurveysQuestionsRelationship).Methods("GET")
	router.HandleFunc("/api/single/survey_question_relationship", controllers.CreateSingleSurveyQuestionRelationship).Methods("POST")
	router.HandleFunc("/api/multiple/survey_question_relationship", controllers.CreateMultipleSurveyQuestionRelationship).Methods("POST")
	router.HandleFunc("/api/single/survey_question_relationship", controllers.DeleteSingleSurveyQuestionRelationship).Methods("DELETE")
	router.HandleFunc("/api/multiple/survey_question_relationship", controllers.DeleteMultipleSurveyQuestionRelationship).Methods("DELETE")
	router.HandleFunc("/api/check/survey_question_relationship", controllers.CheckSurveyQuestionRelationship).Methods("GET")

	router.HandleFunc("/api/single/survey_employee_relationship", controllers.CreateSingleSurveyEmployeeRelationship).Methods("POST")
	router.HandleFunc("/api/multiple/survey_employee_relationship", controllers.CreateMultipleSurveyEmployeeRelationship).Methods("POST")
	router.HandleFunc("/api/single/survey_employee_relationship", controllers.DeleteSingleSurveyEmployeeRelationship).Methods("DELETE")
	router.HandleFunc("/api/multiple/survey_employee_relationship", controllers.DeleteMultipleSurveyEmployeeRelationship).Methods("DELETE")
	router.HandleFunc("/api/single/survey_employee_relationship", controllers.UpdateSingleSurveyEmployeeRelationship).Methods("PUT")

	router.HandleFunc("/api/multiple/answer", controllers.CreateMultipleAnswer).Methods("POST")
}
