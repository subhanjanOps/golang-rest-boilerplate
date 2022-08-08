package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName       string    `gorm:"not null"`
	LastName        string    `gorm:"not null"`
	Fullname        string    `gorm:"not null"`
	Username        string    `gorm:"not null"`
	Email           string    `gorm:"not null"`
	DateOfBirth     time.Time `gorm:"not null"`
	Phone           string    `gorm:"not null"`
	Password        string    `gorm:"not null"`
	Status          int       `gorm:"not null;default:0"`
	EmailVerifiedAt sql.NullTime
}
