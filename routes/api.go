package routes

import (
	"gin/controllers"
	"gin/database"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	db := database.InitializeDB()
	database.AutoMigrate(db)
	router := gin.Default()

	router.GET("/api/players", controllers.GetAllPlayers(db))
	router.GET("/api/players/:id", controllers.ShowPlayer(db))
	router.POST("/api/players", controllers.CreatePlayer(db))
	router.PUT("/api/players/:id", controllers.UpdatePlayer(db))
	router.DELETE("/api/players/:id", controllers.DeletePlayer(db))

	router.GET("/api/teams", controllers.GetAllTeams(db))
	router.GET("/api/teams/:id", controllers.ShowTeam(db))
	router.POST("/api/teams", controllers.CreateTeam(db))
	router.PUT("/api/teams/:id", controllers.UpdateTeam(db))
	router.DELETE("/api/teams/:id", controllers.DeleteTeam(db))

	return router
}
