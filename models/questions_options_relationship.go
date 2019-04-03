package models

type QuestionOptionRelationship struct {
	QuestionID int `json:"question_id"`
	OptionID int `json:"option_id"`
}

func (QuestionOptionRelationship) TableName() string {
	return "questions_options_relationship"
}
