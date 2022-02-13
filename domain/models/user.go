package models

import (
	"time"
)

// Client struct
type User struct {
	ID              int       `gorm:"unique;not null"`
	Email           string    `gorm:"unique;not null"`
	CryptedPassword string    `gorm:"not null"`
	TotpSecret string
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdatedAt       time.Time `json:"UpdatedAt"`
}
