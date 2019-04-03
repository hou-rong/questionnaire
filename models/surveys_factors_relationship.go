package models

type SurveyFactorRelationship struct {
	SurveyID string `json:"survey_id"`
	FactorID int `json:"factor_id"`
}

func (SurveyFactorRelationship) TableName() string {
	return "surveys_factors_relationship"
}
