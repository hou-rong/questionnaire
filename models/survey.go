package models

import "time"

type BetaSurvey struct {
	ID string `gorm:"primary_key" json:"survey_id"`
	Name *string `json:"survey_name"`
	Description *string `json:"survey_description"`
	Category *int `json:"category"`
	Condition *int `json:"condition"`
	Mark *bool `json:"mark"`
	Control *bool `json:"control"`
	StartPeriod time.Time `json:"start_period"`
	EndPeriod time.Time `json:"end_period"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Email *string `json:"email"`
	Blocked bool `json:"blocked"`
}

func (BetaSurvey) TableName() string {
	return "surveys"
}
