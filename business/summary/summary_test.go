package summary_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/courses"
	_mocksCourseUsecase "profcourse/business/courses/mocks"
	"profcourse/business/summary"
	"profcourse/business/users"
	_mockUserUsecase "profcourse/business/users/mocks"
	"testing"
	"time"
)

var usersUsecase _mockUserUsecase.Usecase
var courseUsecase _mocksCourseUsecase.Usecase

var summaryDomain summary.Domain
var summaryService summary.Usecase
var courseSummary courses.Summary
var userSummary users.Summary

func setUpGetAllSummary() {
	summaryService = summary.NewSummaryUsecase(time.Hour*1, &courseUsecase, &usersUsecase)
	summaryDomain = summary.Domain{
		CountCourse:         8,
		CountUser:           9,
		CountSpesialization: 0,
	}
	courseSummary = courses.Summary{CountCourse: 8}
	userSummary = users.Summary{CountUser: 9}
}

func TestSummaryUsecase_GetAllSummary(t *testing.T) {
	t.Run("Test case 1 | success dapat data summary course", func(t *testing.T) {
		setUpGetAllSummary()
		usersUsecase.On("GetCountUser", mock.Anything).Return(&userSummary, nil).Once()
		courseUsecase.On("GetCountCourse", mock.Anything).Return(&courseSummary, nil).Once()

		allSummary, err := summaryService.GetAllSummary(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, summaryDomain.CountUser, allSummary.CountUser)
	})
	t.Run("Test case 2 | ", func(t *testing.T) {
		setUpGetAllSummary()
		usersUsecase.On("GetCountUser", mock.Anything).Return(&userSummary, nil).Once()
		courseUsecase.On("GetCountCourse", mock.Anything).Return(&courseSummary, nil).Once()

		allSummary, err := summaryService.GetAllSummary(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, summaryDomain.CountUser, allSummary.CountUser)
	})
}
