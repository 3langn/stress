package models

type Rating struct {
	RatingStar  float64 `json:"rating_star"`
	RatingCount []int   `json:"rating_count"`
}

type TierVariation struct {
	Images      []string `json:"images"`
	Name        string   `json:"name"`
	Options     []string `json:"options"`
	SummedStock int      `json:"summed_stock"`
}

type Category struct {
	ID       int        `json:"id" gorm:"primary_key;auto_increment"`
	Name     string     `json:"name" gorm:"type:varchar(255);not null"`
	Sub      []Category `json:"sub" gorm:"null;foreignkey:ParentID"`
	ParentID *int       `json:"parent_id"`
}
