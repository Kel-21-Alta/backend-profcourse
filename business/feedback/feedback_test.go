package feedback_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/feedback"
	_mockFeedbackRepo "profcourse/business/feedback/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var mysqlFeedbackRespository _mockFeedbackRepo.Repository

var feedbackService feedback.Usecase
var feedbackDomain feedback.Domain

func setUpCreateFeedback() {
	feedbackService = feedback.NewFeedbackUsecase(&mysqlFeedbackRespository, time.Hour*1)
	feedbackDomain = feedback.Domain{
		ID:       "123",
		UserId:   "234",
		CourseId: "345",
		Review:   "qwer5ty",
		Rating:   4.5,
	}
}

func TestFeedbackUsecase_CreateFeedback(t *testing.T) {
	t.Run("test case 1 | success create feedback", func(t *testing.T) {
		setUpCreateFeedback()
		mysqlFeedbackRespository.On("CreateFeedback", mock.Anything, mock.Anything).Return(feedbackDomain, nil).Once()

		result, err := feedbackService.CreateFeedback(context.Background(), &feedback.Domain{
			UserId:   "234",
			CourseId: "345",
			Review:   "hai",
			Rating:   5.0,
		})

		assert.Nil(t, err)
		assert.Equal(t, feedbackDomain.ID, result.ID)
	})
	t.Run("test case 2 | error db create feedback", func(t *testing.T) {
		setUpCreateFeedback()
		mysqlFeedbackRespository.On("CreateFeedback", mock.Anything, mock.Anything).Return(feedbackDomain, errors.New("db err")).Once()

		_, err := feedbackService.CreateFeedback(context.Background(), &feedback.Domain{
			UserId:   "234",
			CourseId: "345",
			Review:   "hai",
			Rating:   5.0,
		})

		assert.NotNil(t, err)
	})
	t.Run("test case 3 | user empty | create feedback", func(t *testing.T) {
		setUpCreateFeedback()
		_, err := feedbackService.CreateFeedback(context.Background(), &feedback.Domain{
			UserId:   "",
			CourseId: "345",
			Review:   "hai",
			Rating:   5.0,
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})
	t.Run("test case 4 | course empty | create feedback", func(t *testing.T) {
		setUpCreateFeedback()
		_, err := feedbackService.CreateFeedback(context.Background(), &feedback.Domain{
			UserId:   "234",
			CourseId: "",
			Review:   "hai",
			Rating:   5.0,
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})

	t.Run("test case 5 | review empty | create feedback", func(t *testing.T) {
		setUpCreateFeedback()
		_, err := feedbackService.CreateFeedback(context.Background(), &feedback.Domain{
			UserId:   "234",
			CourseId: "345",
			Review:   "",
			Rating:   5.0,
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.REVIEW_EMPTY, err)
	})

	t.Run("test case 6 | rating empty | create feedback", func(t *testing.T) {
		setUpCreateFeedback()
		_, err := feedbackService.CreateFeedback(context.Background(), &feedback.Domain{
			UserId:   "234",
			CourseId: "345",
			Review:   "wqe",
			Rating:   0.0,
		})

		assert.NotNil(t, err)
		assert.Equal(t, controller.RATING_EMPTY, err)
	})
}
