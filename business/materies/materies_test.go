package materies_test

import (
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
	t.Run("Test case modul 2 | handle error modul id empty", func(t *testing.T) {
		setUpCreateMateri()
		_, err := materiesService.CreateMateri(context.Background(), &materies.Domain{
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
		_, err := materiesService.CreateMateri(context.Background(), &materies.Domain{
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
		_, err := materiesService.CreateMateri(context.Background(), &materies.Domain{
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
		_, err := materiesService.CreateMateri(context.Background(), &materies.Domain{
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
		_, err := materiesService.CreateMateri(context.Background(), &materies.Domain{
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
		_, err := materiesService.CreateMateri(context.Background(), &materies.Domain{
			Title:     "Pengenalan",
			ModulId:   "3023a588-70c9-49d5-8698-c1b37939f3d8",
			Order:     1,
			Type:      3,
			UrlMateri: "https://www.youtube.com/watch?v=nyGu8Xn5b3g&list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw&index=1",
		})
		assert.Equal(t, controller.TYPE_MATERI_WRONG, err)
	})
}
