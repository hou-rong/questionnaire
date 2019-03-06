package models

import "time"

type Survey struct {
	ID string `gorm:"primary_key" json:"survey_id"`
	Name string `json:"survey_name"`
	Description *string `json:"survey_description"`
	Status bool `json:"survey_status"`
	StartPeriod *time.Time `json:"start_period"`
	EndPeriod *time.Time `json:"end_period"`
	Questions []Question `json:"questions"`
}
