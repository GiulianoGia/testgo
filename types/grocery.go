package types

import (
	"github.com/google/uuid"
)

type Grocery struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string    `json: "name"`
	Password string    `json: "password"`
	Age      int       `json: "quantity"`
}

type UserGrocery struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	GroceryID int       `gorm:"primaryKey"`
}
