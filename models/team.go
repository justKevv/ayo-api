package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
	Logo string `gorm:"not null" json:"logo"`
	YearEstablished int `gorm:"not null" json:"year_established"`
	Address string `gorm:"not null" json:"address"`
	City string `gorm:"not null" json:"city"`
	Players []Player `gorm:"foreignKey:TeamID" json:"players,omitempty"`
}
