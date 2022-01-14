package users_test

import (
	"context"
	"errors"
	"profcourse/app/middlewares"
	_mockSmtpEmailRepo "profcourse/business/send_email/mocks"
	"profcourse/business/users"
	_mockUserMysqlRepo "profcourse/business/users/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userMysqlRepository _mockUserMysqlRepo.Repository
var smtpEmailRepository _mockSmtpEmailRepo.Repository
var configJwt middlewares.ConfigJwt

var userService users.Usecase
var userDomain users.Domain
var userSummary users.Summary
var adminDomain users.Domain

func setUpGetCountUser() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	userSummary = users.Summary{CountUser: 9}
}

func TestUserUsecase_GetCountUser(t *testing.T) {
	t.Run("Test case 1 | success ", func(t *testing.T) {
		setUpGetCountUser()

		userMysqlRepository.On("GetCountUser", mock.Anything).Return(&userSummary, nil).Once()
		summary, err := userService.GetCountUser(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, 9, summary.CountUser)
	})
	t.Run("test case 2 | handle err db", func(t *testing.T) {
		setUpGetCountUser()

		userMysqlRepository.On("GetCountUser", mock.Anything).Return(&users.Summary{}, errors.New("haha")).Once()
		_, err := userService.GetCountUser(context.Background())

		assert.NotNil(t, err)
	})
}

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
	t.Run("Test Case 2 | Title Empty", func(t *testing.T) {
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

func setUpAdminLogin() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	adminDomain = users.Domain{
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
		Role:         1,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
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
func TestUserUsecase_LoginAdmin(t *testing.T) {
	t.Run("Test Case 1 | cek email empty", func(t *testing.T) {
		setUpAdminLogin()

		_, err := userService.LoginAdmin(context.Background(), users.Domain{Email: "", Password: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_EMAIL, err)
	})
	t.Run("Test Case 2 | cek password empty", func(t *testing.T) {
		setUpAdminLogin()
		_, err := userService.LoginAdmin(context.Background(), users.Domain{
			Email:    "test1@gmail.com",
			Password: "",
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.PASSWORD_EMPTY, err)
	})
	t.Run("Test Case 3 | Cek Email Invalid", func(t *testing.T) {
		setUpAdminLogin()
		_, err := userService.LoginAdmin(context.Background(), users.Domain{
			Email:    "ahfkdhfjdkf",
			Password: "dsaddasd",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.INVALID_EMAIL, err)
	})
	t.Run("Test Case 4 | Cek Password Salah", func(t *testing.T) {
		setUpAdminLogin()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(adminDomain, nil).Once()

		_, err := userService.LoginAdmin(context.Background(), users.Domain{
			Email:    "test1@gmail.com",
			Password: "dafsdfiejq",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.WRONG_PASSWORD, err)
	})
	t.Run("Test Case 5 | Cek Success Login", func(t *testing.T) {
		setUpAdminLogin()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(adminDomain, nil).Once()

		user, err := userService.LoginAdmin(context.Background(), users.Domain{
			Email:    "test1@gmail.com",
			Password: "kQPPSkyR",
		})
		assert.Nil(t, err)
		assert.Equal(t, "test1@gmail.com", user.Email)
	})
	t.Run("Test Case 5 | Cek Email Salah", func(t *testing.T) {
		setUpAdminLogin()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("bla bla")).Once()
		_, err := userService.LoginAdmin(context.Background(), users.Domain{
			Email:    "dsadas@gmail.com",
			Password: "kQPPSkyR",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.WRONG_EMAIL, err)
	})
	t.Run("Test case 6 | handle error role no admin", func(t *testing.T) {
		setUpAdminLogin()
		userMysqlRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		_, err := userService.LoginAdmin(context.Background(), users.Domain{
			Email:    "test1@gmail.com",
			Password: "kQPPSkyR",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.FORBIDDIN_USER, err)
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
	t.Run("Test case 3 | Handle mysql error", func(t *testing.T) {
		setUpCurrentUser()
		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("bla bla")).Once()
		_, err := userService.GetCurrentUser(context.Background(), users.Domain{ID: "dsadsd"})
		assert.NotNil(t, err)
	})
	t.Run("Test case 2 | Success", func(t *testing.T) {
		setUpCurrentUser()
		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		result, err := userService.GetCurrentUser(context.Background(), users.Domain{ID: "dsadsd"})
		assert.Nil(t, err)
		assert.Equal(t, "test1@gmail.com", result.Email)
	})

}

func setUpChangePassword() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	userDomain = users.Domain{
		ID:           "756f702e-69ae-45e2-8ab2-870c11f7ba51",
		Name:         "test",
		Email:        "test1@gmail.com",
		Password:     "kQPPSkyR",
		HashPassword: "$2a$04$nHHmj1KfuzixIZ8nf9PFH.szVVWeCDsBG6bYYqbMGKhdAzGwzh35K",
		PasswordNew:  "test1",
	}
}

func TestUserUsecase_ChangePassword(t *testing.T) {
	t.Run("Test Case 1 | Handle ID empty", func(t *testing.T) {
		setUpChangePassword()
		_, err := userService.ChangePassword(context.Background(), users.Domain{ID: "", Password: "kQPPSkyR", PasswordNew: "test1"})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})

	t.Run("Test Case 2 | Handle error salah password", func(t *testing.T) {
		setUpChangePassword()
		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		_, err := userService.ChangePassword(context.Background(), users.Domain{ID: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Password: "kQPPSkyRs", PasswordNew: "test1"})

		assert.NotNil(t, err)
		assert.Equal(t, controller.WRONG_PASSWORD, err)
	})

	t.Run("Test Case 3 | Handle error salah password", func(t *testing.T) {
		setUpChangePassword()
		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, errors.New("DB error")).Once()
		_, err := userService.ChangePassword(context.Background(), users.Domain{ID: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Password: "kQPPSkyRs", PasswordNew: "test1"})

		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Handle error db Get User By ID tidak ada record", func(t *testing.T) {
		setUpChangePassword()
		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("Record Not Found")).Once()
		user, err := userService.ChangePassword(context.Background(), users.Domain{ID: "adfdafawe", Password: "kQPPSkyR", PasswordNew: "test1"})
		assert.NotNil(t, err)
		assert.Equal(t, users.Domain{}, user)
	})

	t.Run("Test Case 5 | Handle error db upload password", func(t *testing.T) {
		setUpChangePassword()
		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdatePassword", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("Db Error")).Once()

		_, err := userService.ChangePassword(context.Background(), users.Domain{ID: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Password: "kQPPSkyR", PasswordNew: "test1"})

		assert.NotNil(t, err)
	})

	t.Run("Test Case 6 | Handle Db error update password", func(t *testing.T) {
		setUpChangePassword()

		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdatePassword", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("Db Error")).Once()

		_, err := userService.ChangePassword(context.Background(), users.Domain{ID: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Password: "kQPPSkyR", PasswordNew: "test1"})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 7 | Handle Success", func(t *testing.T) {
		setUpChangePassword()

		userMysqlRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userMysqlRepository.On("UpdatePassword", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		user, err := userService.ChangePassword(context.Background(), users.Domain{ID: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Password: "kQPPSkyR", PasswordNew: "test1"})
		assert.Nil(t, err)
		assert.Equal(t, userDomain.Email, user.Email)
	})
}
func setUpUpdateUser() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	userDomain = users.Domain{
		ID:         uuid.NewV4().String(),
		Name:       "Adrian",
		NoHp:       "01293143",
		Birth:      time.Now(),
		BirthPlace: "Medan",
		Bio:        "agheuirhe",
		Role:       2,
		RoleText:   "user",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	t.Run("Test Case 1 | Name Empty ", func(t *testing.T) {
		setUpUpdateUser()
		userMysqlRepository.On("UpdateUser",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(userDomain, nil).Once()
		_, err := userService.UpdateUser(context.Background(), users.Domain{
			ID:     "1",
			IdUser: "1",
			Role:   1,
			Name:   "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_NAME, err)
	})
	t.Run("Test Case 2 | ID Requester Empty ", func(t *testing.T) {
		setUpUpdateUser()
		userMysqlRepository.On("UpdateUser",
			mock.Anything,
			mock.Anything,
		).Return(userDomain, nil).Once()
		_, err := userService.UpdateUser(context.Background(), users.Domain{
			ID:     "",
			IdUser: "1",
			Role:   1,
			Name:   "Adrian",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("Test Case 3 | ID Requested Empty ", func(t *testing.T) {
		setUpUpdateUser()
		userMysqlRepository.On("UpdateUser",
			mock.Anything,
			mock.Anything,
		).Return(userDomain, nil).Once()
		_, err := userService.UpdateUser(context.Background(), users.Domain{
			ID:     "1",
			IdUser: "",
			Role:   1,
			Name:   "Adrian",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("Test Case 4 | Not Admin Update other user ", func(t *testing.T) {
		setUpUpdateUser()
		userMysqlRepository.On("UpdateUser",
			mock.Anything,
			mock.Anything,
		).Return(userDomain, nil).Once()
		_, err := userService.UpdateUser(context.Background(), users.Domain{
			ID:     "1",
			IdUser: "2",
			Role:   2,
			Name:   "Adrian",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.FORBIDDIN_USER, err)
	})
	t.Run("Test Case 5 | Not Admin Update Self Success ", func(t *testing.T) {
		setUpUpdateUser()
		userMysqlRepository.On("UpdateUser",
			mock.Anything,
			mock.Anything,
		).Return(userDomain, nil).Once()
		_, err := userService.UpdateUser(context.Background(), users.Domain{
			ID:     "1",
			IdUser: "1",
			Role:   2,
			Name:   "Adrian",
		})
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("Test Case 6 | Admin Update Success", func(t *testing.T) {
		setUpUpdateUser()
		userMysqlRepository.On("UpdateUser",
			mock.Anything,
			mock.Anything,
		).Return(userDomain, nil).Once()
		_, err := userService.UpdateUser(context.Background(), users.Domain{
			ID:     "2",
			IdUser: "1",
			Role:   1,
			Name:   "Adrian",
		})
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
}
func setUpDelete() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	userDomain = users.Domain{
		ID:           "756f702e-69ae-45e2-8ab2-870c11f7ba51",
		Name:         "test",
		Email:        "test1@gmail.com",
		Password:     "kQPPSkyR",
		HashPassword: "$2a$04$nHHmj1KfuzixIZ8nf9PFH.szVVWeCDsBG6bYYqbMGKhdAzGwzh35K",
		PasswordNew:  "test1",
		Message:      "User dengan id: 68d3127a-9db2-4798-bffa-0a047f25c355 telah dihapus",
	}
}

func TestUserUsecase_DeleteUser(t *testing.T) {
	t.Run("Test case 1 | yang mengirim request adalah bukan admin", func(t *testing.T) {
		setUpDelete()
		_, err := userService.DeleteUser(context.Background(), users.Domain{IdUser: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Role: 2})
		assert.NotNil(t, err)
		assert.Equal(t, controller.FORBIDDIN_USER, err)
	})
	t.Run("Test case 2 | success delete", func(t *testing.T) {
		setUpDelete()
		userMysqlRepository.On("DeleteUser", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		user, err := userService.DeleteUser(context.Background(), users.Domain{IdUser: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Role: 1})
		assert.Nil(t, err)
		assert.Equal(t, userDomain.Message, user.Message)
	})
	t.Run("Test case 3 | Handle error user id kosong", func(t *testing.T) {
		setUpDelete()
		_, err := userService.DeleteUser(context.Background(), users.Domain{IdUser: "", Role: 1})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("Test case 4 | Handle error from db", func(t *testing.T) {
		setUpDelete()
		userMysqlRepository.On("DeleteUser", mock.Anything, mock.Anything).Return(users.Domain{}, errors.New("Error DB")).Once()
		_, err := userService.DeleteUser(context.Background(), users.Domain{IdUser: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Role: 1})
		assert.NotNil(t, err)
	})
}

func setUpUpdateCurrentUset() {
	userService = users.NewUserUsecase(&userMysqlRepository, time.Hour*1, &smtpEmailRepository, configJwt)
	userDomain = users.Domain{
		ID:           "756f702e-69ae-45e2-8ab2-870c11f7ba51",
		Name:         "test",
		Email:        "test1@gmail.com",
		Password:     "kQPPSkyR",
		HashPassword: "$2a$04$nHHmj1KfuzixIZ8nf9PFH.szVVWeCDsBG6bYYqbMGKhdAzGwzh35K",
		PasswordNew:  "test1",
	}
}

func TestUserUsecase_UpdateDataCurrentUser(t *testing.T) {
	t.Run("Test case 1 | user barhasil mengubah data dirinya sendiri", func(t *testing.T) {
		setUpUpdateCurrentUset()
		userMysqlRepository.On("UpdateDataCurrentUser", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		result, err := userService.UpdateDataCurrentUser(context.Background(), &userDomain)

		assert.Nil(t, err)
		assert.Equal(t, userDomain.ID, result.ID)
	})
	t.Run("Test case 2 | handle error ketika user id tidak ada", func(t *testing.T) {
		setUpUpdateCurrentUset()
		_, err := userService.UpdateDataCurrentUser(context.Background(), &users.Domain{ID: ""})

		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("Test case 3 | handle error ketika name kosong", func(t *testing.T) {
		setUpUpdateCurrentUset()
		_, err := userService.UpdateDataCurrentUser(context.Background(), &users.Domain{ID: "756f702e-69ae-45e2-8ab2-870c11f7ba51", Name: ""})

		assert.Equal(t, controller.EMPTY_NAME, err)
	})
}
