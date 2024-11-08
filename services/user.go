package userservice

import (
	userservice "localhost/domain/services"
	userrepository "localhost/domain/repositories"

)

type userService struct{
	userrepo userrepository.UserRepository
}

func NewUserService(usrepo userrepository.UserRepository) userservice.UserService {
	return &userService{
		userrepo: usrepo,
	}
}

func (u *userService) GetUser() string {
	return "GET /user"
}

func (u *userService) GetUserByID() string {
	return "GET /user/:id"
}

func (u *userService) CreateUser() string {
	userrepoResponse := u.userrepo.CreateUser()
	return userrepoResponse
}

func (u *userService) UpdateUser() string {
	return "PUT /user"
}

func (u *userService) DeleteUser() string {
	return "DELETE /user"
}
