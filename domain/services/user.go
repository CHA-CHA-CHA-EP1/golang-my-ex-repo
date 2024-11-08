package userservice;

type UserService interface {
	GetUser() string
	GetUserByID() string	
	CreateUser() string
	UpdateUser() string
	DeleteUser() string
}

