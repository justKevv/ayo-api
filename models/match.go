package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	Date time.Time `gorm:"not null" json:"date"`
	Time string `gorm:"type:varchar(5);not null" json:"time"`
	Team1ID uint `gorm:"not null" json:"team1_id"`
	Team2ID uint `gorm:"not null" json:"team2_id"`
	Team1Score int `gorm:"default:0" json:"team1_score"`
	Team2Score int `gorm:"default:0" json:"team2_score"`
	Status string `gorm:"type:varchar(20);default:'scheduled'" json:"status"`
	Team1 Team `gorm:"foreignKey:Team1ID" json:"team1,omitempty"`
	Team2 Team `gorm:"foreignKey:Team2ID" json:"team2,omitempty"`
	Goals []Goal `gorm:"foreignKey:MatchID" json:"goals,omitempty"`
}
