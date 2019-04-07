package models

type Option struct {
	ID int `gorm:"primary_key" json:"option_id"`
	Text *string `json:"option_text"`
	Position *int `json:"option_position"`
}
