package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `gorm:"type:varchar" json:"name"`
	Description string    `gorm:"type:varchar" json:"description"`
	Products    []Product `gorm:"many2many:products_categories" json:"products"`
}
