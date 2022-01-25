package spesializations_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/spesializations"
	_mockSpesialization "profcourse/business/spesializations/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var mysqlSpesializationRepository _mockSpesialization.Repository

var spesializationService spesializations.Usecase
var spesializationDomain spesializations.Domain
var listSpesializationDomain []spesializations.Domain
var spesializationCourse spesializations.Course
var summary spesializations.Summary

func setUpGetOneSpesialization() {
	spesializationService = spesializations.NewSpesializationUsecase(&mysqlSpesializationRepository, time.Hour*1)
	spesializationCourse = spesializations.Course{
		ID:          "dasdas",
		Title:       "dasdas",
		Rating:      0,
		Description: "dasdasdas",
	}
	spesializationDomain = spesializations.Domain{
		ID:            "123",
		Title:         "Mastering Back End",
		ImageUrl:      "https://placeimg.com/640/480/any",
		Description:   "Manya isinya",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		CertificateId: "",
		Courses:       []spesializations.Course{spesializationCourse, spesializationCourse},
	}
}

func TestSpesializationUsecase_GetOneSpesialization(t *testing.T) {
	t.Run("Test case 1 | Success get one spesailization", func(t *testing.T) {
		setUpGetOneSpesialization()
		mysqlSpesializationRepository.On("GetOneSpesialization", mock.Anything, mock.Anything).Return(spesializationDomain, nil).Once()
		result, err := spesializationService.GetOneSpesialization(context.Background(), &spesializations.Domain{ID: "123"})
		assert.Nil(t, err)
		assert.Equal(t, spesializationDomain.Title, result.Title)
	})
	t.Run("Test case 1 | Success get one spesailization", func(t *testing.T) {
		setUpGetOneSpesialization()
		_, err := spesializationService.GetOneSpesialization(context.Background(), &spesializations.Domain{ID: ""})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
}

func setUpCreateSpesialization() {
	spesializationService = spesializations.NewSpesializationUsecase(&mysqlSpesializationRepository, time.Hour*1)
	spesializationDomain = spesializations.Domain{
		ID:            "123",
		Title:         "Mastering Back End",
		ImageUrl:      "https://placeimg.com/640/480/any",
		Description:   "Manya isinya",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		CertificateId: "",
		CourseIds:     []string{"234", "235"},
	}
}

func TestSpesializationUsecase_CreateSpasialization(t *testing.T) {
	t.Run("Test 0 | hendle maker spesialization must be admin", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole: 2})
		assert.NotNil(t, err)
		assert.Error(t, controller.TITLE_EMPTY, err)
	})
	t.Run("Test 1 | hendle name empty", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole: 1, Title: ""})
		assert.NotNil(t, err)
		assert.Error(t, controller.TITLE_EMPTY, err)
	})
	t.Run("Test 2 | hendle desc empty", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole: 1, Title: "Beck end", Description: ""})
		assert.NotNil(t, err)
		assert.Error(t, controller.DESC_EMPTY, err)
	})
	t.Run("Test 3 | hendle desc empty", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole: 1, Title: "Beck end", Description: "bla bla", ImageUrl: ""})
		assert.NotNil(t, err)
		assert.Error(t, controller.IMAGE_EMPTY, err)
	})
	t.Run("Test 4 | hendle courses empty", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole: 1, Title: "Beck end", Description: "bla bla", ImageUrl: "http://kakdfgk.com", CourseIds: nil})
		assert.NotNil(t, err)
		assert.Error(t, controller.COURSES_SPESIALIZATION_EMPTY, err)
	})

	t.Run("Test 4 | Success create spesialization", func(t *testing.T) {
		setUpCreateSpesialization()
		mysqlSpesializationRepository.On("CreateSpasialization", mock.Anything, mock.Anything).Return(spesializationDomain, nil).Once()
		result, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{
			Title:       "Mastering Back End",
			ImageUrl:    "https://placeimg.com/640/480/any",
			Description: "Manya isinya",
			CourseIds:   []string{"ijadfjkg"},
			MakerRole:   1,
		})
		assert.Nil(t, err)
		assert.Equal(t, "Mastering Back End", result.Title)
		assert.Equal(t, spesializationDomain.ImageUrl, result.ImageUrl)
		assert.Equal(t, spesializationDomain.Description, result.Description)
	})
	t.Run("test 5 | Handle error db", func(t *testing.T) {
		setUpCreateSpesialization()
		mysqlSpesializationRepository.On("CreateSpasialization", mock.Anything, mock.Anything).Return(spesializations.Domain{}, errors.New("DB")).Once()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{
			Title:       "Mastering Back End",
			ImageUrl:    "https://placeimg.com/640/480/any",
			Description: "Manya isinya",
			CourseIds:   []string{"ijadfjkg"},
			MakerRole:   1,
		})
		assert.NotNil(t, err)
	})
}

func setUpGetAllSpesialization() {
	spesializationService = spesializations.NewSpesializationUsecase(&mysqlSpesializationRepository, time.Hour*1)
	spesializationDomain = spesializations.Domain{
		ID:            "123",
		Title:         "Mastering Back End",
		ImageUrl:      "https://placeimg.com/640/480/any",
		Description:   "Manya isinya",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		CertificateId: "",
		CourseIds:     []string{"234", "235"},
	}
	listSpesializationDomain = []spesializations.Domain{spesializationDomain, spesializationDomain}
}

func TestSpesializationUsecase_GetAllSpesializations(t *testing.T) {
	t.Run("Test case 1 | Success mendapatkan data", func(t *testing.T) {
		setUpGetAllSpesialization()
		mysqlSpesializationRepository.On("GetAllSpesializations", mock.Anything, mock.Anything).Return(listSpesializationDomain, nil).Once()
		result, err := spesializationService.GetAllSpesializations(context.Background(), &spesializations.Domain{})
		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})
}

func setUpGetCountSpesialization() {
	spesializationService = spesializations.NewSpesializationUsecase(&mysqlSpesializationRepository, time.Hour*1)
	summary = spesializations.Summary{CountSpesialization: 2}
}


func TestSpesializationUsecase_GetCountSpesializations(t *testing.T) {
	t.Run("test case 1 | success mendapatkan data", func(t *testing.T) {
		setUpGetCountSpesialization()
		mysqlSpesializationRepository.On("GetCountSpesializations", mock.Anything).Return(summary, nil).Once()
		result, err := spesializationService.GetCountSpesializations(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, summary.CountSpesialization, result.CountSpesialization)
	})
	t.Run("test case 2 | err db", func(t *testing.T) {
		setUpGetCountSpesialization()
		mysqlSpesializationRepository.On("GetCountSpesializations", mock.Anything).Return(summary, errors.New("err")).Once()
		_, err := spesializationService.GetCountSpesializations(context.Background())
		assert.NotNil(t, err)
	})
}
