package users

import (
	"context"
	controller "profcourse/controllers"
	"profcourse/helpers/encrypt"
	"profcourse/helpers/randomString"
	"profcourse/helpers/validators"
	"time"
)

type userUsecase struct {
	contextTimeout time.Duration
	userRepository Repository
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
	existedUser, err = u.userRepository.GetUserByEmail(ctx, domain.Email)

	if existedUser != (Domain{}) {
		return Domain{}, controller.EMAIL_UNIQUE
	}

	domain.Password = randomString.RandomString(8)
	domain.HashPassword, err = encrypt.Hash(domain.Password)

	if err != nil {
		return Domain{}, err
	}

	domain.CreatedAt = time.Now()
	domain.UpdatedAt = time.Now()

	// Mengirim domain to layer mysql repository user
	var resultDomain Domain
	resultDomain, err = u.userRepository.CreateUser(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return resultDomain, nil
}

func NewUserUsecase(r Repository, timeout time.Duration) Usecase {
	return &userUsecase{
		contextTimeout: timeout,
		userRepository: r,
	}
}
