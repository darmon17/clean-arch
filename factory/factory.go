package factory

import (
	_userBusiness "project2/features/users/business"
	_userData "project2/features/users/data"
	_userPresentation "project2/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPrsenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)
	return Presenter{
		UserPrsenter: userPresentation,
	}
}
