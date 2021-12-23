package users

import (
	"context"
	"fmt"
	"profcourse/app/middlewares"
	"profcourse/business/smtpEmail"
	controller "profcourse/controllers"
	"profcourse/helpers/encrypt"
	"profcourse/helpers/randomString"
	"profcourse/helpers/validators"
	"time"
)

type userUsecase struct {
	ContextTimeout time.Duration
	UserRepository Repository
	SmtpRepository smtpEmail.Repository
	JWTConfig      middlewares.ConfigJwt
}

func (u userUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	var existedUser Domain

	if domain.Email == "" {
		return Domain{}, controller.EMPTY_EMAIL
	}
	if domain.Password == "" {
		return Domain{}, controller.PASSWORD_EMPTY
	}

	if !validators.CheckEmail(domain.Email) {
		return Domain{}, controller.INVALID_EMAIL
	}

	existedUser, err = u.UserRepository.GetUserByEmail(ctx, domain.Email)

	if err != nil {
		return Domain{}, err
	}

	// Mengecek email apakah benar ada usernya
	if existedUser == (Domain{}) {
		return Domain{}, controller.WRONG_EMAIL
	}

	// Mengecek apakan passwordnya benar
	if !encrypt.ValidateHash(domain.Password, existedUser.HashPassword) {
		return Domain{}, controller.WRONG_PASSWORD
	}

	existedUser.Token, err = u.JWTConfig.GenrateTokenJWT(domain.ID, domain.Role)

	if err != nil {
		fmt.Println("sini")
		return Domain{}, err
	}

	return existedUser, nil
}

func (u userUsecase) CreateUser(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	var existedUser Domain
	if domain.Name == "" {
		return Domain{}, controller.EMPTY_NAME
	}

	if domain.Email == "" {
		return Domain{}, controller.EMPTY_EMAIL
	}

	// Mengecek apakah email yang diberika valid
	if !validators.CheckEmail(domain.Email) {
		return Domain{}, controller.INVALID_EMAIL
	}

	// Mengecek apakan Email telah digunakan atau belum
	existedUser, err = u.UserRepository.GetUserByEmail(ctx, domain.Email)

	if existedUser != (Domain{}) {
		return Domain{}, controller.EMAIL_UNIQUE
	}

	// Melakukan hashing pada password
	domain.Password = randomString.RandomString(8)
	domain.HashPassword, err = encrypt.Hash(domain.Password)

	if err != nil {
		return Domain{}, err
	}

	domain.CreatedAt = time.Now()
	domain.UpdatedAt = time.Now()

	// Mengirim password dengan email
	to := domain.Email
	subject := "Pendaftaran akun di Profcouse"
	message := "<p>Dear " + domain.Name + "</p><br><p>Password anda pada akun profcourse adalah :" + domain.Password + " "

	err = u.SmtpRepository.SendEmail(ctx, to, subject, message)
	if err != nil {
		return Domain{}, err
	}

	// Mengirim domain to layer mysql repository user
	var resultDomain Domain
	resultDomain, err = u.UserRepository.CreateUser(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return resultDomain, nil
}

func NewUserUsecase(r Repository, timeout time.Duration, smtpRepo smtpEmail.Repository, configJwt middlewares.ConfigJwt) Usecase {
	return &userUsecase{
		ContextTimeout: timeout,
		UserRepository: r,
		SmtpRepository: smtpRepo,
		JWTConfig:      configJwt,
	}
}
