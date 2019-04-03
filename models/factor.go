package models

type BetaFactor struct {
	ID int `gorm:"primary_key" json:"factor_id"`
	Name string `json:"factor_name"`
}

func (BetaFactor) TableName() string {
	return "factors"
}
