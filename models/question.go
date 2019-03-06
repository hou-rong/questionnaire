package models

type Question struct {
	ID int `gorm:"primary_key" json:"question_id"`
	Text string `json:"question_text"`
	Widget Widget `json:"widget"`
	Options []Option `json:"options"`
}
