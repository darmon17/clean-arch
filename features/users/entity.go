package users

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Business interface {
	GetAllData(limit int, offset int) (data []Core, err error)
	LoginUser(authData AuthRequestData) (token, name string, err error)
	// CreateData(data Core) (row int, err error) sama mas yusuf
	CreateData(data Core) (row int, err error)
}
type Data interface {
	SelectData(param string) (data []Core, err error)
	LoginUserDB(authData AuthRequestData) (token, name string, err error)
	// CreateData(data Core) (row int, err error) sama mas yusuf
	InsertData(data Core) (row int, err error)
}
