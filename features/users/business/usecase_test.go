package business

import (
	"fmt"
	"project2/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

//tampung mock data (data tiruan)
type mockUserDataSuccess struct{}

//mock data success
//mock SelectData yang di ambil dari data/mysql dan return SelectData yang di ambil dari entity Data
func (mock mockUserDataSuccess) SelectData(param string) (data []users.Core, err error) {
	return []users.Core{
		{ID: 1, Name: "heri", Email: "heri@mail.com", Password: "qwerty"},
	}, nil
}

//mock data failed
type mockUserDataFailed struct{}

func (mock mockUserDataFailed) SelectData(param string) (data []users.Core, err error) {
	return nil, fmt.Errorf("Faild to get all data")
}

func TestGetAlldata(t *testing.T) {
	t.Run("Test Get All Data succes", func(t *testing.T) {
		userBusines := NewUserBusiness(mockUserDataSuccess{})
		result, err := userBusines.GetAllData(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, "alta", result[0].Name)
	})
	t.Run("Test Get All Data Failed", func(t *testing.T) {
		userBusines := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusines.GetAllData(0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
