package users_courses

import (
	"context"
	controller "profcourse/controllers"
)

type UsersCoursesUsecase struct {
	UsersCoursesRepository Repository
}

func (u UsersCoursesUsecase) UserRegisterCourse(ctx context.Context, domain *Domain) (*Domain, error) {
	// Validasi request empty
	if domain.UserId == "" {
		return &Domain{}, controller.EMPTY_USER
	}
	if domain.CourseId == "" {
		return &Domain{}, controller.EMPTY_COURSE
	}

	// Untuk melakukan cek apakah user udah mendaftar apa belum
	existedUserCourse, err := u.UsersCoursesRepository.GetEndRollCourseUserById(ctx, domain)

	if err != nil {
		return &Domain{}, err
	}

	if *existedUserCourse != (Domain{}) {
		return &Domain{}, controller.ALREADY_REGISTERED_COURSE
	}

	userCourseDomain, err := u.UsersCoursesRepository.UserRegisterCourse(ctx, domain)

	if err != nil {
		return &Domain{}, err
	}

	return userCourseDomain, nil
}

func NewUsersCoursesUsecase(r Repository) Usecase {
	return &UsersCoursesUsecase{UsersCoursesRepository: r}
}
