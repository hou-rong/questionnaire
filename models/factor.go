package models

type ExtendedFactor struct {
	ID int `gorm:"primary_key" json:"factor_id"`
	Name string `json:"factor_name"`
	Questions []Question `json:"questions"`
}

func (ExtendedFactor) TableName() string {
	return "factors"
}

type AbridgedFactor struct {
	ID int `gorm:"primary_key" json:"factor_id"`
	Name string `json:"factor_name"`
}

func (AbridgedFactor) TableName() string {
	return "factors"
}
