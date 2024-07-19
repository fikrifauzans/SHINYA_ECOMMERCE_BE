package db

import (
	"app/src/domain/user"
	"app/src/infrastructure/config"
)

func Migrate() {
	config.DB.AutoMigrate(&user.User{})
}
