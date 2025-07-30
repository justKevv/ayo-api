package models

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	MatchID uint `gorm:"not null"`
	PlayerID uint `gorm:"not null"`
	Minute int `gorm:"not null"`
	Match Match `gorm:"foreignKey:MatchID"`
	Player Player `gorm:"foreignKey:PlayerID"`
}
