package summary_test

import (
	"profcourse/business/courses"
	_mocksCourseUsecase "profcourse/business/courses/mocks"
	"profcourse/business/summary"
	"profcourse/business/users"
	_mockUserUsecase "profcourse/business/users/mocks"
	"testing"
)

var usersUsecase _mockUserUsecase.Usecase
var courseUsecase _mocksCourseUsecase.Usecase

var summaryDomain summary.Domain
var summaryService summary.Usecase
var courseSummary courses.Summary
var userSummary users.Summary

func setUpGetAllSummary() {
	//summaryService = summary.NewSummaryUsecase(time.Hour*1, &usersUsecase, &courseUsecase)
	summaryDomain = summary.Domain{
		CountCourse:         8,
		CountUser:           9,
		CountSpesialization: 0,
	}
	courseSummary = courses.Summary{CountCourse: 8}
	userSummary = users.Summary{CountUser: 9}
}

func TestSummaryUsecase_GetAllSummary(t *testing.T) {
	//t.Run("Test case 1 | success dapat data summary course", func(t *testing.T) {
	//	setUpGetAllSummary()
	//	usersUsecase.On("GetCountUser", mock.Anything).Return(&courseSummary, nil).Once()
	//	courseUsecase.On("GetCountCourse", mock.Anything).Return(&userSummary, nil).Once()
	//
	//	allSummary, err := summaryService.GetAllSummary(context.Background())
	//	assert.Nil(t, err)
	//	assert.Equal(t, summaryDomain.CountUser, allSummary.CountUser)
	//})
}
