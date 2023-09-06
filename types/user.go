package types

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string    `json: "name"`
	Password string    `json: "password"`
	Age      int       `json: "age"`
	RoleID   int
	Role     Role
}
