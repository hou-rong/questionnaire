package models

type SurveyOrganizationRelationship struct {
	SurveyID string `json:"survey_id"`
	OrganizationID int `json:"organization_id"`
}

func (SurveyOrganizationRelationship) TableName() string {
	return "surveys_organizations_relationship"
}
