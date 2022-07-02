package data

import (
	"errors"
	"fmt"
	"project2/features/middlewares"
	"project2/features/users"

	_bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) SelectData(data string) (response []users.Core, err error) {
	var dataUsers []User
	result := repo.db.Find(&dataUsers)
	if result.Error != nil {
		return []users.Core{}, result.Error
	}
	return toCoreList(dataUsers), nil
}

//sama mas yusuf
// func (repo *mysqlUserRepository) CreateData(data users.Core) (row int, err error) {
// 	row = repo.db.Create(&data)
// }

func (repo *mysqlUserRepository) InsertData(input users.Core) (row int, err error) {
	user := FromCore(input)
	passwordHashhed, errorHash := _bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorHash != nil {
		fmt.Println("Eror has ", errorHash.Error())
	}
	user.Password = string(passwordHashhed)
	resultcreate := repo.db.Create(&user)
	if resultcreate.Error != nil {
		return 0, resultcreate.Error
	}
	if resultcreate.RowsAffected != 1 {
		return 0, errors.New("Failed to insert data, your email is already registered")
	}
	return int(resultcreate.RowsAffected), nil

}

func (repo *mysqlUserRepository) LoginUserDB(authData users.AuthRequestData) (token string, name string, err error) {
	userData := User{}
	result := repo.db.Where("email = ?", authData.Email).First(&userData)
	if result.Error != nil {
		return "", "", result.Error
	}

	if result.RowsAffected != 1 {
		return "", "", errors.New("Failed to Login")
	}

	errCrypt := _bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(authData.Password))
	if errCrypt != nil {
		return "", "", errors.New("Password incorrect")
	}
	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return "", "", errToken
	}
	return token, userData.Name, nil
}
