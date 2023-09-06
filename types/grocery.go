package types

import "gorm.io/gorm"

type Grocery struct {
	gorm.Model
	ID        uint        `gorm:"primaryKey"`
	Name      string      `json: "name"`
	Quantity  int         `json: "quantity"`
	Done      bool        `json: "done"`
	DeletedAt interface{} `gorm:"-"` // Ignore DeletedAt
}
