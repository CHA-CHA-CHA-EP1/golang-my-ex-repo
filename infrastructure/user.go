package intrastructure;

import (
	userrepository "localhost/domain/repositories"
)

type userRepository struct{}

func NewUserRepository() userrepository.UserRepository {
	return &userRepository{}
}

func (u *userRepository) GetUser() string {
	return "GET /user"
}

func (u *userRepository) GetUserByID() string {
	return "GET /user/:id"
}

func (u *userRepository) CreateUser() string {
	return "[REPO] POST /user"
}

func (u *userRepository) UpdateUser() string {
	return "PUT /user"
}

func (u *userRepository) DeleteUser() string {
	return "DELETE /user"
}
