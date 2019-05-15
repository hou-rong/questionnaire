package models

import "time"

type Organization struct {
	ID int `json:"organization_id"`
	Name string `json:"organization_name"`
	Rang int `json:"organization_rang"`
	Children []*Organization `json:"children"`
}

type OrganizationalStructure struct {
	OrgStructureVersionID int `json:"org_structure_version_id"`
	VerDFrom time.Time `json:"ver_d_from"`
	VerPrevDFrom time.Time `json:"ver_prev_d_from"`
	VerPrevDTo time.Time `json:"ver_prev_d_to"`
	OrganizationID int `json:"organization_id"`
	ParentOrganizationID int `json:"parent_organization_id"`
	OrganizationName string `json:"organization_name"`
	OrganizationRang int `json:"organization_rang"`
	TreeOrganizationID string `json:"tree_organization_id"`
	TreeOrganizationName string `json:"tree_organization_name"`
	Rang1OrganizationID int `json:"rang1_organization_id"`
	Rang1OrganizationName string `json:"rang1_organization_name"`
	CreationDate time.Time `json:"creation_date"`
}
