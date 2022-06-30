package migration

import (
	_mUsers "project2/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{}) //otomatis buat table di database
}
