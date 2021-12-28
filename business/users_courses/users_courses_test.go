package users_courses_test

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"profcourse/business/users_courses"
	_mocksUsersCoursesRepository "profcourse/business/users_courses/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var usersCoursesRepository _mocksUsersCoursesRepository.Repository

var usersCoursesService users_courses.Usecase
var usersCoursesDomain users_courses.Domain

func setupUserRegisterCourse() {
	usersCoursesService = users_courses.NewUsersCoursesUsecase(&usersCoursesRepository)
	usersCoursesDomain = users_courses.Domain{
		ID:          uuid.NewV4().String(),
		UserId:      uuid.NewV4().String(),
		CourseId:    uuid.NewV4().String(),
		Progres:     0,
		LastVideoId: "",
		LastModulId: "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func TestUsersCoursesUsecase_UserRegisterCourse(t *testing.T) {
	t.Run("Test Case 1 | Handle error empty user", func(t *testing.T) {
		setupUserRegisterCourse()
		
		_, err := usersCoursesService.UserRegisterCourse(context.Background(), &users_courses.Domain{UserId: "", CourseId: uuid.NewV4().String()})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_USER, err)
	})
	t.Run("Test Case 2 | Handle error empty course ID", func(t *testing.T) {
		setupUserRegisterCourse()

		_, err := usersCoursesService.UserRegisterCourse(context.Background(), &users_courses.Domain{UserId: uuid.NewV4().String()})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
	t.Run("Test Case 3 | Handle error db Get Endroll course User By id", func(t *testing.T) {
		setupUserRegisterCourse()

		usersCoursesRepository.On("GetEndRollCourseUserById", mock.Anything, mock.Anything).Return(&users_courses.Domain{}, errors.New("Error DB")).Once()

		_, err := usersCoursesService.UserRegisterCourse(context.Background(), &users_courses.Domain{UserId: uuid.NewV4().String(), CourseId: uuid.NewV4().String()})

		assert.NotNil(t, err)
		assert.NotEqual(t, controller.EMPTY_COURSE, err)
		assert.NotEqual(t, controller.EMPTY_USER, err)
		assert.Equal(t, errors.New("Error DB"), err)
	})
	t.Run("Test Case 4 | Handle error user sudah mendaftar kursus", func(t *testing.T) {
		setupUserRegisterCourse()

		usersCoursesRepository.On("GetEndRollCourseUserById", mock.Anything, mock.Anything).Return(&usersCoursesDomain, nil).Once()

		_, err := usersCoursesService.UserRegisterCourse(context.Background(), &users_courses.Domain{UserId: uuid.NewV4().String(), CourseId: uuid.NewV4().String()})

		assert.NotNil(t, err)
		assert.NotEqual(t, controller.EMPTY_COURSE, err)
		assert.NotEqual(t, controller.EMPTY_USER, err)
		assert.Equal(t, controller.ALREADY_REGISTERED_COURSE, err)
	})
	t.Run("Test Case 5 | Handle error db create user_course/ daftar course user", func(t *testing.T) {
		setupUserRegisterCourse()

		usersCoursesRepository.On("GetEndRollCourseUserById", mock.Anything, mock.Anything).Return(&users_courses.Domain{}, nil).Once()

		usersCoursesRepository.On("UserRegisterCourse", mock.Anything, mock.Anything).Return(&users_courses.Domain{}, errors.New("Error DB")).Once()

		_, err := usersCoursesService.UserRegisterCourse(context.Background(), &users_courses.Domain{UserId: uuid.NewV4().String(), CourseId: uuid.NewV4().String()})

		assert.NotNil(t, err)
		assert.NotEqual(t, controller.EMPTY_COURSE, err)
		assert.NotEqual(t, controller.EMPTY_USER, err)
		assert.NotEqual(t, controller.ALREADY_REGISTERED_COURSE, err)
	})
	t.Run("Test Case 6 | Success register course user", func(t *testing.T) {
		setupUserRegisterCourse()

		usersCoursesRepository.On("GetEndRollCourseUserById", mock.Anything, mock.Anything).Return(&users_courses.Domain{}, nil).Once()

		usersCoursesRepository.On("UserRegisterCourse", mock.Anything, mock.Anything).Return(&usersCoursesDomain, nil).Once()

		_, err := usersCoursesService.UserRegisterCourse(context.Background(), &users_courses.Domain{UserId: uuid.NewV4().String(), CourseId: uuid.NewV4().String()})

		assert.Nil(t, err)
		assert.NotEqual(t, controller.EMPTY_COURSE, err)
		assert.NotEqual(t, controller.EMPTY_USER, err)
		assert.NotEqual(t, controller.ALREADY_REGISTERED_COURSE, err)
	})
}
