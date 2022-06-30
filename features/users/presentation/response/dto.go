package response

import (
	"project2/features/users"
	"time"
)

//ambil data yang akan di tampilkan
type User struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

// untuk menampilkan data
func FromCore(data users.Core) User {
	return User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key, _ := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
