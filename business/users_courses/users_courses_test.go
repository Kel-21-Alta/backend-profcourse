package users_courses_test

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"profcourse/business/courses"
	_mocksCourseUsecase "profcourse/business/courses/mocks"
	"profcourse/business/users"
	_mockUserUsecase "profcourse/business/users/mocks"
	"profcourse/business/users_courses"
	_mocksUsersCoursesRepository "profcourse/business/users_courses/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var usersCoursesRepository _mocksUsersCoursesRepository.Repository
var userUsecase _mockUserUsecase.Usecase
var courseUsecase _mocksCourseUsecase.Usecase

var usersCoursesService users_courses.Usecase
var usersCoursesDomain users_courses.Domain
var usersCourseUser users_courses.User
var userDomain users.Domain
var courseDomain courses.Domain

func setupUserRegisterCourse() {
	usersCoursesService = users_courses.NewUsersCoursesUsecase(&usersCoursesRepository, &userUsecase, &courseUsecase, time.Hour*1)
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

func setupGetUserCourseEndroll() {
	usersCoursesService = users_courses.NewUsersCoursesUsecase(&usersCoursesRepository, &userUsecase, &courseUsecase, time.Hour*1)
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
	usersCourseUser = users_courses.User{
		UserID:      "123",
		Name:        "HT",
		CountCourse: 0,
		Courses:     []users_courses.Domain{usersCoursesDomain},
	}
	userDomain = users.Domain{
		ID:         "123",
		Name:       "qwe",
		ImgProfile: "1232",
	}
	courseDomain = courses.Domain{
		ID:     "123",
		Title:  "qwe",
		ImgUrl: "asd",
	}
}

func TestUsersCoursesUsecase_GetUserCourseEndroll(t *testing.T) {
	t.Run("Test 1 | success get all course register", func(t *testing.T) {
		setupGetUserCourseEndroll()
		usersCoursesRepository.On("GetUserCourseEndroll", mock.Anything, mock.Anything).Return(usersCourseUser, nil).Once()
		userUsecase.On("GetCurrentUser", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()

		_, err := usersCoursesService.GetUserCourseEndroll(context.Background(), &users_courses.User{UserID: "123"})

		assert.Nil(t, err)
	})
	t.Run("Test 2 | user id empty", func(t *testing.T) {
		setupGetUserCourseEndroll()
		_, err := usersCoursesService.GetUserCourseEndroll(context.Background(), &users_courses.User{UserID: ""})

		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("Test 3 | db error get user course endroll", func(t *testing.T) {
		setupGetUserCourseEndroll()
		usersCoursesRepository.On("GetUserCourseEndroll", mock.Anything, mock.Anything).Return(usersCourseUser, errors.New("db err")).Once()
		_, err := usersCoursesService.GetUserCourseEndroll(context.Background(), &users_courses.User{UserID: "123"})

		assert.NotNil(t, err)
	})
	t.Run("Test 4 | error user usecase", func(t *testing.T) {
		setupGetUserCourseEndroll()
		usersCoursesRepository.On("GetUserCourseEndroll", mock.Anything, mock.Anything).Return(usersCourseUser, nil).Once()
		userUsecase.On("GetCurrentUser", mock.Anything, mock.Anything).Return(userDomain, errors.New("use case err")).Once()

		_, err := usersCoursesService.GetUserCourseEndroll(context.Background(), &users_courses.User{UserID: "123"})

		assert.NotNil(t, err)
	})
	t.Run("Test 5 | error course usecase", func(t *testing.T) {
		setupGetUserCourseEndroll()
		usersCoursesRepository.On("GetUserCourseEndroll", mock.Anything, mock.Anything).Return(usersCourseUser, nil).Once()
		userUsecase.On("GetCurrentUser", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, errors.New("use case err")).Once()

		_, err := usersCoursesService.GetUserCourseEndroll(context.Background(), &users_courses.User{UserID: "123"})

		assert.NotNil(t, err)
	})
}

func setupUpdateScoreCourse() {
	usersCoursesService = users_courses.NewUsersCoursesUsecase(&usersCoursesRepository, &userUsecase, &courseUsecase, time.Hour*1)
	usersCoursesDomain = users_courses.Domain{
		ID:          uuid.NewV4().String(),
		UserId:      uuid.NewV4().String(),
		CourseId:    uuid.NewV4().String(),
		Progres:     33,
		LastVideoId: "",
		LastModulId: "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func TestUsersCoursesUsecase_UpdateScoreCourse(t *testing.T) {
	t.Run("Test 1 | success", func(t *testing.T) {
		setupUpdateScoreCourse()

		usersCoursesRepository.On("UpdateScoreCourse", mock.Anything, mock.Anything).Return(usersCoursesDomain, nil).Once()

		result, err := usersCoursesService.UpdateScoreCourse(context.Background(), &users_courses.Domain{
			ID:       "123",
			UserId:   "234",
			CourseId: "345",
		})

		assert.Nil(t, err)
		assert.Equal(t, usersCoursesDomain.Progres, result.Progres)
	})
	t.Run("Test 2 | course id empty", func(t *testing.T) {
		setupUpdateScoreCourse()
		_, err := usersCoursesService.UpdateScoreCourse(context.Background(), &users_courses.Domain{
			ID:       "123",
			UserId:   "234",
			CourseId: "",
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
	t.Run("Test 3 | user id empty", func(t *testing.T) {
		setupUpdateScoreCourse()
		_, err := usersCoursesService.UpdateScoreCourse(context.Background(), &users_courses.Domain{
			ID:       "123",
			UserId:   "",
			CourseId: "345",
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("Test 4 | id empty", func(t *testing.T) {
		setupUpdateScoreCourse()
		_, err := usersCoursesService.UpdateScoreCourse(context.Background(), &users_courses.Domain{
			ID:       "",
			UserId:   "345",
			CourseId: "123",
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
	t.Run("Test 5 | error db", func(t *testing.T) {
		setupUpdateScoreCourse()

		usersCoursesRepository.On("UpdateScoreCourse", mock.Anything, mock.Anything).Return(usersCoursesDomain, errors.New("db err")).Once()

		_, err := usersCoursesService.UpdateScoreCourse(context.Background(), &users_courses.Domain{
			ID:       "123",
			UserId:   "234",
			CourseId: "345",
		})

		assert.NotNil(t, err)
	})
}

func setupOneUserCourse() {
	usersCoursesService = users_courses.NewUsersCoursesUsecase(&usersCoursesRepository, &userUsecase, &courseUsecase, time.Hour*1)
	usersCoursesDomain = users_courses.Domain{
		ID:          uuid.NewV4().String(),
		UserId:      "345",
		CourseId:    "123",
		Progres:     33,
		LastVideoId: "",
		LastModulId: "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func TestUsersCoursesUsecase_GetOneUserCourse(t *testing.T) {
	t.Run("Test 1 | success ", func(t *testing.T) {
		setupOneUserCourse()
		usersCoursesRepository.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(usersCoursesDomain, nil).Once()

		result, err := usersCoursesService.GetOneUserCourse(context.Background(), &users_courses.Domain{
			CourseId: "123",
			UserId:   "345",
		})
		assert.Nil(t, err)
		assert.Equal(t, usersCoursesDomain.CourseId, result.CourseId)
	})
	t.Run("Test 2 | err db", func(t *testing.T) {
		setupOneUserCourse()
		usersCoursesRepository.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(usersCoursesDomain, errors.New("err db")).Once()

		_, err := usersCoursesService.GetOneUserCourse(context.Background(), &users_courses.Domain{
			CourseId: "123",
			UserId:   "345",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test 3 | course id empty", func(t *testing.T) {
		setupOneUserCourse()

		_, err := usersCoursesService.GetOneUserCourse(context.Background(), &users_courses.Domain{
			CourseId: "",
			UserId:   "345",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
	t.Run("Test 4 | user id empty", func(t *testing.T) {
		setupOneUserCourse()

		_, err := usersCoursesService.GetOneUserCourse(context.Background(), &users_courses.Domain{
			CourseId: "123",
			UserId:   "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
}
