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

func (p *Player) BeforeCreate(tx *gorm.DB) error {
    var count int64
    tx.Model(&Player{}).Where("team_id = ? AND jersey_number = ?", p.TeamID, p.JerseyNumber).Count(&count)
    if count > 0 {
        return gorm.ErrDuplicatedKey
    }
    return nil
}
