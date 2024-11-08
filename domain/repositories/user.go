package userrepository;

type UserRepository interface {
	GetUser() string
	GetUserByID() string
	CreateUser() string
	UpdateUser() string
	DeleteUser() string
}

