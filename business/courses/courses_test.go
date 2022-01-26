package courses_test

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"profcourse/business/courses"
	_mocksCourseRepository "profcourse/business/courses/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var courseMysqlRepository _mocksCourseRepository.Repository

var courseService courses.Usecase
var courseDomain courses.Domain
var courseSummary courses.Summary
var listCourse []courses.Domain

func setGetCountCourse() {
	courseService = courses.NewCourseUseCase(&courseMysqlRepository, time.Hour*1)
	courseSummary = courses.Summary{
		CountCourse: 9,
	}
}

func TestCoursesUsecase_GetCountCourse(t *testing.T) {
	t.Run("Test case 1 | Success get count course", func(t *testing.T) {
		setGetCountCourse()
		courseMysqlRepository.On("GetCountCourse", mock.Anything).Return(&courseSummary, nil).Once()
		summary, err := courseService.GetCountCourse(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 9, summary.CountCourse)
	})
	t.Run("Test case 2 | handle err db", func(t *testing.T) {
		setGetCountCourse()
		courseMysqlRepository.On("GetCountCourse", mock.Anything).Return(&courses.Summary{}, errors.New("hahaha")).Once()
		_, err := courseService.GetCountCourse(context.Background())
		assert.NotNil(t, err)
	})
}

func setupCreateCouse() {
	courseService = courses.NewCourseUseCase(&courseMysqlRepository, time.Hour*1)
	courseDomain = courses.Domain{
		ID:            uuid.NewV4().String(),
		Title:         "Docker Pemula",
		Description:   "Docker untuk pemula",
		ImgUrl:        "./public/img/courses/iahguid.png",
		TeacherId:     uuid.NewV4().String(),
		TeacherName:   "",
		Status:        2,
		StatusText:    "",
		CertificateId: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
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
		assert.Equal(t, controller.IMAGE_EMPTY, err)
	})

	t.Run("Test case 3 | Success", func(t *testing.T) {
		setupCreateCouse()

		courseMysqlRepository.On("CreateCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		course, err := courseService.CreateCourse(context.Background(), &courses.Domain{Title: "Docker Pemula",
			Description: "Docker untuk pemula", TeacherId: uuid.NewV4().String(), ImgUrl: "dafdsfj"})
		assert.Nil(t, err)
		assert.Equal(t, courseDomain.Title, course.Title)
	})

	t.Run("Test case 5 | Handle Error DB", func(t *testing.T) {
		setupCreateCouse()

		courseMysqlRepository.On("CreateCourse", mock.Anything, mock.Anything).Return(&courses.Domain{}, errors.New("Error DB")).Once()
		_, err := courseService.CreateCourse(context.Background(), &courses.Domain{Title: "Docker Pemula",
			Description: "Docker untuk pemula", TeacherId: uuid.NewV4().String(), ImgUrl: "dasfsdfg"})
		assert.NotNil(t, err)
	})

}

