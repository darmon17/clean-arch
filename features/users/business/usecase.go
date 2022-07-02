package business

import (
	"errors"
	"project2/features/users"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (uc *userUsecase) GetAllData(limit, offset int) (resp []users.Core, err error) {
	var param string
	resp, err = uc.userData.SelectData(param)
	return resp, err
}

// sama mas yusuf
// func (uc *userUsecase) CreateData(data users.Core) (row int, err error) {
// 	row, err = uc.userData.CreateData(data)
// 	return row, err
// }

func (uc *userUsecase) CreateData(input users.Core) (row int, err error) {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return -1, errors.New("please make sure all fields are filled in correclt")
	}
	row, err = uc.userData.InsertData(input)
	return row, err
}

func (uc *userUsecase) LoginUser(authData users.AuthRequestData) (token, name string, err error) {
	token, name, err = uc.userData.LoginUserDB(authData)
	return token, name, err
}
