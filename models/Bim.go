package models

import "time"

type Bim struct {
	ID        uint `gorm:"primaryKey"`
	Guid      string
	LicenseId int
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
