package models

import (
	"time"
)

// GORM prefer convention over configuration,
// by default, GORM uses ID as primary key,
// pluralize struct name to snake_cases as table name,
// snake_case as column name, and uses CreatedAt, UpdatedAt to track creating/updating time
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
