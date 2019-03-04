package models

type Factor struct {
	FactorID   int    `gorm:"primary_key" json:"factor_id"`
	FactorName string `gorm:"not null;unique" json:"factor_name"`
}
