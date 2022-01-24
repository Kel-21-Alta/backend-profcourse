package materies_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/materies"
	_mocksMateriesRepository "profcourse/business/materies/mocks"
	"profcourse/business/moduls"
	_mockModulUsecase "profcourse/business/moduls/mocks"
	"profcourse/business/users_courses"
	_mocksUsersCoursesUsecase "profcourse/business/users_courses/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var mysqlMateriesRepository _mocksMateriesRepository.Repository
var userCourseUsecase _mocksUsersCoursesUsecase.Usecase
var modulUsecase _mockModulUsecase.Usecase

var materiesService materies.Usecase
var materiesDomain materies.Domain
var allMaterisDomain materies.AllMateriModul
var userCourseDomain users_courses.Domain
var modulDomamin moduls.Domain

func setUpCreateMateri() {
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, &userCourseUsecase, &modulUsecase, time.Hour*1)
	materiesDomain = materies.Domain{
		ID:        "3ee0c5e0-ab38-4c4a-8c74-346ebcfa04e8",
		Title:     "Pengenalan Golang",
		ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
		Order:     1,
		Type:      2,
		UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func TestMateriesUsecase_CreateMateri(t *testing.T) {
	t.Run("Test case 1 | success create materies", func(t *testing.T) {
		setUpCreateMateri()
		mysqlMateriesRepository.On("CreateMateri", mock.Anything, mock.Anything).Return(materiesDomain, nil).Once()
		result, err := materiesService.CreateMateri(context.Background(), &materies.Domain{
			Title:     "Pengenalan Golang",
			ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
			Order:     1,
			Type:      2,
			UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		})
		assert.Nil(t, err)
		assert.Equal(t, materiesDomain.ModulId, result.ModulId)
	})
}

func setUpDeleteMateri() {
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, &userCourseUsecase, &modulUsecase, time.Hour*1)
	materiesDomain = materies.Domain{
		ID: "3ee0c5e0-ab38-4c4a-8c74-346ebcfa04e8",
	}
}

func TestMateriesUsecase_DeleteMateri(t *testing.T) {
	t.Run("Test case 1 | success delete materi", func(t *testing.T) {
		setUpDeleteMateri()
		mysqlMateriesRepository.On("DeleteMateri", mock.Anything, mock.Anything).Return(materiesDomain, nil).Once()
		result, err := materiesService.DeleteMateri(context.Background(), &materiesDomain)
		assert.Nil(t, err)
		assert.Equal(t, materiesDomain.ID, result.ID)
	})
	t.Run("Test case 2 | handle id materi empty", func(t *testing.T) {
		setUpDeleteMateri()
		_, err := materiesService.DeleteMateri(context.Background(), &materies.Domain{ID: ""})
		assert.Equal(t, controller.ID_MATERI_EMPTY, err)
	})
}

func TestMateriesUsecase_ValidasiMateri(t *testing.T) {
	t.Run("Test case 1 | success create materies", func(t *testing.T) {
		setUpCreateMateri()
		result, err := materiesService.ValidasiMateri(context.Background(), &materies.Domain{
			Title:     "Pengenalan Golang",
			ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
			Order:     1,
			Type:      2,
			UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		})
		assert.Nil(t, err)
		assert.Equal(t, materiesDomain.ModulId, result.ModulId)
	})
	t.Run("Test case modul 2 | handle error modul id empty", func(t *testing.T) {
		setUpCreateMateri()
		_, err := materiesService.ValidasiMateri(context.Background(), &materies.Domain{
			Title:     "Pengenalan Golang",
			ModulId:   "",
			Order:     1,
			Type:      2,
			UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		})
		assert.Equal(t, controller.EMPTY_MODUL_ID, err)
	})
	t.Run("Testcase 3 | handle error title materi empty", func(t *testing.T) {
		setUpCreateMateri()
		_, err := materiesService.ValidasiMateri(context.Background(), &materies.Domain{
			Title:     "",
			ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
			Order:     1,
			Type:      2,
			UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		})
		assert.Equal(t, controller.TITLE_EMPTY, err)
	})
	t.Run("Testcase 4 | handle error materi empty", func(t *testing.T) {
		setUpCreateMateri()
		_, err := materiesService.ValidasiMateri(context.Background(), &materies.Domain{
			Title:     "Pengenalan GOlang",
			ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
			Order:     1,
			Type:      2,
			UrlMateri: "",
		})
		assert.Equal(t, controller.EMPTY_FILE_MATERI, err)
	})
	t.Run("Testcase 5 | handle error order/urutan materi empty", func(t *testing.T) {
		setUpCreateMateri()
		_, err := materiesService.ValidasiMateri(context.Background(), &materies.Domain{
			Title:     "Pengenalan",
			ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
			Order:     0,
			Type:      2,
			UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		})
		assert.Equal(t, controller.ORDER_MATERI_EMPTY, err)
	})
	t.Run("Testcase 6 | handle error type materi empty", func(t *testing.T) {
		setUpCreateMateri()
		_, err := materiesService.ValidasiMateri(context.Background(), &materies.Domain{
			Title:     "Pengenalan",
			ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
			Order:     1,
			Type:      0,
			UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		})
		assert.Equal(t, controller.TYPE_MATERI_EMPTY, err)
	})
	t.Run("Testcase 7 | handle error type materi tidak valid", func(t *testing.T) {
		setUpCreateMateri()
		_, err := materiesService.ValidasiMateri(context.Background(), &materies.Domain{
			Title:     "Pengenalan",
			ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
			Order:     1,
			Type:      3,
			UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		})
		assert.Equal(t, controller.TYPE_MATERI_WRONG, err)
	})
}

