package users_test

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"profcourse/business/users"
	_mockUserMysqlRepo "profcourse/business/users/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var userMysqlRepository _mockUserMysqlRepo.Repository

var userService users.Usecase
var userDomain users.Domain

func setUpCreateUser() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1)
	userDomain = users.Domain{
		ID:           uuid.NewV4().String(),
		Name:         "test",
		Email:        "test@gmail.com",
		Password:     "dafsdfiejq",
		HashPassword: "fadfijiweojq",
		NoHp:         "01293143",
		Birth:        time.Now(),
		BirthPlace:   "Medan",
		Bio:          "agheuirhe",
		ImgProfile:   "fasdfihruie",
		Role:         2,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

func TestUserUsecase_CreateUser(t *testing.T) {
	t.Run("Test Case 1 | Email Empty", func(t *testing.T) {
		setUpCreateUser()

		_, err := userService.CreateUser(context.Background(), users.Domain{
			Name:  "test",
			Email: "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_EMAIL, err)
	})
	t.Run("Test Case 2 | Name Empty", func(t *testing.T) {
		setUpCreateUser()

		_, err := userService.CreateUser(context.Background(), users.Domain{
			Name:  "",
			Email: "test@gmail.com",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_NAME, err)
	})
	t.Run("Test Case 3 | Invalid Email", func(t *testing.T) {
		setUpCreateUser()
		userMysqlRepository.On("CreateUser", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		_, err := userService.CreateUser(context.Background(), users.Domain{
			Name:  "test",
			Email: "test@sadas.com",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.INVALID_EMAIL, err)
	})
	t.Run("Test Case 4 | Success Create User", func(t *testing.T) {
		setUpCreateUser()
		userMysqlRepository.On("CreateUser", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(users.Domain{}, errors.New("Email sudah digunakan")).Once()

		user, err := userService.CreateUser(context.Background(), users.Domain{
			Name:  "test",
			Email: "test@gmail.com",
		})
		assert.Nil(t, err)
		assert.Equal(t, "test", user.Name)
		assert.Equal(t, "test@gmail.com", user.Email)
	})
	t.Run("Test Case 4 | Cek Email Unique", func(t *testing.T) {
		setUpCreateUser()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		_, err := userService.CreateUser(context.Background(), users.Domain{
			Name:  "test",
			Email: "test@gmail.com",
		})
		assert.NotNil(t, err)
	})
}
