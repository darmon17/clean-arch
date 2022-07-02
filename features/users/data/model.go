package data

import (
	"project2/features/users"

	"gorm.io/gorm"
)

//database model (isi dari table yang di auto migrate di file migration)
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

// dto
func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
func toCoreList(data []User) []users.Core {
	result := []users.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
func FromCore(core users.Core) User {
	return User{
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
	}
}
