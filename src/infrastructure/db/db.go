package db

import (
	"app/src/domain/user"
	"app/src/infrastructure/config"
)

func Migrate() {
	// Drop existing table if exists
	config.DB.Migrator().DropTable(&user.User{})

	// Run the new migration
	config.DB.AutoMigrate(&user.User{})
}
