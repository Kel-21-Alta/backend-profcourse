package quizs_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/moduls"
	_mockModulUsecase "profcourse/business/moduls/mocks"
	"profcourse/business/quizs"
	_mocksQuizRepository "profcourse/business/quizs/mocks"
	"profcourse/business/users_courses"
	_mocksUsersCoursesUsecase "profcourse/business/users_courses/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var mysqlQuizsRepository _mocksQuizRepository.Repository
var userCourseUsecase _mocksUsersCoursesUsecase.Usecase
var modulsUsecase _mockModulUsecase.Usecase
var usecaseQuiz _mocksUsersCoursesUsecase.Usecase

var quizsService quizs.Usecase
var quizsDomain quizs.Domain
var listDomain []quizs.Domain
var modulDomain moduls.Domain
var userCourseDomain users_courses.Domain
var scoreUserModulDomain moduls.ScoreUserModul

func setUpCreateQuizs() {
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, &modulsUsecase, &userCourseUsecase, time.Hour*1)
	quizsDomain = quizs.Domain{
		ID:         "7c1ec4be-8565-4b25-82cf-244d7730c398",
		Pilihan:    []string{"a", "b", "c"},
		Pertanyaan: "Makan apa?",
		Jawaban:    "a",
		ModulId:    "36d8d8bc-87cb-467d-97c0-2902920457df",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}

func TestNewQuizUsecase(t *testing.T) {
	t.Run("Test case 1 | success create quiz", func(t *testing.T) {
		setUpCreateQuizs()
		mysqlQuizsRepository.On("CreateQuiz", mock.Anything, mock.Anything).Return(quizsDomain, nil).Once()
		result, err := quizsService.CreateQuiz(context.Background(), &quizs.Domain{
			Pilihan:    []string{"a", "b", "c"},
			Pertanyaan: "Makan apa?",
			Jawaban:    "a",
			ModulId:    "36d8d8bc-87cb-467d-97c0-2902920457df",
		})
		assert.Nil(t, err)
		assert.Equal(t, quizsDomain.ID, result.ID)
	})
	t.Run("Test case 2 | success create quiz", func(t *testing.T) {
		setUpCreateQuizs()
		mysqlQuizsRepository.On("CreateQuiz", mock.Anything, mock.Anything).Return(quizs.Domain{}, errors.New("db error")).Once()
		_, err := quizsService.CreateQuiz(context.Background(), &quizs.Domain{
			Pilihan:    []string{"a", "b", "c"},
			Pertanyaan: "Makan apa?",
			Jawaban:    "a",
			ModulId:    "36d8d8bc-87cb-467d-97c0-2902920457df",
		})
		assert.NotNil(t, err)
	})
}

