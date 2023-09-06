package types

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID   int    `gorm:"primaryKey"`
	Name string `json: "name"`
}

type RoleEnum int

const (
	Admin  RoleEnum = 1
	Member RoleEnum = 2
)

func GetRole(r RoleEnum) string {
	switch r {
	case Admin:
		return "Admin"
	case Member:
		return "Member"
	default:
		return "unkown"
	}
}

func GetRoleId(r string) int {
	switch r {
	case "Admin":
		return 1
	case "Member":
		return 2
	default:
		return 0
	}
}
