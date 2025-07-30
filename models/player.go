package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Height float64 `gorm:"not null"`
	Weight float64 `gorm:"not null"`
	Position string `gorm:"not null"`
	TeamID uint `gorm:"not null"`
	JerseyNumber uint `gorm:"not null;uniqueIndex:idx_team_jersey"`
	Team Team `gorm:"foreignKey:TeamID"`
	Goals []Goal `gorm:"foreignKey:PlayerID"`
}
