package users_test

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"profcourse/app/middlewares"
	_mockSmtpEmailRepo "profcourse/business/smtpEmail/mocks"
	"profcourse/business/users"
	_mockUserMysqlRepo "profcourse/business/users/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var userMysqlRepository _mockUserMysqlRepo.Repository
var smtpEmailRepository _mockSmtpEmailRepo.Repository
var configJwt middlewares.ConfigJwt

var userService users.Usecase
var userDomain users.Domain

func setUpCreateUser() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
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
		smtpEmailRepository.On("SendEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

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

func setUpLogin() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	userDomain = users.Domain{
		ID:           uuid.NewV4().String(),
		Name:         "test",
		Email:        "test1@gmail.com",
		Password:     "kQPPSkyR",
		HashPassword: "$2a$04$nHHmj1KfuzixIZ8nf9PFH.szVVWeCDsBG6bYYqbMGKhdAzGwzh35K",
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

func TestUserUsecase_Login(t *testing.T) {
	t.Run("Test Case 1 | cek email empty", func(t *testing.T) {
		setUpLogin()

		_, err := userService.Login(context.Background(), users.Domain{Email: "", Password: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_EMAIL, err)
	})
	t.Run("Test Case 2 | cek password empty", func(t *testing.T) {
		setUpLogin()
		_, err := userService.Login(context.Background(), users.Domain{
			Email:    "test1@gmail.com",
			Password: "",
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.PASSWORD_EMPTY, err)
	})
	t.Run("Test Case 3 | Cek Email Invalid", func(t *testing.T) {
		setUpLogin()
		_, err := userService.Login(context.Background(), users.Domain{
			Email:    "ahfkdhfjdkf",
			Password: "dsaddasd",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.INVALID_EMAIL, err)
	})
	t.Run("Test Case 4 | Cek Password Salah", func(t *testing.T) {
		setUpLogin()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		_, err := userService.Login(context.Background(), users.Domain{
			Email:    "test1@gmail.com",
			Password: "dafsdfiejq",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.WRONG_PASSWORD, err)
	})
	t.Run("Test Case 5 | Cek Success Login", func(t *testing.T) {
		setUpLogin()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Login(context.Background(), users.Domain{
			Email:    "test1@gmail.com",
			Password: "kQPPSkyR",
		})
		assert.Nil(t, err)
		assert.Equal(t, "test1@gmail.com", user.Email)
	})
	t.Run("Test Case 5 | Cek Email Salah", func(t *testing.T) {
		setUpLogin()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("bla bla")).Once()
		_, err := userService.Login(context.Background(), users.Domain{
			Email:    "dsadas@gmail.com",
			Password: "kQPPSkyR",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.WRONG_EMAIL, err)
	})
}

func setForgetPassword() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	userDomain = users.Domain{
		ID:           uuid.NewV4().String(),
		Name:         "test",
		Email:        "test1@gmail.com",
		Password:     "kQPPSkyR",
		HashPassword: "$2a$04$nHHmj1KfuzixIZ8nf9PFH.szVVWeCDsBG6bYYqbMGKhdAzGwzh35K",
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

func TestUserUsecase_ForgetPassword(t *testing.T) {
	t.Run("Test Case 1 | Email Empty", func(t *testing.T) {
		setForgetPassword()

		_, err := userService.ForgetPassword(context.Background(), users.Domain{Email: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_EMAIL, err)
	})

	t.Run("Test Case 2 | Cek Email Valid", func(t *testing.T) {
		setForgetPassword()
		_, err := userService.ForgetPassword(context.Background(), users.Domain{Email: "sdajdhf"})
		assert.NotNil(t, err)
		assert.Equal(t, controller.INVALID_EMAIL, err)
	})

	t.Run("Test Case 3 | Cek Email Salah", func(t *testing.T) {
		setForgetPassword()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("bla bla")).Once()
		_, err := userService.ForgetPassword(context.Background(), users.Domain{
			Email: "tfadfd@gmail.com",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.WRONG_EMAIL, err)
	})

	t.Run("Test Case 4 | Cek Success Forget Password", func(t *testing.T) {
		setForgetPassword()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdatePassword", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		smtpEmailRepository.On("SendEmail", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()

		user, err := userService.ForgetPassword(context.Background(), users.Domain{Email: "test1@gmail.com"})

		assert.Nil(t, err)
		assert.Equal(t, "test", user.Name)
	})

	t.Run("Test Case 5 | Cek Handle Error DB", func(t *testing.T) {
		setForgetPassword()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, errors.New("Record Empty")).Once()

		_, err := userService.ForgetPassword(context.Background(), users.Domain{Email: "test1@gmail.com"})

		assert.NotNil(t, err)
	})
	t.Run("Test Case 6 | Cek Handle Error DB Update", func(t *testing.T) {
		setForgetPassword()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdatePassword", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(userDomain, errors.New("Err s")).Once()

		_, err := userService.ForgetPassword(context.Background(), users.Domain{Email: "test1@gmail.com"})

		assert.NotNil(t, err)
	})
	t.Run("Test Case 7 | Cek Handle Error Send Email", func(t *testing.T) {
		setForgetPassword()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdatePassword", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		smtpEmailRepository.On("SendEmail", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("Err")).Once()

		_, err := userService.ForgetPassword(context.Background(), users.Domain{Email: "test1@gmail.com"})

		assert.NotNil(t, err)
	})
}

func setUpCurrentUser() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	userDomain = users.Domain{
		ID:           uuid.NewV4().String(),
		Name:         "test",
		Email:        "test1@gmail.com",
		Password:     "kQPPSkyR",
		HashPassword: "$2a$04$nHHmj1KfuzixIZ8nf9PFH.szVVWeCDsBG6bYYqbMGKhdAzGwzh35K",
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
func TestUserUsecase_GetCurrentUser(t *testing.T) {
	t.Run("Test case 1 | ID tidak ada", func(t *testing.T) {
		setUpCurrentUser()

		_, err := userService.GetCurrentUser(context.Background(), users.Domain{ID: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("Test case 2 | Success", func(t *testing.T) {
		setUpCurrentUser()
		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil)
		result, err := userService.GetCurrentUser(context.Background(), users.Domain{ID: "dsadsd"})
		assert.Nil(t, err)
		assert.Equal(t, "test1@gmail.com", result.Email)
	})
}
