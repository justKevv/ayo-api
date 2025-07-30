package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null;unique"`
	Logo string `gorm:"not null"`
	YearEstablished int `gorm:"not null"`
	Address string `gorm:"not null"`
	City string `gorm:"not null"`
	Players []Player `gorm:"foreignKey:TeamID"`
}
