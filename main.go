package main

import (
	"gin/config"
	"gin/routes"
)

func main() {
	router := routes.Setup()
	router.Run(":" + config.LoadConfig().Server.Port)
}
