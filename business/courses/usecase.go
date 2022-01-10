package courses

import (
	"context"
	controller "profcourse/controllers"
	"profcourse/helpers"
	"time"
)

type coursesUsecase struct {
	CourseMysqlRepository Repository
	ContextTimeOut        time.Duration
}

func (c *coursesUsecase) DeleteCourse(ctx context.Context, id string) (Domain, error) {
	if id == "" {
		return Domain{}, controller.ID_EMPTY
	}
	result, err := c.CourseMysqlRepository.DeleteCourse(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (c *coursesUsecase) UpdateCourse(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Title == "" {
		return Domain{}, controller.TITLE_EMPTY
	}
	if domain.Description == "" {
		return Domain{}, controller.DESC_EMPTY
	}
	if domain.ImgUrl == "" {
		return Domain{}, controller.IMAGE_EMPTY
	}

	result, err := c.CourseMysqlRepository.UpdateCourse(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	if int8(result.Status) == 1 {
		result.StatusText = "Publish"
	} else if int8(result.Status) == 3 {
		result.StatusText = "Pending"
	} else if int8(result.Status) == 2 {
		result.StatusText = "Draft"
	}
	return result, nil
}

func (c *coursesUsecase) GetCountCourse(ctx context.Context) (*Summary, error) {
	domain, err := c.CourseMysqlRepository.GetCountCourse(ctx)
	if err != nil {
		return &Summary{}, err
	}
	return domain, nil
}

func (c *coursesUsecase) GetOneCourse(ctx context.Context, domain *Domain) (*Domain, error) {

	course, err := c.CourseMysqlRepository.GetOneCourse(ctx, domain)
	if err != nil {
		return &Domain{}, err
	}
	course.InfoUser.CurrentUser = domain.InfoUser.CurrentUser
	return course, nil
}

func (c *coursesUsecase) GetAllCourses(ctx context.Context, domain *Domain) (*[]Domain, error) {

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

func (c *coursesUsecase) CreateCourse(ctx context.Context, domain *Domain) (*Domain, error) {
	// Validasi
	if domain.Title == "" {
		return &Domain{}, controller.TITLE_EMPTY
	}

	if domain.Description == "" {
		return &Domain{}, controller.DESC_EMPTY
	}

	if domain.ImgUrl == "" {
		return &Domain{}, controller.IMAGE_EMPTY
	}
	course, err := c.CourseMysqlRepository.CreateCourse(ctx, domain)

	if err != nil {
		return &Domain{}, err
	}
	return course, nil
}

func NewCourseUseCase(r Repository, timeout time.Duration) Usecase {
	return &coursesUsecase{
		CourseMysqlRepository: r,
		ContextTimeOut:        timeout,
	}
}
