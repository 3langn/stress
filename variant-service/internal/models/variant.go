package models

type Variant struct {
	ID    int64  `json:"id" gorm:"primary_key;auto_increment"`
	Name  string `json:"name" gorm:"type:varchar(255);not null"`
	Value string `json:"value" gorm:"type:varchar(255);not null"`
}

func (Variant) TableName() string {
	return "variants"
}
