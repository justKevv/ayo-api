package controllers

import (
	"gin/models"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPlayers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var players []models.Player

		if err := db.Preload("Team").Find(&players).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to fetch players")
			return
		}

		utils.RespondWithSuccess(ctx, http.StatusOK, "Players retrieved successfully", players)
	}
}

func ShowPlayer(db *gorm.DB) gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		var player models.Player
		if err := db.Preload("Team").First(&player, ctx.Param("id")).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusNotFound, "Player not found")
			return
		}
		utils.RespondWithSuccess(ctx, http.StatusOK, "Player retrieved successfully", player)
	}
}

func CreatePlayer(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var player models.Player
		if err := ctx.ShouldBindJSON(&player); err != nil {
			utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}
		var team models.Team
		if err := db.First(&team, player.TeamID).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusNotFound, "Team not found")
			return
		}

		if err := db.Create(&player).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to create player")
			return
		}

		if err := db.Preload("Team").First(&player, player.ID).Error; err != nil {
            utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve created player")
            return
        }

		utils.RespondWithSuccess(ctx, http.StatusCreated, "Player created successfully", player)
	}
}

func UpdatePlayer(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var player models.Player
		if err := db.First(&player, ctx.Param("id")).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusNotFound, "Player not found")
			return
		}
		if err := ctx.ShouldBindJSON(&player); err != nil {
			utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}
		if err := db.Save(&player).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to update player")
			return
		}
		utils.RespondWithSuccess(ctx, http.StatusOK, "Player updated successfully", player)
	}
}

func DeletePlayer(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var player models.Player
		if err := db.Delete(&player, ctx.Param("id")).Error; err != nil {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to delete player")
			return
		}

		utils.RespondWithSuccess(ctx, http.StatusOK, "Player deleted successfully", nil)
	}
}
