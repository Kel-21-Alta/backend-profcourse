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
var listModul []moduls.Domain
var scoreModulDomain moduls.ScoreUserModul

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
	t.Run("Test case 7 | Handle error database modul", func(t *testing.T) {
		setUpUpdateMateri()
		courseUsecase.On("GetOneCourse", mock.Anything, mock.Anything).Return(&courseDomain, nil).Once()
		mysqlModulrepository.On("UpdateModul", mock.Anything, mock.Anything).Return(modulDomain, errors.New("Err db")).Once()
		_, err := modulServices.UpdateModul(context.Background(), &moduls.Domain{Title: "Pengenalan", Order: 1, CourseId: modulDomain.CourseId, RoleUser: 1, UserMakeModul: "asd"})
		assert.NotNil(t, err)
	})
}

func setUpDeleteModul() {
	modulServices = moduls.NewModulUsecase(&mysqlModulrepository, &courseUsecase, time.Hour*1)
	modulDomain = moduls.Domain{UserMakeModul: "9cc1fc86-4c02-4fe2-8c57-93b4c56225ee"}
}

func TestModulUsecase_DeleteModul(t *testing.T) {
	t.Run("Test case 1 | berhasil delete modul", func(t *testing.T) {
		setUpDeleteModul()
		mysqlModulrepository.On("GetOneModulWithCourse", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		mysqlModulrepository.On(
			"DeleteModul",
			mock.Anything,
			mock.AnythingOfType("string")).Return(moduls.Message("Success"), nil).Once()

		result, err := modulServices.DeleteModul(
			context.Background(),
			&moduls.Domain{ID: "4a418b6a-68b8-436e-ad29-12504d63bb2d", UserMakeModul: "9cc1fc86-4c02-4fe2-8c57-93b4c56225ee"})

		assert.Nil(t, err)
		assert.NotEqual(t, "", result)
	})
	t.Run("Test case 2 | admin boleh mendelete modul user", func(t *testing.T) {
		setUpDeleteModul()
		mysqlModulrepository.On("GetOneModulWithCourse", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		mysqlModulrepository.On(
			"DeleteModul",
			mock.Anything,
			mock.AnythingOfType("string")).Return(moduls.Message("Success"), nil).Once()

		_, err := modulServices.DeleteModul(
			context.Background(),
			&moduls.Domain{ID: "4a418b6a-68b8-436e-ad29-12504d63bb2d", RoleUser: 1, UserMakeModul: "cc1fc86-4c02-4fe2-8c57-93b4c56225ee"})

		assert.Nil(t, err)
	})
	t.Run("Test case 3 | user tidak diizinkan mendelete", func(t *testing.T) {
		setUpDeleteModul()
		mysqlModulrepository.On("GetOneModulWithCourse", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		mysqlModulrepository.On(
			"DeleteModul",
			mock.Anything,
			mock.AnythingOfType("string")).Return(moduls.Message("Success"), nil).Once()

		_, err := modulServices.DeleteModul(
			context.Background(),
			&moduls.Domain{ID: "4a418b6a-68b8-436e-ad29-12504d63bb2d", RoleUser: 2, UserMakeModul: "cc1fc86-4c02-4fe2-8c57-93b4c56225ee"})

		assert.Equal(t, controller.FORBIDDIN_USER, err)
	})
	t.Run("Test case 4 | error id modul empty", func(t *testing.T) {
		setUpDeleteModul()

		_, err := modulServices.DeleteModul(
			context.Background(),
			&moduls.Domain{ID: "", RoleUser: 2, UserMakeModul: "9cc1fc86-4c02-4fe2-8c57-93b4c56225ee"})

		assert.Equal(t, controller.ID_EMPTY, err)
	})
}

func setUpGetAllModul() {
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
	listModul = []moduls.Domain{modulDomain, modulDomain}
}

func TestModulUsecase_GetAllModulCourse(t *testing.T) {
	t.Run("test case 1 | succes get modul form course", func(t *testing.T) {
		setUpGetAllModul()

		mysqlModulrepository.On("GetAllModulCourse", mock.Anything, mock.Anything).Return(listModul, nil).Once()

		_, err := modulServices.GetAllModulCourse(context.Background(), &moduls.Domain{CourseId: "41f99e51-bd6a-4e67-b631-25a58acb39f4"})

		assert.Nil(t, err)
	})
	t.Run("test case 1 | succes get modul form course", func(t *testing.T) {
		setUpGetAllModul()

		_, err := modulServices.GetAllModulCourse(context.Background(), &moduls.Domain{CourseId: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
	t.Run("test case 1 | succes get modul form course", func(t *testing.T) {
		setUpGetAllModul()

		mysqlModulrepository.On("GetAllModulCourse", mock.Anything, mock.Anything).Return(listModul, errors.New("err db")).Once()

		_, err := modulServices.GetAllModulCourse(context.Background(), &moduls.Domain{CourseId: "41f99e51-bd6a-4e67-b631-25a58acb39f4"})

		assert.NotNil(t, err)
	})
}

func setUpCalculateScoreCourse() {
	modulServices = moduls.NewModulUsecase(&mysqlModulrepository, &courseUsecase, time.Hour*1)
	scoreModulDomain = moduls.ScoreUserModul{
		Nilai: 4,
	}
}

func TestModulUsecase_CalculateScoreCourse(t *testing.T) {
	t.Run("Test case 1 | Success calculate", func(t *testing.T) {
		setUpCalculateScoreCourse()
		mysqlModulrepository.On("CalculateScoreCourse", mock.Anything, mock.Anything).Return(scoreModulDomain, nil).Once()

		result, err := modulServices.CalculateScoreCourse(context.Background(), &moduls.ScoreUserModul{UserCourseId: "123"})

		assert.Nil(t, err)
		assert.NotEqual(t, "", result.Nilai)
	})
	t.Run("Test case 2 | User course id empty", func(t *testing.T) {
		setUpCalculateScoreCourse()

		_, err := modulServices.CalculateScoreCourse(context.Background(), &moduls.ScoreUserModul{UserCourseId: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
	t.Run("Test case 3 | Error db", func(t *testing.T) {
		setUpCalculateScoreCourse()
		mysqlModulrepository.On("CalculateScoreCourse", mock.Anything, mock.Anything).Return(scoreModulDomain, errors.New("error db")).Once()

		_, err := modulServices.CalculateScoreCourse(context.Background(), &moduls.ScoreUserModul{UserCourseId: "123"})

		assert.NotNil(t, err)
	})
}

func setUpCreateScoreModul() {
	modulServices = moduls.NewModulUsecase(&mysqlModulrepository, &courseUsecase, time.Hour*1)
	scoreModulDomain = moduls.ScoreUserModul{
		Nilai:        4,
		ModulID:      "123",
		UserCourseId: "1234",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

func TestModulUsecase_CreateScoreModul(t *testing.T) {
	t.Run("Test case 1 | success create score modul", func(t *testing.T) {
		setUpCreateScoreModul()
		mysqlModulrepository.On("CreateScoreModul", mock.Anything, mock.Anything).Return(scoreModulDomain, nil).Once()

		result, err := modulServices.CreateScoreModul(context.Background(), &scoreModulDomain)

		assert.Nil(t, err)
		assert.Equal(t, scoreModulDomain.ModulID, result.ModulID)
	})
	t.Run("Test case 2 | modul id empty", func(t *testing.T) {
		setUpCreateScoreModul()

		_, err := modulServices.CreateScoreModul(context.Background(), &moduls.ScoreUserModul{
			ModulID:      "",
			UserCourseId: "123",
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_MODUL_ID, err)
	})
	t.Run("Test case 3 | user course id empty", func(t *testing.T) {
		setUpCreateScoreModul()

		_, err := modulServices.CreateScoreModul(context.Background(), &moduls.ScoreUserModul{
			ModulID:      "123",
			UserCourseId: "",
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
	t.Run("Test case 4 | err db", func(t *testing.T) {
		setUpCreateScoreModul()
		mysqlModulrepository.On("CreateScoreModul", mock.Anything, mock.Anything).Return(scoreModulDomain, errors.New("error ni")).Once()

		_, err := modulServices.CreateScoreModul(context.Background(), &scoreModulDomain)

		assert.NotNil(t, err)
	})
}
