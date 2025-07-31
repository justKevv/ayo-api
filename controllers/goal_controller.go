package controllers

import (
    "gin/models"
    "gin/utils"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func CreateGoal(db *gorm.DB) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var goal models.Goal
        if err := ctx.ShouldBindJSON(&goal); err != nil {
            utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request body")
            return
        }

        // Validate match exists
        var match models.Match
        if err := db.First(&match, goal.MatchID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusNotFound, "Match not found")
            return
        }

        // Validate player exists
        var player models.Player
        if err := db.First(&player, goal.PlayerID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusNotFound, "Player not found")
            return
        }

        // Validate team exists
        var team models.Team
        if err := db.First(&team, goal.TeamID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusNotFound, "Team not found")
            return
        }

        // Validate player belongs to the team
        if player.TeamID != goal.TeamID {
            utils.RespondWithError(ctx, http.StatusBadRequest, "Player does not belong to the specified team")
            return
        }

        // Validate team is playing in the match
        if goal.TeamID != match.Team1ID && goal.TeamID != match.Team2ID {
            utils.RespondWithError(ctx, http.StatusBadRequest, "Team is not playing in this match")
            return
        }

        if err := db.Create(&goal).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to create goal")
            return
        }

        // Update match score
        if goal.TeamID == match.Team1ID {
            match.Team1Score++
        } else {
            match.Team2Score++
        }
        db.Save(&match)

        // Reload with related data
        if err := db.Preload("Match").Preload("Player").Preload("Team").First(&goal, goal.ID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve created goal")
            return
        }

        utils.RespondWithSuccess(ctx, http.StatusCreated, "Goal created successfully", goal)
    }
}

func GetMatchGoals(db *gorm.DB) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var goals []models.Goal
        matchID := ctx.Param("id")

        if err := db.Preload("Player").Preload("Team").Where("match_id = ?", matchID).Find(&goals).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to fetch goals")
            return
        }

        utils.RespondWithSuccess(ctx, http.StatusOK, "Goals retrieved successfully", goals)
    }
}
