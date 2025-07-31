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
	api := router.Group("/api")
	{
		api.GET("/players", controllers.GetAllPlayers(db))
		api.GET("/players/:id", controllers.ShowPlayer(db))
		api.POST("/players", controllers.CreatePlayer(db))
		api.PUT("/players/:id", controllers.UpdatePlayer(db))
		api.DELETE("/players/:id", controllers.DeletePlayer(db))

		api.GET("/teams", controllers.GetAllTeams(db))
		api.GET("/teams/:id", controllers.ShowTeam(db))
		api.POST("/teams", controllers.CreateTeam(db))
		api.PUT("/teams/:id", controllers.UpdateTeam(db))
		api.DELETE("/teams/:id", controllers.DeleteTeam(db))

        api.GET("/matches", controllers.GetAllMatches(db))
        api.GET("/matches/:id", controllers.ShowMatch(db))
        api.POST("/matches", controllers.CreateMatch(db))
        api.PUT("/matches/:id/score", controllers.UpdateMatchScore(db))
        api.GET("/matches/:id/report", controllers.GetMatchReport(db))

        api.POST("/goals", controllers.CreateGoal(db))
        api.GET("/matches/:id/goals", controllers.GetMatchGoals(db))
	}

	return router
}
