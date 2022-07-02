package response

import (
	"project2/features/users"
	"time"
)

//ambil data yang akan di tampilkan
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
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
