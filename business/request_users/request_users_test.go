package request_users_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/request_users"
	_mockRequestuser "profcourse/business/request_users/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var mysqlRequestUser _mockRequestuser.Repository

var requestUserService request_users.Usecase
var requestUserDoamin request_users.Domain
var listCategory []request_users.Category
var listRequest []request_users.Domain

func setupCreateRequestuser() {
	requestUserService = request_users.NewRequestUserUsecase(&mysqlRequestUser, time.Hour*1)

	requestUserDoamin = request_users.Domain{
		Id:         "123",
		UserId:     "234",
		CategoryID: "345",
		Request:    "Hai",
		Category: request_users.Category{
			ID:        "234",
			Title:     "Online Course",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func TestRequestUserUsecase_CreateRequest(t *testing.T) {
	t.Run("test case 1 | success create request", func(t *testing.T) {
		setupCreateRequestuser()

		mysqlRequestUser.On("CreateRequest", mock.Anything, mock.Anything).Return(requestUserDoamin, nil).Once()
		mysqlRequestUser.On("GetOneRequest", mock.Anything, mock.Anything).Return(requestUserDoamin, nil).Once()

		result, err := requestUserService.CreateRequest(context.Background(), &request_users.Domain{
			CategoryID: "345",
			UserId:     "234",
			Request:    "Hai",
		})
		assert.Nil(t, err)
		assert.Equal(t, requestUserDoamin.Request, result.Request)
	})
	t.Run("test case 2 | db error create request", func(t *testing.T) {
		setupCreateRequestuser()

		mysqlRequestUser.On("CreateRequest", mock.Anything, mock.Anything).Return(requestUserDoamin, errors.New("err")).Once()

		_, err := requestUserService.CreateRequest(context.Background(), &request_users.Domain{
			CategoryID: "345",
			UserId:     "234",
			Request:    "Hai",
		})
		assert.NotNil(t, err)
	})
	t.Run("test case 3 | db error get one request", func(t *testing.T) {
		setupCreateRequestuser()

		mysqlRequestUser.On("CreateRequest", mock.Anything, mock.Anything).Return(requestUserDoamin, nil).Once()
		mysqlRequestUser.On("GetOneRequest", mock.Anything, mock.Anything).Return(requestUserDoamin, errors.New("err")).Once()

		_, err := requestUserService.CreateRequest(context.Background(), &request_users.Domain{
			CategoryID: "345",
			UserId:     "234",
			Request:    "Hai",
		})
		assert.NotNil(t, err)
	})
	t.Run("test case 4 | category empty create request", func(t *testing.T) {
		setupCreateRequestuser()

		_, err := requestUserService.CreateRequest(context.Background(), &request_users.Domain{
			CategoryID: "",
			UserId:     "234",
			Request:    "Hai",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.CATEGORY_EMPTY, err)
	})

	t.Run("test case 5 | user id empty create request", func(t *testing.T) {
		setupCreateRequestuser()

		_, err := requestUserService.CreateRequest(context.Background(), &request_users.Domain{
			CategoryID: "234",
			UserId:     "",
			Request:    "Hai",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})

	t.Run("test case 6 | request empty create request", func(t *testing.T) {
		setupCreateRequestuser()

		_, err := requestUserService.CreateRequest(context.Background(), &request_users.Domain{
			CategoryID: "234",
			UserId:     "23",
			Request:    "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controller.REQUEST_EMPTY, err)
	})
}

func setupGetAllCategory() {
	requestUserService = request_users.NewRequestUserUsecase(&mysqlRequestUser, time.Hour*1)
	listCategory = []request_users.Category{
		request_users.Category{
			ID:        "123",
			Title:     "123",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
}

func TestRequestUserUsecase_GetAllCategoryRequest(t *testing.T) {
	t.Run("Test case 1 | success", func(t *testing.T) {
		setupGetAllCategory()

		mysqlRequestUser.On("GetAllCategoryRequest", mock.Anything, mock.Anything).Return(listCategory, nil).Once()

		_, err := requestUserService.GetAllCategoryRequest(context.Background())

		assert.Nil(t, err)
	})
	t.Run("Test case 1 | error db", func(t *testing.T) {
		setupGetAllCategory()

		mysqlRequestUser.On("GetAllCategoryRequest", mock.Anything, mock.Anything).Return(listCategory, errors.New("db")).Once()

		_, err := requestUserService.GetAllCategoryRequest(context.Background())

		assert.NotNil(t, err)
	})
}

func setupGetAllRequestUser() {
	requestUserService = request_users.NewRequestUserUsecase(&mysqlRequestUser, time.Hour*1)

	requestUserDoamin = request_users.Domain{
		Id:         "123",
		UserId:     "234",
		CategoryID: "345",
		Request:    "Hai",
		Category: request_users.Category{
			ID:        "234",
			Title:     "Online Course",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	listRequest = []request_users.Domain{requestUserDoamin}
}

func TestRequestUserUsecase_GetAllRequestUser(t *testing.T) {
	t.Run("Test case 1 | success", func(t *testing.T) {
		
	})
}
