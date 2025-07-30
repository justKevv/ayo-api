package database

import (
	"fmt"

	"gin/config"
	"gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	dbUser := config.LoadConfig().Database.User
	dbPassword := config.LoadConfig().Database.Password
	dbHost := config.LoadConfig().Database.Host
	dbPort := config.LoadConfig().Database.Port
	dbName := config.LoadConfig().Database.Name

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Team{},
		&models.Player{},
		&models.Match{},
		&models.Goal{},
	)
}
