package presentation

import (
	"net/http"
	"project2/features/users"

	_requestUser "project2/features/users/presentation/request"
	_responseUser "project2/features/users/presentation/response"
	_helper "project2/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.userBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    _responseUser.FromCoreList(result),
	})
}

// sama mas yusuf
// func (h *UserHandler) Register(c echo.Context) error {
// 	//name, email, password
// 	dataUser := request.User{} //manggil request dto (struct)
// 	c.Bind(&dataUser)
// 	data := request.ToCore(dataUser)
// 	result, err := h.userBusiness.CreateData(data)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("Failed to register"))
// 	}
// 	if result == 1 {
// 		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("Success to register"))
// 	} else {
// 		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("Failed to insert"))
// 	}
// }

func (h *UserHandler) PostUser(c echo.Context) error {
	userReq := _requestUser.User{}
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("Failed to bind data, check your input"))
	}

	dataUser := _requestUser.ToCore(userReq)
	row, errCreate := h.userBusiness.CreateData(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filed in correctly"))
	}
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("Your email is alredy registered"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("Success"))
}

func (h *UserHandler) LoginAuth(c echo.Context) error {
	authData := users.AuthRequestData{}
	c.Bind(&authData)
	token, name, e := h.userBusiness.LoginUser(authData)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("Email or password incorrect"))
	}
	data := map[string]interface{}{
		"token": token,
		"name":  name,
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("Login Success", data))
}
