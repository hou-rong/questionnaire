package models

type FactorQuestionRelationship struct {
	FactorID int `json:"factor_id"`
	QuestionID int `json:"question_id"`
}

func (FactorQuestionRelationship) TableName() string {
	return "factors_questions_relationship"
}