func setupGetAllCourses() {
	courseService = courses.NewCourseUseCase(&courseMysqlRepository, time.Hour*1)
	courseDomain = courses.Domain{
		ID:            uuid.NewV4().String(),
		Title:         "Docker Pemula",
		Description:   "Docker untuk pemula",
		ImgUrl:        "./public/img/courses/iahguid.png",
		TeacherId:     uuid.NewV4().String(),
		TeacherName:   "",
		Status:        2,
		StatusText:    "",
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
			Limit:  2,
			SortBy: "dsc",
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

func setupGetOneCourses() {
	courseService = courses.NewCourseUseCase(&courseMysqlRepository, time.Hour*1)
	courseDomain = courses.Domain{
		ID:            uuid.NewV4().String(),
		Title:         "Docker Pemula",
		Description:   "Docker untuk pemula",
		ImgUrl:        "./public/img/courses/iahguid.png",
		TeacherId:     uuid.NewV4().String(),
		TeacherName:   "",
		Status:        2,
		StatusText:    "",
		CertificateId: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	listCourse = []courses.Domain{
		courseDomain, courseDomain,
	}
}

func TestCoursesUsecase_GetOneCourse(t *testing.T) {
	t.Run("Test 1 | Success get one course", func(t *testing.T) {
		setupGetOneCourses()
		courseMysqlRepository.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		_, err := courseService.GetOneCourse(context.Background(), &courses.Domain{ID: "ece59f46-71df-42d2-8885-619061e2def8"})
		assert.Nil(t, err)
	})
	t.Run("Test 2 | Handle Error db get one course", func(t *testing.T) {
		setupGetOneCourses()
		courseMysqlRepository.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courses.Domain{}, errors.New("Error db")).Once()
		_, err := courseService.GetOneCourse(context.Background(), &courses.Domain{ID: "ece59f46-71df-42d2-8885-619061e2def8"})
		assert.NotNil(t, err)
	})
}

func setUpUpdateCourse() {
	courseService = courses.NewCourseUseCase(&courseMysqlRepository, time.Hour*1)
	courseDomain = courses.Domain{
		ID:            uuid.NewV4().String(),
		Title:         "Docker Pemula",
		Description:   "Docker untuk pemula",
		ImgUrl:        "./public/img/courses/iahguid.png",
		TeacherId:     uuid.NewV4().String(),
		TeacherName:   "",
		Status:        2,
		StatusText:    "",
		CertificateId: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
}

func TestCoursesUsecase_UpdateCourse(t *testing.T) {
	t.Run("Test case 1 | seuccess update course dengan admin", func(t *testing.T) {
		setUpUpdateCourse()

		courseMysqlRepository.On("UpdateCourseForAdmin", mock.Anything, mock.Anything).Return(courseDomain, nil).Once()

		result, err := courseService.UpdateCourse(context.Background(), &courses.Domain{ID: "asdas", Title: "Docker Pemula", Description: "Docker untuk pemula", ImgUrl: "https://placeimg.com/640/480/tech"}, &courses.Token{UserId: "535763b7-d499-4478-87be-36463758474c", Role: 1})

		assert.Nil(t, err)
		assert.Equal(t, courseDomain.Title, result.Title)
	})

	t.Run("Test case 1 | seuccess update course dengan user", func(t *testing.T) {
		setUpUpdateCourse()

		courseMysqlRepository.On("UpdateCourseForUser", mock.Anything, mock.Anything, mock.Anything).Return(courseDomain, nil).Once()

		result, err := courseService.UpdateCourse(context.Background(), &courses.Domain{ID: "asdas", Title: "Docker Pemula", Description: "Docker untuk pemula", ImgUrl: "https://placeimg.com/640/480/tech"}, &courses.Token{UserId: "535763b7-d499-4478-87be-36463758474c", Role: 2})

		assert.Nil(t, err)
		assert.Equal(t, courseDomain.Title, result.Title)
	})

	t.Run("Test case 2 | error title kosong", func(t *testing.T) {
		setUpUpdateCourse()

		_, err := courseService.UpdateCourse(context.Background(), &courses.Domain{ID: "f9c00064-bb11-46d8-bbde-bbda5f97b831", Title: "", Description: "Docker untuk pemula", ImgUrl: "https://placeimg.com/640/480/tech"}, &courses.Token{UserId: "535763b7-d499-4478-87be-36463758474c", Role: 1})

		assert.NotNil(t, err)
		assert.Equal(t, controller.TITLE_EMPTY, err)
	})
	t.Run("Test case 3 | error id kosong", func(t *testing.T) {
		setUpUpdateCourse()

		_, err := courseService.UpdateCourse(context.Background(), &courses.Domain{ID: "", Title: "Docker", Description: "Docker untuk pemula", ImgUrl: "https://placeimg.com/640/480/tech"}, &courses.Token{UserId: "535763b7-d499-4478-87be-36463758474c", Role: 1})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})

	t.Run("Test case 4 | error description kosong", func(t *testing.T) {
		setUpUpdateCourse()

		_, err := courseService.UpdateCourse(context.Background(), &courses.Domain{ID: "f9c00064-bb11-46d8-bbde-bbda5f97b831", Title: "Docker", Description: "", ImgUrl: "https://placeimg.com/640/480/tech"}, &courses.Token{UserId: "535763b7-d499-4478-87be-36463758474c", Role: 1})

		assert.NotNil(t, err)
		assert.Equal(t, controller.DESC_EMPTY, err)
	})

	t.Run("Test case 1 | image url kosong", func(t *testing.T) {
		setUpUpdateCourse()

		_, err := courseService.UpdateCourse(context.Background(), &courses.Domain{ID: "asdas", Title: "Docker Pemula", Description: "Docker untuk pemula", ImgUrl: ""}, &courses.Token{UserId: "535763b7-d499-4478-87be-36463758474c", Role: 2})

		assert.NotNil(t, err)
		assert.Equal(t, controller.IMAGE_EMPTY, err)
	})
}

func setUpDeleteCourse() {
	courseService = courses.NewCourseUseCase(&courseMysqlRepository, time.Hour*1)
	courseDomain = courses.Domain{
		ID:            uuid.NewV4().String(),
		Title:         "Docker Pemula",
		Description:   "Docker untuk pemula",
		ImgUrl:        "./public/img/courses/iahguid.png",
		TeacherId:     uuid.NewV4().String(),
		TeacherName:   "",
		Status:        2,
		StatusText:    "",
		CertificateId: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
}

func TestCoursesUsecase_DeleteCourse(t *testing.T) {
	t.Run("Test case 1 | Success delete course dengan admin", func(t *testing.T) {
		setUpDeleteCourse()
		courseMysqlRepository.On("DeleteCourseForAdmin", mock.Anything, mock.AnythingOfType("string")).Return(courseDomain, nil).Once()
		_, err := courseService.DeleteCourse(context.Background(), "535763b7-d499-4478-87be-36463758474c", courses.Token{
			UserId: "f9c00064-bb11-46d8-bbde-bbda5f97b831",
			Role:   1,
		})
		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Success delete course dengan user", func(t *testing.T) {
		setUpDeleteCourse()
		courseMysqlRepository.On("DeleteCourseForUser", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(courseDomain, nil).Once()
		_, err := courseService.DeleteCourse(context.Background(), "535763b7-d499-4478-87be-36463758474c", courses.Token{
			UserId: "f9c00064-bb11-46d8-bbde-bbda5f97b831",
			Role:   2,
		})
		assert.Nil(t, err)
	})

	t.Run("Test case 3 | Error karena tidak ada id course", func(t *testing.T) {
		setUpDeleteCourse()
		_, err := courseService.DeleteCourse(context.Background(), "", courses.Token{
			UserId: "f9c00064-bb11-46d8-bbde-bbda5f97b831",
			Role:   2,
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
}

func TestCoursesUsecase_GetAllCourseUser(t *testing.T) {
	t.Run("test case 1 | Success", func(t *testing.T) {
		setupGetAllCourses()
		courseMysqlRepository.On("GetAllCourseUser", mock.Anything, mock.Anything).Return(listCourse, nil).Once()

		_, err := courseService.GetAllCourseUser(context.Background(), &courses.Domain{TeacherId: "123"})

		assert.Nil(t, err)
	})
	t.Run("test case 2 | Error user id empty", func(t *testing.T) {
		setupGetAllCourses()

		_, err := courseService.GetAllCourseUser(context.Background(), &courses.Domain{TeacherId: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("test case 3 | Success walau sort diisi dsc", func(t *testing.T) {
		setupGetAllCourses()
		courseMysqlRepository.On("GetAllCourseUser", mock.Anything, mock.Anything).Return(listCourse, nil).Once()

		_, err := courseService.GetAllCourseUser(context.Background(), &courses.Domain{TeacherId: "123", Sort: "dsc"})

		assert.Nil(t, err)
	})
	t.Run("test case 4 | err db", func(t *testing.T) {
		setupGetAllCourses()
		courseMysqlRepository.On("GetAllCourseUser", mock.Anything, mock.Anything).Return(listCourse, errors.New("err")).Once()

		_, err := courseService.GetAllCourseUser(context.Background(), &courses.Domain{TeacherId: "123", Sort: "dsc"})

		assert.NotNil(t, err)
	})
}
