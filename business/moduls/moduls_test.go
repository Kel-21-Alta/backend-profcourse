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
var materi moduls.Materi

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

func setUpGetOneModul() {
	modulServices = moduls.NewModulUsecase(&mysqlModulrepository, &courseUsecase, time.Hour*1)
	modulDomain = moduls.Domain{
		ID:            "985c0e69-8a38-4774-9a12-2c279f7f258d",
		Title:         "Pengenalan Docker",
		Order:         1,
		CourseId:      "41f99e51-bd6a-4e67-b631-25a58acb39f4",
		Materi:        []moduls.Materi{materi, materi},
		JumlahMateri:  5,
		UserMakeModul: "9cc1fc86-4c02-4fe2-8c57-93b4c56225ee",
		RoleUser:      1,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
}

func TestModulUsecase_GetOneModul(t *testing.T) {
	t.Run("Test case 1 | success get one modul", func(t *testing.T) {
		setUpGetOneModul()
		mysqlModulrepository.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		result, err := modulServices.GetOneModul(context.Background(), &moduls.Domain{ID: modulDomain.ID})
		assert.Nil(t, err)
		assert.Equal(t, modulDomain.ID, result.ID)
	})
	t.Run("Test case 2 | handle error modul id kosong", func(t *testing.T) {
		setUpGetOneModul()
		_, err := modulServices.GetOneModul(context.Background(), &moduls.Domain{ID: ""})
		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_MODUL_ID, err)
	})
}

func setUpUpdateMateri() {
	modulServices = moduls.NewModulUsecase(&mysqlModulrepository, &courseUsecase, time.Hour*1)
	modulDomain = moduls.Domain{
		ID:            "985c0e69-8a38-4774-9a12-2c279f7f258d",
		Title:         "Pengenalan Docker",
		Order:         1,
		CourseId:      "41f99e51-bd6a-4e67-b631-25a58acb39f4",
		Materi:        []moduls.Materi{materi, materi},
		JumlahMateri:  5,
		UserMakeModul: "9cc1fc86-4c02-4fe2-8c57-93b4c56225ee",
		RoleUser:      1,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	courseDomain = courses.Domain{
		ID:          "41f99e51-bd6a-4e67-b631-25a58acb39f4",
		Title:       "",
		Description: "",
		ImgUrl:      "",
		TeacherId:   modulDomain.UserMakeModul,
		TeacherName: "",
		Status:      0,
	}
}

func TestModulUsecase_UpdateModul(t *testing.T) {
	t.Run("Test case 1 | success update modul", func(t *testing.T) {
		setUpUpdateMateri()
		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		mysqlModulrepository.On("UpdateModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()

		result, err := modulServices.UpdateModul(context.Background(), &moduls.Domain{
			ID:            modulDomain.ID,
			Title:         modulDomain.Title,
			Order:         modulDomain.Order,
			CourseId:      courseDomain.ID,
			UserMakeModul: courseDomain.TeacherId,
			RoleUser:      2,
		})

		assert.Nil(t, err)
		assert.Equal(t, modulDomain.Title, result.Title)
		assert.Equal(t, courseDomain.TeacherId, result.UserMakeModul)
	})
	t.Run("Test case 2 | title kosong", func(t *testing.T) {
		_, err := modulServices.UpdateModul(context.Background(), &moduls.Domain{Title: ""})
		assert.Equal(t, controller.TITLE_EMPTY, err)
	})
	t.Run("Test case 3 | Order 0", func(t *testing.T) {
		_, err := modulServices.UpdateModul(context.Background(), &moduls.Domain{Title: "Pengenalan", Order: 0})
		assert.Equal(t, controller.ORDER_MODUL_EMPTY, err)
	})
	t.Run("Test case 4 | Course id kosong", func(t *testing.T) {
		_, err := modulServices.UpdateModul(context.Background(), &moduls.Domain{Title: "Pengenalan", Order: 1, CourseId: ""})
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
	t.Run("Test case 5 | Handle error form course usecase", func(t *testing.T) {
		setUpUpdateMateri()
		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, errors.New("Err form usercase course")).Once()
		_, err := modulServices.UpdateModul(context.Background(), &moduls.Domain{Title: "Pengenalan", Order: 1, CourseId: modulDomain.CourseId})
		assert.NotNil(t, err)
	})
	t.Run("Test case 6 | Handle error user tidak diizinkan mengubah modul", func(t *testing.T) {
		setUpUpdateMateri()
		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		_, err := modulServices.UpdateModul(context.Background(), &moduls.Domain{Title: "Pengenalan", Order: 1, CourseId: modulDomain.CourseId, RoleUser: 2, UserMakeModul: "asd"})
		assert.NotNil(t, err)
		assert.Equal(t, controller.FORBIDDIN_USER, err)
	})
	t.Run("Test case 6 | Handle error database modul", func(t *testing.T) {
		setUpUpdateMateri()
		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		mysqlModulrepository.On("UpdateModul", mock.Anything, mock.Anything).Return(modulDomain, errors.New("Err db")).Once()
		_, err := modulServices.UpdateModul(context.Background(), &moduls.Domain{Title: "Pengenalan", Order: 1, CourseId: modulDomain.CourseId, RoleUser: 1, UserMakeModul: "asd"})
		assert.NotNil(t, err)
	})
}
