package users_spesializations_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/spesializations"
	_mockSpesialization "profcourse/business/spesializations/mocks"
	"profcourse/business/users_courses"
	_mocksUsersCoursesUsecase "profcourse/business/users_courses/mocks"
	"profcourse/business/users_spesializations"
	_mysqlUserRepository "profcourse/business/users_spesializations/mocks"
	"testing"
	"time"
)

var mysqlUserSpesializationRepository _mysqlUserRepository.Repository
var usecaseSpesialization _mockSpesialization.Usecase
var usecaseUserCourse _mocksUsersCoursesUsecase.Usecase

var userSpesializationService users_spesializations.Usecase
var userSpesializationDomain users_spesializations.Domain
var spesializationDomain spesializations.Domain

func setRegisterSpesialization() {
	userSpesializationService = users_spesializations.NewUsersSpesializationsUsecase(&mysqlUserSpesializationRepository, &usecaseSpesialization, &usecaseUserCourse, time.Hour*1)
	userSpesializationDomain = users_spesializations.Domain{
		ID:               "123",
		UserID:           "234",
		SpesializationID: "345",
		Progress:         0,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}
	spesializationDomain = spesializations.Domain{
		ID:            "",
		Title:         "",
		ImageUrl:      "",
		Description:   "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		CertificateId: "",
		Courses: []spesializations.Course{
			spesializations.Course{
				ID:          "123",
				Title:       "qwe",
				Rating:      0,
				Description: "qwe",
			},
		},
	}
}

func TestUsersSpesializationsUsecase_RegisterSpesialization(t *testing.T) {
	t.Run("Test case 1 | success register", func(t *testing.T) {
		setRegisterSpesialization()
		mysqlUserSpesializationRepository.On("RegisterSpesialization", mock.Anything, mock.Anything).Return(userSpesializationDomain, nil).Once()
		mysqlUserSpesializationRepository.On("GetEndRollSpesializationById", mock.Anything, mock.Anything).Return(users_spesializations.Domain{}, nil).Once()
		usecaseSpesialization.On("GetOneSpesialization", mock.Anything, mock.Anything).Return(spesializationDomain, nil).Once()
		usecaseUserCourse.On("UserRegisterCourse", mock.Anything, mock.Anything).Return(&users_courses.Domain{}, nil).Once()

		_, err := userSpesializationService.RegisterSpesialization(context.Background(), &users_spesializations.Domain{UserID: "123", SpesializationID: "234"})

		assert.Nil(t, err)
	})
	t.Run("Test case 2 | spesialization id empty", func(t *testing.T) {
		setRegisterSpesialization()
		_, err := userSpesializationService.RegisterSpesialization(context.Background(), &users_spesializations.Domain{UserID: "123", SpesializationID: ""})

		assert.NotNil(t, err)
	})
	t.Run("Test case 3 | user id empty", func(t *testing.T) {
		setRegisterSpesialization()
		_, err := userSpesializationService.RegisterSpesialization(context.Background(), &users_spesializations.Domain{UserID: "", SpesializationID: "234"})

		assert.NotNil(t, err)
	})
	t.Run("Test case 1 | success register", func(t *testing.T) {
		setRegisterSpesialization()
		mysqlUserSpesializationRepository.On("GetEndRollSpesializationById", mock.Anything, mock.Anything).Return(users_spesializations.Domain{ID: "123"}, nil).Once()

		_, err := userSpesializationService.RegisterSpesialization(context.Background(), &users_spesializations.Domain{UserID: "123", SpesializationID: "234"})

		assert.NotNil(t, err)
	})
}
