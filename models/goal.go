package models

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	MatchID uint `gorm:"not null" json:"match_id"`
	PlayerID uint `gorm:"not null" json:"player_id"`
	Minute int `gorm:"not null" json:"minute"`
	Match Match `gorm:"foreignKey:MatchID" json:"match,omitempty"`
	Player Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}
