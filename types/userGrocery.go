package types

import "github.com/google/uuid"

type UserGrocery struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	GroceryID int       `gorm:"primaryKey"`
}