func TestQuizeUsecase_ValidasiQuiz(t *testing.T) {
	t.Run("Test case 1 | validasi berhasil", func(t *testing.T) {
		setUpCreateQuizs()
		result, err := quizsService.ValidasiQuiz(context.Background(), &quizs.Domain{
			Pilihan:    []string{"a", "b", "c"},
			Pertanyaan: "Makan apa?",
			Jawaban:    "a",
			ModulId:    "36d8d8bc-87cb-467d-97c0-2902920457df",
		})
		assert.Nil(t, err)
		assert.Equal(t, "Makan apa?", result.Pertanyaan)
	})
	t.Run("Test case 2 | handle modulid empty", func(t *testing.T) {
		setUpCreateQuizs()
		_, err := quizsService.ValidasiQuiz(context.Background(), &quizs.Domain{
			Pilihan:    []string{"a", "b", "c"},
			Pertanyaan: "Makan apa?",
			Jawaban:    "a",
			ModulId:    "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_MODUL_ID, err)
	})
	t.Run("Test case 3 | handle jawaban empty", func(t *testing.T) {
		setUpCreateQuizs()
		_, err := quizsService.ValidasiQuiz(context.Background(), &quizs.Domain{
			Pilihan:    []string{"a", "b", "c"},
			Pertanyaan: "Makan apa?",
			Jawaban:    "",
			ModulId:    "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.JAWABAN_QUIZ_EMPTY, err)
	})
	t.Run("Test case 4 | handdle pertanyaan kosong", func(t *testing.T) {
		setUpCreateQuizs()
		_, err := quizsService.ValidasiQuiz(context.Background(), &quizs.Domain{
			Pilihan:    []string{"a", "b", "c"},
			Pertanyaan: "",
			Jawaban:    "a",
			ModulId:    "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.PERTANYAAN_QUIZ_EMPTY, err)
	})
	t.Run("Test case 5 | handle error pilihan kosong", func(t *testing.T) {
		setUpCreateQuizs()
		_, err := quizsService.ValidasiQuiz(context.Background(), &quizs.Domain{
			Pilihan:    nil,
			Pertanyaan: "Makan apa?",
			Jawaban:    "a",
			ModulId:    "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.PILIHAN_QUIZ_EMPTY, err)
	})
	t.Run("Test case 6 | handle error pilihan kurang dari 2", func(t *testing.T) {
		setUpCreateQuizs()
		_, err := quizsService.ValidasiQuiz(context.Background(), &quizs.Domain{
			Pilihan:    []string{"a"},
			Pertanyaan: "Makan apa?",
			Jawaban:    "a",
			ModulId:    "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.PILIHAN_QUIZ_MINUS, err)
	})
}

func setUpUpdateQuiz() {
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, &modulsUsecase, &userCourseUsecase, time.Hour*1)

	quizsDomain = quizs.Domain{
		ID:         "7c1ec4be-8565-4b25-82cf-244d7730c398",
		Pilihan:    []string{"a", "b", "c"},
		Pertanyaan: "Makan apa?",
		Jawaban:    "a",
		ModulId:    "36d8d8bc-87cb-467d-97c0-2902920457df",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}

func TestQuizeUsecase_UpdateQuiz(t *testing.T) {
	t.Run("Test 1 | success update quiz", func(t *testing.T) {
		setUpUpdateQuiz()
		mysqlQuizsRepository.On("UpdateQuiz", mock.Anything, mock.Anything).Return(quizsDomain, nil).Once()

		result, err := quizsService.UpdateQuiz(context.Background(), &quizsDomain)
		assert.Nil(t, err)
		assert.Equal(t, quizsDomain.ID, result.ID)
	})
	t.Run("Test 2 | err db", func(t *testing.T) {
		setUpUpdateQuiz()

		mysqlQuizsRepository.On("UpdateQuiz", mock.Anything, mock.Anything).Return(quizs.Domain{}, errors.New("db error")).Once()

		_, err := quizsService.UpdateQuiz(context.Background(), &quizsDomain)
		assert.NotNil(t, err)
	})
	t.Run("Test 3 | id quiz kososng", func(t *testing.T) {
		setUpUpdateQuiz()

		_, err := quizsService.UpdateQuiz(context.Background(), &quizs.Domain{ID: ""})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_QUIZ_EMPTY, err)
	})
}

func setUpDeleteQuiz() {
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, &modulsUsecase, &userCourseUsecase, time.Hour*1)

}

func TestQuizeUsecase_DeleteQuiz(t *testing.T) {
	t.Run("Test 1 | success delete quiz", func(t *testing.T) {
		setUpDeleteQuiz()
		mysqlQuizsRepository.On("DeleteQuiz", mock.Anything, mock.Anything).Return("7c1ec4be-8565-4b25-82cf-244d7730c398", nil).Once()
		result, err := quizsService.DeleteQuiz(context.Background(), "7c1ec4be-8565-4b25-82cf-244d7730c398")

		assert.Nil(t, err)
		assert.Equal(t, "7c1ec4be-8565-4b25-82cf-244d7730c398", result)
	})
	t.Run("Test 2 | id quiz kosong", func(t *testing.T) {
		setUpDeleteQuiz()
		_, err := quizsService.DeleteQuiz(context.Background(), "")

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_QUIZ_EMPTY, err)
	})
}

func setUpGetAllQuiz() {
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, &modulsUsecase, &userCourseUsecase, time.Hour*1)

	quizsDomain = quizs.Domain{
		ID:         "7c1ec4be-8565-4b25-82cf-244d7730c398",
		Pilihan:    []string{"a", "b", "c"},
		Pertanyaan: "Makan apa?",
		Jawaban:    "a",
		ModulId:    "36d8d8bc-87cb-467d-97c0-2902920457df",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
	listDomain = []quizs.Domain{quizsDomain, quizsDomain}
}

func TestQuizeUsecase_GetAllQuizModul(t *testing.T) {
	t.Run("test case 1 | success get all quiz from modul", func(t *testing.T) {
		setUpGetAllQuiz()
		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, nil).Once()

		_, err := quizsService.GetAllQuizModul(context.Background(), &quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df"})

		assert.Nil(t, err)
	})
	t.Run("test case 2 | handle err modul id empty", func(t *testing.T) {
		setUpGetAllQuiz()

		_, err := quizsService.GetAllQuizModul(context.Background(), &quizs.Domain{ModulId: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_MODUL_ID, err)
	})
	t.Run("test case 1 | success get all quiz from modul", func(t *testing.T) {
		setUpGetAllQuiz()
		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return([]quizs.Domain{}, errors.New("db error ni")).Once()

		_, err := quizsService.GetAllQuizModul(context.Background(), &quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df"})

		assert.NotNil(t, err)
	})
}

func setUpGetOneQuiz() {
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, &modulsUsecase, &userCourseUsecase, time.Hour*1)

	quizsDomain = quizs.Domain{
		ID:         "7c1ec4be-8565-4b25-82cf-244d7730c398",
		Pilihan:    []string{"a", "b", "c"},
		Pertanyaan: "Makan apa?",
		Jawaban:    "a",
		ModulId:    "36d8d8bc-87cb-467d-97c0-2902920457df",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}

func TestQuizeUsecase_GetOneQuiz(t *testing.T) {
	t.Run("Test case 1 | get one quiz success", func(t *testing.T) {
		setUpGetOneQuiz()
		mysqlQuizsRepository.On("GetOneQuiz", mock.Anything, mock.Anything).Return(quizsDomain, nil).Once()

		result, err := quizsService.GetOneQuiz(context.Background(), &quizsDomain)

		assert.Nil(t, err)
		assert.Equal(t, quizsDomain.ID, result.ID)
	})
	t.Run("Test case 2 | empty quiz id", func(t *testing.T) {
		setUpGetOneQuiz()

		_, err := quizsService.GetOneQuiz(context.Background(), &quizs.Domain{ID: ""})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_QUIZ_EMPTY, err)
	})
	t.Run("Test case 3 | get one quiz db error", func(t *testing.T) {
		setUpGetOneQuiz()
		mysqlQuizsRepository.On("GetOneQuiz", mock.Anything, mock.Anything).Return(quizsDomain, errors.New("eror guys")).Once()

		_, err := quizsService.GetOneQuiz(context.Background(), &quizsDomain)

		assert.NotNil(t, err)
	})
}

func setUpCalculateScoreQuiz() {
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, &modulsUsecase, &userCourseUsecase, time.Hour*1)

	quizsDomain = quizs.Domain{
		ID:      "7c1ec4be-8565-4b25-82cf-244d7730c398",
		Jawaban: "a",
		ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df",
	}

	var quizsDomain2 = quizs.Domain{
		ID:      "7c1ec4be-8565-4b25-82cf-244d7730c396",
		Jawaban: "b",
		ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df",
	}

	listDomain = []quizs.Domain{
		quizsDomain,
		quizsDomain2,
	}

	modulDomain = moduls.Domain{
		ID:       "36d8d8bc-87cb-467d-97c0-2902920457df",
		Title:    "ikan",
		Order:    1,
		CourseId: "123",
	}

	userCourseDomain = users_courses.Domain{
		ID:          "123",
		UserId:      "123",
		CourseId:    "123",
		Progres:     0,
		LastVideoId: "",
		LastModulId: "",
		Score:       0,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	scoreUserModulDomain = moduls.ScoreUserModul{
		ID:           "123",
		Nilai:        2,
		ModulID:      "36d8d8bc-87cb-467d-97c0-2902920457df",
		UserCourseId: userCourseDomain.ID,
	}
}

func TestQuizeUsecase_CalculateScoreQuiz(t *testing.T) {
	t.Run("Test case 1 | success calculate Score and if jawaban benar dapat score 2 dan jika tidak dijawab mendapat 0", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, nil).Once()
		modulsUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()
		modulsUsecase.On("CreateScoreModul", mock.Anything, mock.Anything).Return(scoreUserModulDomain, nil).Once()
		modulsUsecase.On("CalculateScoreCourse", mock.Anything, mock.Anything).Return(scoreUserModulDomain, nil).Once()
		userCourseUsecase.On("UpdateScoreCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()

		result, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: "7c1ec4be-8565-4b25-82cf-244d7730c398"},
		}, "123")

		assert.Nil(t, err)
		assert.Equal(t, quizsDomain.ModulId, result.ModulId)
		assert.Equal(t, 2, result.Skor)
	})
	t.Run("Test case 2 | success calculate Score and if jawaban benar dapat score 2 dan jika dijawab kosong mendapat 0", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, nil).Once()
		modulsUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()
		modulsUsecase.On("CreateScoreModul", mock.Anything, mock.Anything).Return(scoreUserModulDomain, nil).Once()
		modulsUsecase.On("CalculateScoreCourse", mock.Anything, mock.Anything).Return(scoreUserModulDomain, nil).Once()
		userCourseUsecase.On("UpdateScoreCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()

		result, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: "7c1ec4be-8565-4b25-82cf-244d7730c398"},
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.Nil(t, err)
		assert.Equal(t, quizsDomain.ModulId, result.ModulId)
		assert.Equal(t, 2, result.Skor)
	})
	t.Run("Test case 3 | success calculate Score and if jawaban benar dapat score 2 dan jika dijawab salah mendapat -1", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, nil).Once()
		modulsUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()
		modulsUsecase.On("CalculateScoreCourse", mock.Anything, mock.Anything).Return(scoreUserModulDomain, nil).Once()
		modulsUsecase.On("CreateScoreModul", mock.Anything, mock.Anything).Return(scoreUserModulDomain, nil).Once()
		userCourseUsecase.On("UpdateScoreCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()

		result, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: "7c1ec4be-8565-4b25-82cf-244d7730c398"},
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "c", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.Nil(t, err)
		assert.Equal(t, quizsDomain.ModulId, result.ModulId)
		assert.Equal(t, 1, result.Skor)
	})
	t.Run("Test case 4 | id quiz ditemukan kosong", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		_, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: ""},
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "c", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_QUIZ_EMPTY, err)
	})
	t.Run("Test case 5 | id modul ditemukan kosong", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		_, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "", Jawaban: "a", ID: "123"},
			quizs.Domain{ModulId: "", Jawaban: "c", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_MODUL_ID, err)
	})
	t.Run("Test case 6 | error db get all quiz", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, errors.New("db err")).Once()

		_, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: "7c1ec4be-8565-4b25-82cf-244d7730c398"},
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "c", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.NotNil(t, err)
	})
	t.Run("Test case 7 | error business get one modul", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, nil).Once()
		modulsUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomain, errors.New("usecase err")).Once()

		_, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: "7c1ec4be-8565-4b25-82cf-244d7730c398"},
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "c", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.NotNil(t, err)
	})
	t.Run("Test case 8 | error business get one modul", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, nil).Once()
		modulsUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, errors.New("usecase err")).Once()

		_, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: "7c1ec4be-8565-4b25-82cf-244d7730c398"},
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "c", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.NotNil(t, err)
	})

	t.Run("Test case 9 | error business CreateScoreModul", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, nil).Once()
		modulsUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()
		modulsUsecase.On("CreateScoreModul", mock.Anything, mock.Anything).Return(scoreUserModulDomain, errors.New("usecase err")).Once()

		_, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: "7c1ec4be-8565-4b25-82cf-244d7730c398"},
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "c", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.NotNil(t, err)
	})

	t.Run("Test case 10 | error business UpdateScoreCourse", func(t *testing.T) {
		setUpCalculateScoreQuiz()

		mysqlQuizsRepository.On("GetAllQuizModul", mock.Anything, mock.Anything).Return(listDomain, nil).Once()
		modulsUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomain, nil).Once()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()
		modulsUsecase.On("CreateScoreModul", mock.Anything, mock.Anything).Return(scoreUserModulDomain, nil).Once()
		modulsUsecase.On("CalculateScoreCourse", mock.Anything, mock.Anything).Return(scoreUserModulDomain, nil).Once()
		userCourseUsecase.On("UpdateScoreCourse", mock.Anything, mock.Anything).Return(userCourseDomain, errors.New("usecase err")).Once()

		_, err := quizsService.CalculateScoreQuiz(context.Background(), []quizs.Domain{
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "a", ID: "7c1ec4be-8565-4b25-82cf-244d7730c398"},
			quizs.Domain{ModulId: "36d8d8bc-87cb-467d-97c0-2902920457df", Jawaban: "c", ID: "7c1ec4be-8565-4b25-82cf-244d7730c396"},
		}, "123")

		assert.NotNil(t, err)
	})
}
