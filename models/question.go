package models

type BetaQuestion struct {
	ID int `gorm:"primary_key" json:"question_id"`
	Text *string `json:"question_text"`
	Widget *int `json:"widget"`
}

func (BetaQuestion) TableName() string {
	return "questions"
}
