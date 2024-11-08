package api

import (
	models "localhost/domain/models"
	userservice "localhost/domain/services"


	"github.com/labstack/echo/v4"
)

type UserAPI interface {
	GetUser(c echo.Context) error
	GetUserByID(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userAPI struct {
	us userservice.UserService
}

func NewUserAPI(us userservice.UserService) UserAPI {
	return &userAPI{
		us: us,
	}
}

func (ap *userAPI) GetUser(c echo.Context) error {
	return c.String(200, "GET /user")
}

func (ap *userAPI) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, "GET /user/"+id)
}

func (ap *userAPI) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.String(400, "Bad Request")
	}

	return c.String(200, "POST /user")
}

func (ap *userAPI) UpdateUser(c echo.Context) error {
	return c.String(200, "PUT /user")
}

func (ap *userAPI) DeleteUser(c echo.Context) error {
	return c.String(200, "DELETE /user")
}

