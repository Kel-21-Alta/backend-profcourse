package users_courses

import (
	"context"
	"profcourse/business/courses"
	"profcourse/business/users"
	controller "profcourse/controllers"
	"time"
)

type UsersCoursesUsecase struct {
	UsersCoursesRepository Repository
	UsersUsecase           users.Usecase
	CourseUsecase          courses.Usecase
	ContextTime            time.Duration
}

func (u *UsersCoursesUsecase) GetUserCourseEndroll(ctx context.Context, domain *User) (User, error) {
	if domain.UserID == "" {
		return User{}, nil
	}

	result, err := u.UsersCoursesRepository.GetUserCourseEndroll(ctx, domain)
	if err != nil {
		return User{}, err
	}

	result.UserID = domain.UserID

	user, err := u.UsersUsecase.GetCurrentUser(ctx, users.Domain{ID: domain.UserID})

	if err != nil {
		return User{}, err
	}
	result.Name = user.Name

	var listCourse []Domain

	for _, courseid := range result.Courses {
		course, err := u.CourseUsecase.GetOneCourse(ctx, &courses.Domain{ID: courseid.CourseId})
		if err != nil {
			return User{}, err
		}
		listCourse = append(listCourse, Domain{
			ID:          courseid.ID,
			UserId:      courseid.UserId,
			CourseId:    courseid.CourseId,
			Progres:     courseid.Progres,
			LastVideoId: courseid.LastVideoId,
			LastModulId: courseid.LastModulId,
			Score:       courseid.Score,
			CourseTitle: course.Title,
			UrlImage:    course.ImgUrl,
			CreatedAt:   courseid.CreatedAt,
			UpdatedAt:   courseid.UpdatedAt,
		})
	}
	result.Courses = listCourse
	return result, nil
}

func (u *UsersCoursesUsecase) UpdateScoreCourse(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.CourseId == "" {
		return Domain{}, controller.EMPTY_COURSE
	}
	if domain.UserId == "" {
		return Domain{}, controller.ID_EMPTY
	}
	if domain.ID == "" {
		return Domain{}, controller.EMPTY_COURSE
	}
	result, err := u.UsersCoursesRepository.UpdateScoreCourse(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (u *UsersCoursesUsecase) GetOneUserCourse(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.CourseId == "" {
		return Domain{}, controller.EMPTY_COURSE
	}

	if domain.UserId == "" {
		return Domain{}, controller.ID_EMPTY
	}

	result, err := u.UsersCoursesRepository.GetOneUserCourse(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (u *UsersCoursesUsecase) UpdateProgressCourse(ctx context.Context, domain *Domain) (Domain, error) {

	if domain.CourseId == "" {
		return Domain{}, controller.EMPTY_COURSE
	}

	if domain.UserId == "" {
		return Domain{}, controller.ID_EMPTY
	}

	if domain.LastVideoId == "" {
		return Domain{}, controller.LAST_MATERI_EMPTY
	}

	result, err := u.UsersCoursesRepository.UpdateProgressCourse(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (u *UsersCoursesUsecase) UserRegisterCourse(ctx context.Context, domain *Domain) (*Domain, error) {
	// Validasi request empty
	if domain.UserId == "" {
		return &Domain{}, controller.EMPTY_USER
	}
	if domain.CourseId == "" {
		return &Domain{}, controller.EMPTY_COURSE
	}

	// Untuk melakukan cek apakah user udah mendaftar apa belum
	existedUserCourse, _ := u.UsersCoursesRepository.GetEndRollCourseUserById(ctx, domain)

	if *existedUserCourse != (Domain{}) {
		return &Domain{}, controller.ALREADY_REGISTERED_COURSE
	}

	userCourseDomain, err := u.UsersCoursesRepository.UserRegisterCourse(ctx, domain)

	if err != nil {
		return &Domain{}, err
	}

	return userCourseDomain, nil
}

func NewUsersCoursesUsecase(r Repository, user users.Usecase, courseUsecase courses.Usecase, timeout time.Duration) Usecase {
	return &UsersCoursesUsecase{UsersCoursesRepository: r, ContextTime: timeout, UsersUsecase: user, CourseUsecase: courseUsecase}
}
