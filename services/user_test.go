package userservice_test

import (
	"testing"

	"localhost/mocks"
	userservice "localhost/services"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	mockUserRepo.EXPECT().CreateUser().Return("[REPO] POST /user")

	userService := userservice.NewUserService(mockUserRepo)
	result := userService.CreateUser()

	assert.Equal(t, "[REPO] POST /user", result)
}


