package models

type Condition struct {
	ID int `gorm:"primary_key" json:"condition_id"`
	Name string `json:"condition_name"`
}
