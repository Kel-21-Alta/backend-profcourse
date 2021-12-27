package courses

import (
	"context"
	"profcourse/business/locals"
	controller "profcourse/controllers"
	"time"
)

type coursesUsecase struct {
	CourseMysqlRepository Repository
	ContextTimeOut        time.Duration
	LocalRepository       locals.Repository
}

func (c coursesUsecase) CreateCourse(ctx context.Context, domain *Domain) (*Domain, error) {
	// Validasi
	if domain.Title == "" {
		return &Domain{}, controller.TITLE_EMPTY
	}

	if domain.Description == "" {
		return &Domain{}, controller.DESC_EMPTY
	}

	if domain.FileImage == nil {
		return &Domain{}, controller.FILE_IMAGE_EMPTY
	}

	localRepoDomain, err := c.LocalRepository.UploadImage(ctx, domain.FileImage, "/img/courses/")
	if err != nil {
		return &Domain{}, err
	}

	domain.ImgUrl = localRepoDomain.ResultUrl
	course, err := c.CourseMysqlRepository.CreateCourse(ctx, domain)

	if err != nil {
		return &Domain{}, err
	}
	return course, nil
}

func NewCourseUseCase(r Repository, timeout time.Duration, local locals.Repository) Usecase {
	return &coursesUsecase{
		CourseMysqlRepository: r,
		ContextTimeOut:        timeout,
		LocalRepository:       local,
	}
}
