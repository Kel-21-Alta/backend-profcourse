package users_spesializations

import (
	"golang.org/x/net/context"
	"profcourse/business/spesializations"
	"profcourse/business/users_courses"
	controller "profcourse/controllers"
	"time"
)

type UsersSpesializationsUsecase struct {
	UserSpesializationRepository Repository
	SpesializationUsecase        spesializations.Usecase
	UserCourseUsecase            users_courses.Usecase
	ConTextTimeout               time.Duration
}

func (u UsersSpesializationsUsecase) RegisterSpesialization(ctx context.Context, domain *Domain) (Domain, error) {

	if domain.SpesializationID == "" {
		return Domain{}, controller.SPESIALIZATION_ID_EMPTY
	}

	if domain.UserID == "" {
		return Domain{}, controller.ID_EMPTY
	}

	endroll, _ := u.UserSpesializationRepository.GetEndRollSpesializationById(ctx, domain)

	if endroll != (Domain{}) {
		return Domain{}, controller.ALREADY_REGISTERED_SPESIALIZATION
	}

	spesialization, err := u.SpesializationUsecase.GetOneSpesialization(ctx, &spesializations.Domain{ID: domain.SpesializationID})

	if err != nil {
		return Domain{}, err
	}

	result, err := u.UserSpesializationRepository.RegisterSpesialization(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	for _, course := range spesialization.Courses {

		_, err := u.UserCourseUsecase.UserRegisterCourse(ctx, &users_courses.Domain{
			UserId:   domain.UserID,
			CourseId: course.ID,
		})

		if err != nil && err != controller.ALREADY_REGISTERED_COURSE {
			return Domain{}, err
		}

	}

	return result, nil
}

func NewUsersSpesializationsUsecase(repo Repository, spesialization spesializations.Usecase, userCourse users_courses.Usecase, timeout time.Duration) Usecase {
	return &UsersSpesializationsUsecase{
		UserSpesializationRepository: repo,
		ConTextTimeout:               timeout,
		SpesializationUsecase:        spesialization,
		UserCourseUsecase:            userCourse,
	}
}
