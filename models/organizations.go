package models

type Organization struct {
	ID int `json:"organization_id"`
	Name string `json:"organization_name"`
	Rang int `json:"organization_rang"`
	Children []*Organization `json:"children"`
}
