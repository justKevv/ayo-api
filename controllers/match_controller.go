package controllers

import (
	"gin/models"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllMatches(db *gorm.DB) gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		var matches []models.Match
		if err := db.Preload("Team1").Preload("Team2").Preload("Goals.Player").Find(&matches).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to fetch matches")
			return
		}
		utils.RespondWithSuccess(ctx, http.StatusOK, "Matches retrieved successfully", matches)
	}
}

func ShowMatch(db *gorm.DB) gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		var match models.Match
		if err := db.Preload("Team1").Preload("Team2").Preload("Goals.Player").First(&match, ctx.Param("id")).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusNotFound, "Match not found")
			return
		}
		utils.RespondWithSuccess(ctx, http.StatusOK, "Match retrieved successfully", match)
	}
}

func CreateMatch(db *gorm.DB) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var match models.Match
        if err := ctx.ShouldBindJSON(&match); err != nil {
            utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request body")
            return
        }

        // Validate teams exist
        var team1, team2 models.Team
        if err := db.First(&team1, match.Team1ID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusNotFound, "Team 1 not found")
            return
        }
        if err := db.First(&team2, match.Team2ID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusNotFound, "Team 2 not found")
            return
        }

        // Validate teams are different
        if match.Team1ID == match.Team2ID {
            utils.RespondWithError(ctx, http.StatusBadRequest, "Team cannot play against itself")
            return
        }

        if err := db.Create(&match).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to create match")
            return
        }

        // Reload with team data
        if err := db.Preload("Team1").Preload("Team2").First(&match, match.ID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve created match")
            return
        }

        utils.RespondWithSuccess(ctx, http.StatusCreated, "Match created successfully", match)
    }
}

func UpdateMatchScore(db *gorm.DB) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var match models.Match
        if err := db.First(&match, ctx.Param("id")).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusNotFound, "Match not found")
            return
        }

        var updateData struct {
            Team1Score *int    `json:"team1_score"`
            Team2Score *int    `json:"team2_score"`
            Status     *string `json:"status"`
        }

        if err := ctx.ShouldBindJSON(&updateData); err != nil {
            utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request body")
            return
        }

        if updateData.Team1Score != nil {
            match.Team1Score = *updateData.Team1Score
        }
        if updateData.Team2Score != nil {
            match.Team2Score = *updateData.Team2Score
        }
        if updateData.Status != nil {
            match.Status = *updateData.Status
        }

        if err := db.Save(&match).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to update match")
            return
        }

        utils.RespondWithSuccess(ctx, http.StatusOK, "Match updated successfully", match)
    }
}
