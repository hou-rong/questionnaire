package models

type Factor struct {
	ID int `gorm:"primary_key" json:"factor_id"`
	Name string `json:"factor_name"`
	Questions []Question `json:"questions"`
}
