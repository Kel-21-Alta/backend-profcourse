package materies_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"profcourse/business/materies"
	_mocksMateriesRepository "profcourse/business/materies/mocks"
	controller "profcourse/controllers"
	"testing"
	"time"
)

var mysqlMateriesRepository _mocksMateriesRepository.Repository

var materiesService materies.Usecase
var materiesDomain materies.Domain

func setUpCreateMateri() {
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, time.Hour*1)
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
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, time.Hour*1)
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
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, time.Hour*1)
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
	materiesService = materies.NewMateriesUsecase(&mysqlMateriesRepository, time.Hour*1)
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
	t.Run("Testcase 1 | handle err db", func(t *testing.T) {
		setUpGetOneMateri()
		mysqlMateriesRepository.On("GetOnemateri", mock.Anything, mock.Anything).Return(materies.Domain{}, errors.New("error db")).Once()
		_, err := materiesService.GetOneMateri(context.Background(), &materiesDomain)
		assert.NotNil(t, err)
	})
}
