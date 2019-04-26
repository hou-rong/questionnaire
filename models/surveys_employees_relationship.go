package models

type SurveyEmployeeRelationship struct {
	SurveyID string `json:"survey_id"`
	Employee string `json:"employee"`
	Status bool `json:"status"`
	Send bool `json:"send"`
}

func (SurveyEmployeeRelationship) TableName() string {
	return "surveys_employees_relationship"
}
