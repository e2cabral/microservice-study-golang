package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string     `gorm:"type:varchar" json:"name"`
	Description string     `gorm:"type:varchar" json:"description"`
	Categories  []Category `gorm:"many2many:products_categories" json:"category"`
	Reviews     []Review   `gorm:"foreignKey:ProductID" json:"reviews"`
	ID          uint       `gorm:"primaryKey; type:int" json:"id"`
	Price       float64    `gorm:"type:float" json:"price"`
}
