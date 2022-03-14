package models

import (
	"time"
)

type User struct {
	Licenses       []License `gorm:"foreignKey:UserId"`
	ID             uint      `gorm:"primaryKey"`
	FullName       string
	Company        string
	Email          string
	CompanyAddress string
	Phone          string
	City           string
	Project        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
