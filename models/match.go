package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Date time.Time `gorm:"not null"`
	Time string `gorm:"type:varchar(5);not null"`
	Team1ID uint `gorm:"not null"`
	Team2ID uint `gorm:"not null"`
	Team1Score int `gorm:"default:0"`
	Team2Score int `gorm:"default:0"`
	Status string `gorm:"type:varchar(20);default:'scheduled'"`
	Team1 Team `gorm:"foreignKey:Team1ID"`
	Team2 Team `gorm:"foreignKey:Team2ID"`
	Goals []Goal `gorm:"foreignKey:MatchID"`
}
