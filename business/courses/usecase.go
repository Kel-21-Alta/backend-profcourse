package courses

import (
	"context"
	controller "profcourse/controllers"
	"time"
)

type coursesUsecase struct {
	UserMysqlRepository Repository
	ContextTimeOut      time.Duration
}

func (c coursesUsecase) CreateCourse(ctx context.Context, domain Domain) (Domain, error) {
	// Validasi
	if domain.Title == "" {
		return Domain{}, controller.TITLE_EMPTY
	}

	if domain.Description == "" {
		return Domain{}, controller.DESC_EMPTY
	}

	if domain.FileImage == nil {
		return Domain{}, controller.FILE_IMAGE_EMPTY
	}

}

func NewCourseUseCase(r Repository, timeout time.Duration) Usecase {
	return &coursesUsecase{
		UserMysqlRepository: r,
		ContextTimeOut:      timeout,
	}
}
