package courses

import (
	"context"
	"profcourse/business/locals"
	controller "profcourse/controllers"
	"profcourse/helpers"
	"time"
)

type coursesUsecase struct {
	CourseMysqlRepository Repository
	ContextTimeOut        time.Duration
	LocalRepository       locals.Repository
}

func (c coursesUsecase) GetOneCourse(ctx context.Context, domain *Domain) (*Domain, error) {

	course, err := c.CourseMysqlRepository.GetOneCourse(ctx, domain)
	if err != nil {
		return &Domain{}, err
	}
	return course, nil
}

func (c coursesUsecase) GetAllCourses(ctx context.Context, domain *Domain) (*[]Domain, error) {

	if domain.SortBy == "" {
		domain.SortBy = "asc"
	}

	if domain.SortBy == "dsc" {
		domain.SortBy = "desc"
	}

	if domain.Sort == "" {
		domain.Sort = "created_at"
	}

	// menvalidasi sort by yang diizinkan
	sortByAllow := []string{"asc", "desc"}
	if !helpers.CheckItemInSlice(sortByAllow, domain.SortBy) {
		return &[]Domain{}, controller.INVALID_PARAMS
	}

	// Menvalidasi sort yang diizinkan
	sortAllow := []string{"created_at", "title"} // TODO: disini kurang sort review dan sort popular
	if !helpers.CheckItemInSlice(sortAllow, domain.Sort) {
		return &[]Domain{}, controller.INVALID_PARAMS
	}

	listCourseDomain, err := c.CourseMysqlRepository.GetAllCourses(ctx, domain)
	if err != nil {
		return &[]Domain{}, err
	}

	return listCourseDomain, nil
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