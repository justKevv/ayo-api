package controllers

import (
	"gin/models"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTeams(db *gorm.DB) gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		var teams []models.Team

		if err := db.Find(&teams).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to fetch teams")
			return
		}

		utils.RespondWithSuccess(ctx, http.StatusOK, "Teams retrieved successfully", teams)
	}
}

func ShowTeam(db *gorm.DB) gin.HandlerFunc  {
    return func(ctx *gin.Context) {
        var team models.Team
        if err := db.Preload("Players").First(&team, ctx.Param("id")).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusNotFound, "Team not found")
            return
        }

        // Set team field to nil for each player to avoid circular reference
        for i := range team.Players {
            team.Players[i].Team = nil
        }

        utils.RespondWithSuccess(ctx, http.StatusOK, "Team retrieved successfully", team)
    }
}

func CreateTeam(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var team models.Team
		if err := ctx.ShouldBindJSON(&team); err != nil {
			utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}
		if err := db.Create(&team).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to create team")
			return
		}

		utils.RespondWithSuccess(ctx, http.StatusCreated, "Team created successfully", team)
	}
}

func UpdateTeam(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var team models.Team
		if err := db.First(&team, ctx.Param("id")).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusNotFound, "Team not found")
			return
		}
		var updateData map[string]interface{}
		if err := ctx.ShouldBindJSON(&updateData); err != nil {
			utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}
		if err := db.Model(&team).Updates(updateData).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to update team")
			return
		}
		if err := db.First(&team, team.ID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to fetch updated team")
            return
        }
		utils.RespondWithSuccess(ctx, http.StatusOK, "Team updated successfully", team)
	}
}

func DeleteTeam(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var team models.Team
		if err := db.First(&team, ctx.Param("id")).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusNotFound, "Team not found")
			return
		}
		if err := db.Delete(&team).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to delete team")
			return
		}
		utils.RespondWithSuccess(ctx, http.StatusOK, "Team deleted successfully", team)
	}
}
