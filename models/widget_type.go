package models

type WidgetType struct {
	WidgetTypeID   int    `gorm:"primary_key" json:"widget_type_id"`
	WidgetTypeName string `gorm:"not null;unique" json:"widget_type_name"`
}

func (WidgetType) TableName() string {
	return "widget_type"
}
