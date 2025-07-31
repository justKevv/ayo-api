package controllers

import (
    "gin/models"
    "gin/utils"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type MatchReport struct {
    models.Match
    MatchStatus    string      `json:"match_status"` // Home Win/Away Win/Draw
    TopGoalScorer  *PlayerGoalCount `json:"top_goal_scorer,omitempty"`
    Team1Wins      int64       `json:"team1_cumulative_wins"`
    Team2Wins      int64       `json:"team2_cumulative_wins"`
}

type PlayerGoalCount struct {
    PlayerID   uint   `json:"player_id"`
    PlayerName string `json:"player_name"`
    GoalCount  int64  `json:"goal_count"`
}

func GetMatchReport(db *gorm.DB) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var match models.Match
        if err := db.Preload("Team1").Preload("Team2").Preload("Goals.Player").First(&match, ctx.Param("id")).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusNotFound, "Match not found")
            return
        }

        // Determine match status
        var matchStatus string
        if match.Team1Score > match.Team2Score {
            matchStatus = "Home Win"
        } else if match.Team2Score > match.Team1Score {
            matchStatus = "Away Win"
        } else {
            matchStatus = "Draw"
        }

        // Get top goal scorer for this match
        var topScorer *PlayerGoalCount
        var playerGoals []struct {
            PlayerID   uint   `json:"player_id"`
            PlayerName string `json:"player_name"`
            GoalCount  int64  `json:"goal_count"`
        }

        db.Table("goals").
            Select("goals.player_id, players.name as player_name, COUNT(*) as goal_count").
            Joins("JOIN players ON goals.player_id = players.id").
            Where("goals.match_id = ?", match.ID).
            Group("goals.player_id, players.name").
            Order("goal_count DESC").
            Limit(1).
            Find(&playerGoals)

        if len(playerGoals) > 0 {
            topScorer = &PlayerGoalCount{
                PlayerID:   playerGoals[0].PlayerID,
                PlayerName: playerGoals[0].PlayerName,
                GoalCount:  playerGoals[0].GoalCount,
            }
        }

        var team1Wins, team2Wins int64
        db.Model(&models.Match{}).Where("(team1_id = ? AND team1_score > team2_score) OR (team2_id = ? AND team2_score > team1_score)", match.Team1ID, match.Team1ID).Count(&team1Wins)
        db.Model(&models.Match{}).Where("(team1_id = ? AND team1_score > team2_score) OR (team2_id = ? AND team2_score > team1_score)", match.Team2ID, match.Team2ID).Count(&team2Wins)

        report := MatchReport{
            Match:         match,
            MatchStatus:   matchStatus,
            TopGoalScorer: topScorer,
            Team1Wins:     team1Wins,
            Team2Wins:     team2Wins,
        }

        utils.RespondWithSuccess(ctx, http.StatusOK, "Match report retrieved successfully", report)
    }
}
