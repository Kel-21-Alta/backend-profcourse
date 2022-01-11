package moduls_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/courses"
	_mocksCourseUsecase "profcourse/business/courses/mocks"
	"profcourse/business/moduls"
	_mysqlMockModulRepository "profcourse/business/moduls/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var courseUsecase _mocksCourseUsecase.Usecase
var mysqlModulrepository _mysqlMockModulRepository.Repository

var modulDomain moduls.Domain
var modulServices moduls.Usecase
var courseDomain courses.Domain

func setUpCreateModul() {
	modulServices = moduls.NewModulUsecase(&mysqlModulrepository, &courseUsecase, time.Hour*1)
	courseDomain = courses.Domain{
		ID:          "1234",
		Title:       "",
		Description: "",
		ImgUrl:      "",
		TeacherId:   "123",
		TeacherName: "",
		Status:      0,
	}
}

func TestModulUsecase_CreateModul(t *testing.T) {
	t.Run("Test case 1 | seorang user success create modul miliknya ", func(t *testing.T) {
		setUpCreateModul()

		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		mysqlModulrepository.On("CreateModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()

		_, err := modulServices.CreateModul(context.Background(), &moduls.Domain{CourseId: "1234", UserMakeModul: "123", Title: "Pengenalan", Order: 1, RoleUser: 2})
		assert.Nil(t, err)
	})

	t.Run("Test case 2 | seorang admin success create modul miliknya ", func(t *testing.T) {
		setUpCreateModul()

		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		mysqlModulrepository.On("CreateModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()

		_, err := modulServices.CreateModul(context.Background(), &moduls.Domain{CourseId: "1234", UserMakeModul: "1", Title: "Pengenalan", Order: 1, RoleUser: 1})
		assert.Nil(t, err)
	})

	t.Run("Test case 3 | Menampilkan error jika user yang manambahkan modul bukan pemilik dari modul ", func(t *testing.T) {
		setUpCreateModul()

		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		mysqlModulrepository.On("CreateModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()

		_, err := modulServices.CreateModul(context.Background(), &moduls.Domain{CourseId: "1234", UserMakeModul: "1", Title: "Pengenalan", Order: 1, RoleUser: 2})
		assert.NotNil(t, err)
		assert.Equal(t, controller.FORBIDDIN_USER, err)
	})

	t.Run("Test case 4 | Menampilkan error jika user title tidak ada ", func(t *testing.T) {
		setUpCreateModul()

		_, err := modulServices.CreateModul(context.Background(), &moduls.Domain{CourseId: "1234", UserMakeModul: "1", Title: "", Order: 1, RoleUser: 2})
		assert.NotNil(t, err)
		assert.Equal(t, controller.TITLE_EMPTY, err)
	})

	t.Run("Test case 5 | Menampilkan error jika order bernilai 0 ", func(t *testing.T) {
		setUpCreateModul()

		_, err := modulServices.CreateModul(context.Background(), &moduls.Domain{CourseId: "1234", UserMakeModul: "1", Title: "1", Order: 0, RoleUser: 2})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ORDER_MODUL_EMPTY, err)
	})

	t.Run("Test case 6 | Menampilkan error jika course id kosong ", func(t *testing.T) {
		setUpCreateModul()

		_, err := modulServices.CreateModul(context.Background(), &moduls.Domain{CourseId: "", UserMakeModul: "1", Title: "1", Order: 1, RoleUser: 2})
		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})

	t.Run("Test case 7 | handle error form course usecase ", func(t *testing.T) {
		setUpCreateModul()

		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courses.Domain{}, errors.New("error")).Once()

		_, err := modulServices.CreateModul(context.Background(), &moduls.Domain{CourseId: "1234", UserMakeModul: "123", Title: "Pengenalan", Order: 1, RoleUser: 2})
		assert.NotNil(t, err)
	})

}
