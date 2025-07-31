package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Height float64 `gorm:"not null" json:"height"`
	Weight float64 `gorm:"not null" json:"weight"`
	Position string `gorm:"not null" json:"position"`
	TeamID uint `gorm:"not null" json:"team_id"`
	JerseyNumber uint `gorm:"not null;uniqueIndex:idx_team_jersey" json:"jersey_number"`
	Team *Team `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Goals []Goal `gorm:"foreignKey:PlayerID" json:"goals,omitempty"`
}
