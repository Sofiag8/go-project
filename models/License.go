package models

import "time"

type License struct {
	Bims      []Bim `gorm:"foreignKey:LicenseId"`
	ID        uint  `gorm:"primaryKey"`
	License   string
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
