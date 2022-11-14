package entities

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID   uint    `gorm:"type:int" json:"product_id"`
	Product     Product `gorm:"foreignKey:ProductID" json:"product"`
	UserID      uint    `gorm:"type:int" json:"user_id"`
	Rating      int     `gorm:"type:int" json:"rating"`
	Description string  `gorm:"type:varchar" json:"description"`
}
