package config

import (
	"log"

	"github.com/mfuadfakhruzaki/Jadwalin/models"
)

func MigrateDB() {
	// Automatically migrate models to the database
	err := DB.AutoMigrate(
		&models.User{},        // Migrating User model
		&models.UserCourse{},
		&models.FCMToken{},	// Migrating FCMToken model
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		log.Println("Database migrated successfully")
	}
}
