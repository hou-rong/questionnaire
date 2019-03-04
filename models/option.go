package models

type Option struct {
	OptionID   int    `gorm:"primary_key" json:"option_id"`
	OptionText string `gorm:"not null;unique" json:"option_text"`
}

func (Option) TableName() string {
	return "option"
}