func setUpUpdateMateri() {
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, &userCourseUsecase, &modulUsecase, time.Hour*1)
	materiesDomain = materies.Domain{
		ID:        "3ee0c5e0-ab38-4c4a-8c74-346ebcfa04e8",
		Title:     "Pengenalan Golang",
		ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
		Order:     1,
		Type:      2,
		UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func TestMateriesUsecase_UpdateMateri(t *testing.T) {
	t.Run("Test case 1 | success update materi", func(t *testing.T) {
		setUpUpdateMateri()
		mysqlMateriesRepository.On("UpdateMateri", mock.Anything, mock.Anything).Return(materiesDomain, nil).Once()
		result, err := materiesService.UpdateMateri(context.Background(), &materiesDomain)
		assert.Nil(t, err)
		assert.Equal(t, materiesDomain.ID, result.ID)
	})
	t.Run("Test case 2 | handle id materi empty", func(t *testing.T) {
		setUpUpdateMateri()
		_, err := materiesService.UpdateMateri(context.Background(), &materies.Domain{ID: ""})
		assert.Equal(t, controller.ID_MATERI_EMPTY, err)
	})
	t.Run("Test case 3 | handle error db", func(t *testing.T) {
		setUpUpdateMateri()
		mysqlMateriesRepository.On("UpdateMateri", mock.Anything, mock.Anything).Return(materies.Domain{}, errors.New("db error")).Once()
		_, err := materiesService.UpdateMateri(context.Background(), &materiesDomain)
		assert.NotNil(t, err)
	})
}

func setUpGetOneMateri() {
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, &userCourseUsecase, &modulUsecase, time.Hour*1)
	materiesDomain = materies.Domain{
		ID:        "3ee0c5e0-ab38-4c4a-8c74-346ebcfa04e8",
		Title:     "Pengenalan Golang",
		ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
		Order:     1,
		Type:      2,
		UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		User: materies.CurrentUser{
			ID:          "c56780b2-dee3-45c6-9bb0-2496f7a13b94",
			CurrentTime: "",
			IsComplate:  false,
		},
	}
}

func TestMateriesUsecase_GetOneMateri(t *testing.T) {
	t.Run("Testcase 1 | success get one materi", func(t *testing.T) {
		setUpGetOneMateri()
		mysqlMateriesRepository.On("GetOnemateri", mock.Anything, mock.Anything).Return(materiesDomain, nil).Once()
		result, err := materiesService.GetOneMateri(context.Background(), &materiesDomain)
		assert.Nil(t, err)
		assert.Equal(t, materiesDomain.ID, result.ID)
	})
	t.Run("Testcase 2 | handle id materi empty", func(t *testing.T) {
		setUpGetOneMateri()
		_, err := materiesService.GetOneMateri(context.Background(), &materies.Domain{ID: ""})
		assert.Equal(t, controller.ID_MATERI_EMPTY, err)
	})
	t.Run("Testcase 3 | handle err db", func(t *testing.T) {
		setUpGetOneMateri()
		mysqlMateriesRepository.On("GetOnemateri", mock.Anything, mock.Anything).Return(materies.Domain{}, errors.New("error db")).Once()
		_, err := materiesService.GetOneMateri(context.Background(), &materiesDomain)
		assert.NotNil(t, err)
	})
}

func setUpGetAllMateri() {
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, &userCourseUsecase, &modulUsecase, time.Hour*1)
	allMaterisDomain = materies.AllMateriModul{
		JawabanMateri: 2,
		Materi:        []materies.Domain{materiesDomain, materiesDomain},
	}
	modulDomamin = moduls.Domain{
		ID:       "123",
		CourseId: "321",
	}
	userCourseDomain = users_courses.Domain{
		ID:       "123",
		UserId:   "312",
		CourseId: "321",
	}
}

