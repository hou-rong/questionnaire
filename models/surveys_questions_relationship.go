package models

type SurveyQuestionRelationship struct {
	SurveyID string `json:"survey_id"`
	QuestionID int `json:"question_id"`
}

func (SurveyQuestionRelationship) TableName() string {
	return "surveys_questions_relationship"
}
