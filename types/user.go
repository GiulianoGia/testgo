package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string    `json: "name"`
	Password string    `json: "password"`
	Age      int       `json: "quantity"`
	RoleID   int
	Role     Role
}
