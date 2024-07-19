package db

import (
	"app/config"
	"app/domain/user"
)

func Migrate() {
	config.DB.AutoMigrate(&user.User{})
}
