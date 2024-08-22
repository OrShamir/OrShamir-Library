package postgres

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Role     string // "customer" or "employee"
	Password string // Consider hashing in a real application
}
