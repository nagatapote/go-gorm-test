package models

import (
	"time"
)

// Client struct
type File struct {
	ID         int       `gorm:"unique;not null"`
	UploadName string    `gorm:"unique;not null"`
	FilesName  string    `gorm:"not null"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
}
