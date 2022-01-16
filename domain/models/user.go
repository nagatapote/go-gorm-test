package models

import (
	"time"
)

// Client struct
type User struct {
	ID             int       `gorm:"unique;not null"`
	Email          string    `gorm:"unique;not null"`
	PasswordDigest string    `gorm:"unique;not null"`
	CreatedAt      time.Time `json:"CreatedAt"`
	UpdatedAt      time.Time `json:"UpdatedAt"`
}