func TestMateriesUsecase_GetAllMateri(t *testing.T) {
	t.Run("Test case 1 | success get all materi", func(t *testing.T) {
		setUpGetAllMateri()
		mysqlMateriesRepository.On("GetAllMateri", mock.Anything, mock.Anything).Return(allMaterisDomain, nil).Once()

		modulUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomamin, nil).Once()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()

		result, err := materiesService.GetAllMateri(context.Background(), &materies.Domain{ModulId: "d0b4fac1-09bf-4455-b3ec-74e5b54d2c7f", User: materies.CurrentUser{ID: "727c0932-1c4b-497a-af40-f373d519d242"}})

		assert.Nil(t, err)
		assert.Equal(t, allMaterisDomain.JawabanMateri, result.JawabanMateri)
	})
	t.Run("Test case 2 | handle modul id empty", func(t *testing.T) {
		setUpGetAllMateri()

		_, err := materiesService.GetAllMateri(context.Background(), &materies.Domain{ModulId: "", User: materies.CurrentUser{ID: "727c0932-1c4b-497a-af40-f373d519d242"}})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_MODUL_ID, err)
	})

	t.Run("Test case 3 | handle id user kosong", func(t *testing.T) {
		setUpGetAllMateri()

		_, err := materiesService.GetAllMateri(context.Background(), &materies.Domain{ModulId: "d0b4fac1-09bf-4455-b3ec-74e5b54d2c7f", User: materies.CurrentUser{ID: ""}})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})

	t.Run("Test case 4 | handle error db", func(t *testing.T) {
		setUpGetAllMateri()
		mysqlMateriesRepository.On("GetAllMateri", mock.Anything, mock.Anything).Return(materies.AllMateriModul{}, errors.New("db err ni")).Once()

		modulUsecase.On("GetOneModul", mock.Anything, mock.Anything).Return(modulDomamin, nil).Once()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()

		_, err := materiesService.GetAllMateri(context.Background(), &materies.Domain{ModulId: "d0b4fac1-09bf-4455-b3ec-74e5b54d2c7f", User: materies.CurrentUser{ID: "727c0932-1c4b-497a-af40-f373d519d242"}})

		assert.NotNil(t, err)
	})
}

func setUpUpdateProgressMateri() {
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, &userCourseUsecase, &modulUsecase, time.Hour*1)
	userCourseDomain = users_courses.Domain{
		ID:          "",
		UserId:      "",
		CourseId:    "",
		Progres:     0,
		LastVideoId: "",
		LastModulId: "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	materiesDomain = materies.Domain{
		ID:         "",
		Title:      "",
		ModulId:    "",
		Order:      0,
		Type:       0,
		TypeString: "",
		UrlMateri:  "",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		User:       materies.CurrentUser{},
		UserCourse: materies.UserCourse{},
	}
}

func TestMateriesUsecase_UpdateProgressMateri(t *testing.T) {
	t.Run("Test case 1 | success update progress", func(t *testing.T) {
		setUpUpdateProgressMateri()
		userCourseUsecase.On("GetOneUserCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()
		mysqlMateriesRepository.On("UpdateProgressMateri", mock.Anything, mock.Anything).Return(materiesDomain, nil).Once()
		mysqlMateriesRepository.On("GetCountMateriFinish", mock.Anything, mock.Anything).Return(10, nil).Once()
		mysqlMateriesRepository.On("GetCountMateriCourse", mock.Anything, mock.Anything).Return(20, nil).Once()
		userCourseUsecase.On("UpdateProgressCourse", mock.Anything, mock.Anything).Return(userCourseDomain, nil).Once()

		_, err := materiesService.UpdateProgressMateri(context.Background(), &materies.Domain{ID: "c56780b2-dee3-45c6-9bb0-2496f7a13b94", User: materies.CurrentUser{ID: "123", CourseId: "123"}, UserCourse: materies.UserCourse{UserCourseId: "123"}})

		assert.Nil(t, err)
	})
	t.Run("Test case 2 | error empty id materi", func(t *testing.T) {
		setUpUpdateProgressMateri()

		_, err := materiesService.UpdateProgressMateri(context.Background(), &materies.Domain{ID: "", User: materies.CurrentUser{ID: "123", CourseId: "123"}, UserCourse: materies.UserCourse{UserCourseId: "123"}})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_MATERI_EMPTY, err)
	})

	t.Run("Test case 3 | error empty id user", func(t *testing.T) {
		setUpUpdateProgressMateri()

		_, err := materiesService.UpdateProgressMateri(context.Background(), &materies.Domain{ID: "123", User: materies.CurrentUser{ID: "", CourseId: "123"}, UserCourse: materies.UserCourse{UserCourseId: "123"}})

		assert.NotNil(t, err)
		assert.Equal(t, controller.ID_EMPTY, err)
	})

	t.Run("Test case 4 | error empty id course", func(t *testing.T) {
		setUpUpdateProgressMateri()

		_, err := materiesService.UpdateProgressMateri(context.Background(), &materies.Domain{ID: "123", User: materies.CurrentUser{ID: "123", CourseId: ""}, UserCourse: materies.UserCourse{UserCourseId: "123"}})

		assert.NotNil(t, err)
		assert.Equal(t, controller.EMPTY_COURSE, err)
	})
}
