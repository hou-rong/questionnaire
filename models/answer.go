package models

type SingleAnswer struct {
	SurveyID string `json:"survey_id"`
	Employee string `json:"employee"`
	QuestionID int `json:"question_id"`
	QuestionText string `json:"question_text"`
	OptionID int `json:"option_id"`
	OptionText string `json:"option_text"`
}

func (SingleAnswer) TableName() string {
	return "answers"
}

type MultipleAnswer struct {
	SurveyID string `json:"survey_id"`
	Employee string `json:"employee"`
	QuestionID []int `json:"question_id"`
	QuestionText []string `json:"question_text"`
	OptionID []int `json:"option_id"`
	OptionText []string `json:"option_text"`
}

func (MultipleAnswer) TableName() string {
	return "answers"
}
