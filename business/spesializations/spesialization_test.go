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

func setUpCreateSpesialization() {
	spesializationService = spesializations.NewSpesializationUsecase(&mysqlSpesializationRepository, time.Hour*1)
	spesializationDomain = spesializations.Domain{
		ID:            "123",
		Name:          "Mastering Back End",
		ImageUrl:      "https://placeimg.com/640/480/any",
		Description:   "Manya isinya",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		CertificateId: "",
		Courses:       []string{"234", "235"},
	}
}

func TestSpesializationUsecase_CreateSpasialization(t *testing.T) {
	t.Run("Test 0 | hendle maker spesialization must be admin", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole:2,})
		assert.NotNil(t, err)
		assert.Error(t, controller.TITLE_EMPTY, err)
	})
	t.Run("Test 1 | hendle name empty", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole:1,Name: ""})
		assert.NotNil(t, err)
		assert.Error(t, controller.TITLE_EMPTY, err)
	})
	t.Run("Test 2 | hendle desc empty", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole:1,Name: "Beck end", Description: ""})
		assert.NotNil(t, err)
		assert.Error(t, controller.DESC_EMPTY, err)
	})
	t.Run("Test 3 | hendle desc empty", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole:1,Name: "Beck end", Description: "bla bla", ImageUrl: ""})
		assert.NotNil(t, err)
		assert.Error(t, controller.IMAGE_EMPTY, err)
	})
	t.Run("Test 4 | hendle courses empty", func(t *testing.T) {
		setUpCreateSpesialization()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{MakerRole:1,Name: "Beck end", Description: "bla bla", ImageUrl: "http://kakdfgk.com", Courses: nil})
		assert.NotNil(t, err)
		assert.Error(t, controller.COURSES_SPESIALIZATION_EMPTY, err)
	})

	t.Run("Test 4 | Success create spesialization", func(t *testing.T) {
		setUpCreateSpesialization()
		mysqlSpesializationRepository.On("CreateSpasialization", mock.Anything, mock.Anything).Return(&spesializationDomain, nil).Once()
		result, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{
			Name:        "Mastering Back End",
			ImageUrl:    "https://placeimg.com/640/480/any",
			Description: "Manya isinya",
			Courses:     []string{"ijadfjkg"},
			MakerRole:1,
		})
		assert.Nil(t, err)
		assert.Equal(t, "Mastering Back End", result.Name)
		assert.Equal(t, spesializationDomain.ImageUrl, result.ImageUrl)
		assert.Equal(t, spesializationDomain.Description, result.Description)
	})
	t.Run("test 5 | Handle error db", func(t *testing.T) {
		setUpCreateSpesialization()
		mysqlSpesializationRepository.On("CreateSpasialization", mock.Anything, mock.Anything).Return(&spesializations.Domain{}, errors.New("DB")).Once()
		_, err := spesializationService.CreateSpasialization(context.Background(), &spesializations.Domain{
			Name:        "Mastering Back End",
			ImageUrl:    "https://placeimg.com/640/480/any",
			Description: "Manya isinya",
			Courses:     []string{"ijadfjkg"},
			MakerRole:1,
		})
		assert.NotNil(t, err)
	})
}
