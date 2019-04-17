package models

type AlphaQuestion struct {
	ID int `gorm:"primary_key" json:"question_id"`
	Text *string `json:"question_text"`
	Widget *int `json:"widget"`
	Required *bool `json:"required"`
	Position *int `json:"question_position"`
	Category *int `json:"category"`
	Options []Option `json:"options"`
}

func (AlphaQuestion) TableName() string {
	return "questions"
}

type BetaQuestion struct {
	ID int `gorm:"primary_key" json:"question_id"`
	Text *string `json:"question_text"`
	Widget *int `json:"widget"`
	Required *bool `json:"required"`
	Position *int `json:"question_position"`
	Category *int `json:"category"`
}

func (BetaQuestion) TableName() string {
	return "questions"
}
