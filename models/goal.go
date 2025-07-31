package models

import "gorm.io/gorm"

type Goal struct {
    gorm.Model
    MatchID    uint      `json:"match_id" gorm:"not null"`
    PlayerID   uint      `json:"player_id" gorm:"not null"`
    TeamID     uint      `json:"team_id" gorm:"not null"`
    GoalTime   int       `json:"goal_time" gorm:"not null"` // Goal time in minutes
    GoalType   string    `json:"goal_type" gorm:"type:varchar(20);default:'normal'"` // normal, penalty, own_goal
    Match      *Match  `json:"match,omitempty" gorm:"foreignKey:MatchID"`
    Player     *Player `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
    Team       *Team   `json:"team,omitempty" gorm:"foreignKey:TeamID"`
}
