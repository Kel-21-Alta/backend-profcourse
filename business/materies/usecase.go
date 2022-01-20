package materies

import (
	"fmt"
	"golang.org/x/net/context"
	"profcourse/business/users_courses"
	controller "profcourse/controllers"
	"time"
)

type MateriesUsecase struct {
	MateriesRepository Repository
	ContextTimeout     time.Duration
	UserCourse         users_courses.Usecase
}

func (u MateriesUsecase) UpdateProgressMateri(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_MATERI_EMPTY
	}
	if domain.User.ID == "" {
		return Domain{}, controller.ID_EMPTY
	}

	if domain.User.CourseId == "" {
		return Domain{}, controller.EMPTY_COURSE
	}

	// mencari user_course_id
	userCourseDomain, err := u.UserCourse.GetOneUserCourse(ctx, &users_courses.Domain{CourseId: domain.User.CourseId, UserId: domain.User.ID})

	domain.UserCourse.UserCourseId = userCourseDomain.ID

	resultUpdateProgressMateri, err := u.MateriesRepository.UpdateProgressMateri(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	// menghitung progress
	materiFinist, err := u.MateriesRepository.GetCountMateriFinish(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	// jumlah materi => SELECT COUNT(*) FROM materis INNER JOIN moduls m on materis.modul_id = m.id INNER JOIN courses c on m.course_id = c.id where course_id = "a7234d7d-ebc5-495c-ad41-782f3eb906b8"
	allMateriCourse, err := u.MateriesRepository.GetCountMateriCourse(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	// materi yang selesai / jumlah materi *100

	progress := (materiFinist*100)  / allMateriCourse
	fmt.Println(materiFinist)
	fmt.Println(allMateriCourse)
	fmt.Println(progress)
	_, err = u.UserCourse.UpdateProgressCourse(ctx, &users_courses.Domain{
		UserId:      resultUpdateProgressMateri.User.ID,
		CourseId:    resultUpdateProgressMateri.User.CourseId,
		Progres:     progress,
		LastVideoId: resultUpdateProgressMateri.ID,
		LastModulId: resultUpdateProgressMateri.ModulId,
	})

	if err != nil {
		return Domain{}, err
	}
	// TODO: Update Progress Spesialization
	return resultUpdateProgressMateri, nil
}

func (u MateriesUsecase) GetAllMateri(ctx context.Context, domain *Domain) (AllMateriModul, error) {
	if domain.ModulId == "" {
		return AllMateriModul{}, controller.EMPTY_MODUL_ID
	}
	if domain.User.ID == "" {
		return AllMateriModul{}, controller.ID_EMPTY
	}

	result, err := u.MateriesRepository.GetAllMateri(ctx, domain)

	if err != nil {
		return AllMateriModul{}, err
	}

	return result, nil
}

func (u MateriesUsecase) GetOneMateri(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_MATERI_EMPTY
	}

	// TODO: Validasi apakah user terdaftar di course materi ini

	result, err := u.MateriesRepository.GetOnemateri(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (u MateriesUsecase) DeleteMateri(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_MATERI_EMPTY
	}

	result, err := u.MateriesRepository.DeleteMateri(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (u MateriesUsecase) ValidasiMateri(ctx context.Context, domain *Domain) (*Domain, error) {

	if domain.ModulId == "" {
		return &Domain{}, controller.EMPTY_MODUL_ID
	}
	// TODO: Validasi apakah modul id ada
	// TODO: Validasi apakah user yang sedang login adalah pemilik course atau admin
	if domain.Title == "" {
		return &Domain{}, controller.TITLE_EMPTY
	}
	if domain.UrlMateri == "" {
		return &Domain{}, controller.EMPTY_FILE_MATERI
	}
	if domain.Order == 0 {
		return &Domain{}, controller.ORDER_MATERI_EMPTY
	}
	if domain.Type == 0 {
		return &Domain{}, controller.TYPE_MATERI_EMPTY
	}
	if domain.Type < 1 || domain.Type > 2 {
		return &Domain{}, controller.TYPE_MATERI_WRONG
	}

	return domain, nil
}

func (u MateriesUsecase) UpdateMateri(ctx context.Context, domain *Domain) (Domain, error) {

	if domain.ID == "" {
		return Domain{}, controller.ID_MATERI_EMPTY
	}

	resultValidasi, err := u.ValidasiMateri(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	materi, err := u.MateriesRepository.UpdateMateri(ctx, resultValidasi)

	if err != nil {
		return Domain{}, err
	}

	return materi, nil
}

func (u MateriesUsecase) CreateMateri(ctx context.Context, domain *Domain) (Domain, error) {

	resultValidasi, err := u.ValidasiMateri(ctx, domain)

	materi, err := u.MateriesRepository.CreateMateri(ctx, resultValidasi)

	if err != nil {
		return Domain{}, err
	}

	return materi, nil
}

func NewMateriesUsecase(repo Repository, userCourse users_courses.Usecase, timeout time.Duration) Usecase {
	return &MateriesUsecase{MateriesRepository: repo, UserCourse: userCourse, ContextTimeout: timeout}
}
