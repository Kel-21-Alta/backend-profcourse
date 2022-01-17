package quizs_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/quizs"
	_mocksQuizRepository "profcourse/business/quizs/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var mysqlQuizsRepository _mocksQuizRepository.Repository

var quizsService quizs.Usecase
var quizsDomain quizs.Domain

func setUpCreateQuizs() {
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, time.Hour*1)
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
	t.Run("Test case 1 | success create quiz", func(t *testing.T) {
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
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, time.Hour*1)
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
}

func setUpDeleteQuiz() {
	quizsService = quizs.NewQuizUsecase(&mysqlQuizsRepository, time.Hour*1)
}

func TestQuizeUsecase_DeleteQuiz(t *testing.T) {
	t.Run("Test 1 | success delete quiz", func(t *testing.T) {
		setUpDeleteQuiz()
		mysqlQuizsRepository.On("DeleteQuiz", mock.Anything, mock.Anything).Return("7c1ec4be-8565-4b25-82cf-244d7730c398", nil).Once()

	})
}
