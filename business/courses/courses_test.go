package courses_test

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"profcourse/business/courses"
	_mocksCourseRepository "profcourse/business/courses/mocks"
	"profcourse/business/locals"
	_mocksLocalRepository "profcourse/business/locals/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var courseMysqlRepository _mocksCourseRepository.Repository
var localyRepository _mocksLocalRepository.Repository

var courseService courses.Usecase
var courseDomain courses.Domain
var localDomain locals.Domain
var listCourse []courses.Domain

func setupCreateCouse() {
	courseService = courses.NewCourseUseCase(&courseMysqlRepository, time.Hour*1, &localyRepository)
	courseDomain = courses.Domain{
		ID:            uuid.NewV4().String(),
		Title:         "Docker Pemula",
		Description:   "Docker untuk pemula",
		ImgUrl:        "./public/img/courses/iahguid.png",
		TeacherId:     uuid.NewV4().String(),
		TeacherName:   "",
		Status:        2,
		StatusText:    "",
		FileImage:     nil,
		CertificateId: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	localDomain = locals.Domain{
		File:        &multipart.FileHeader{},
		Destination: "/img/courses",
		ResultUrl:   "./public/img/courses/adfjakg.jpg",
		FileName:    "adfjakg.jpg",
	}
}

func TestCoursesUsecase_CreateCourse(t *testing.T) {
	t.Run("Test case 1 | title empaty", func(t *testing.T) {
		setupCreateCouse()
		_, err := courseService.CreateCourse(context.Background(), &courses.Domain{
			TeacherId:   uuid.NewV4().String(),
			Title:       "",
			Description: "Docker dari pemula"})
		assert.NotNil(t, err)
		assert.Equal(t, controller.TITLE_EMPTY, err)
	})
	t.Run("Test case 2 | decription empaty", func(t *testing.T) {
		setupCreateCouse()
		_, err := courseService.CreateCourse(context.Background(), &courses.Domain{Title: "Docker Pemula", TeacherId: uuid.NewV4().String()})
		assert.NotNil(t, err)
		assert.Equal(t, controller.DESC_EMPTY, err)
	})

	t.Run("Test case 2 | File empty", func(t *testing.T) {
		setupCreateCouse()
		_, err := courseService.CreateCourse(context.Background(), &courses.Domain{Title: "Docker Pemula", Description: "Docker untuk pemula", TeacherId: uuid.NewV4().String()})
		assert.NotNil(t, err)
		assert.Equal(t, controller.FILE_IMAGE_EMPTY, err)
	})

	t.Run("Test case 3 | Success", func(t *testing.T) {
		setupCreateCouse()
		localyRepository.On("UploadImage", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(localDomain, nil).Once()
		courseMysqlRepository.On("CreateCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		course, err := courseService.CreateCourse(context.Background(), &courses.Domain{Title: "Docker Pemula",
			Description: "Docker untuk pemula", TeacherId: uuid.NewV4().String(), FileImage: &multipart.FileHeader{}})
		assert.Nil(t, err)
		assert.Equal(t, courseDomain.Title, course.Title)
	})

	t.Run("Test case 4 | handle Error Local Upload", func(t *testing.T) {
		setupCreateCouse()
		localyRepository.On("UploadImage", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(locals.Domain{}, errors.New("Local error")).Once()
		_, err := courseService.CreateCourse(context.Background(), &courses.Domain{Title: "Docker Pemula",
			Description: "Docker untuk pemula", TeacherId: uuid.NewV4().String(), FileImage: &multipart.FileHeader{}})
		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Handle Error DB", func(t *testing.T) {
		setupCreateCouse()
		localyRepository.On("UploadImage", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(localDomain, nil).Once()
		courseMysqlRepository.On("CreateCourse", mock.Anything, mock.Anything).Return(&courses.Domain{}, errors.New("Error DB")).Once()
		_, err := courseService.CreateCourse(context.Background(), &courses.Domain{Title: "Docker Pemula",
			Description: "Docker untuk pemula", TeacherId: uuid.NewV4().String(), FileImage: &multipart.FileHeader{}})
		assert.NotNil(t, err)
	})

}

func setupGetAllCourses() {
	courseService = courses.NewCourseUseCase(&courseMysqlRepository, time.Hour*1, &localyRepository)
	courseDomain = courses.Domain{
		ID:            uuid.NewV4().String(),
		Title:         "Docker Pemula",
		Description:   "Docker untuk pemula",
		ImgUrl:        "./public/img/courses/iahguid.png",
		TeacherId:     uuid.NewV4().String(),
		TeacherName:   "",
		Status:        2,
		StatusText:    "",
		FileImage:     nil,
		CertificateId: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	listCourse = []courses.Domain{
		courseDomain, courseDomain,
	}
}

func TestCoursesUsecase_GetAllCourses(t *testing.T) {
	t.Run("Test case 1 | invalid params sort", func(t *testing.T) {
		setupGetAllCourses()

		_, err := courseService.GetAllCourses(context.Background(), &courses.Domain{Sort: "adfsd", SortBy: "asc"})
		assert.NotNil(t, err)
		assert.Equal(t, controller.INVALID_PARAMS, err)
	})

	t.Run("Test case 2 | invalid params sort by", func(t *testing.T) {
		setupGetAllCourses()

		_, err := courseService.GetAllCourses(context.Background(), &courses.Domain{Sort: "", SortBy: "ascsd"})
		assert.NotNil(t, err)
		assert.Equal(t, controller.INVALID_PARAMS, err)
	})

	t.Run("Test case 3 | success get list courses with limit 2", func(t *testing.T) {
		setupGetAllCourses()
		courseMysqlRepository.On("GetAllCourses", mock.Anything, mock.Anything).Return(&listCourse, nil).Once()
		coursesList, err := courseService.GetAllCourses(context.Background(), &courses.Domain{
			Limit: 2,
			SortBy:  "dsc",
		})
		assert.Nil(t, err)
		assert.Len(t, *coursesList, 2)
	})

	t.Run("Test case 4 | Handle error db get all courses", func(t *testing.T) {
		setupGetAllCourses()
		courseMysqlRepository.On("GetAllCourses", mock.Anything, mock.Anything).Return(&[]courses.Domain{}, errors.New("Error DB")).Once()
		_, err := courseService.GetAllCourses(context.Background(), &courses.Domain{
			Limit: 2,
		})
		assert.NotNil(t, err)
	})

}
