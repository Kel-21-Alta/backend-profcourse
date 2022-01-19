package materies

import (
	"golang.org/x/net/context"
	"time"
)

type Domain struct {
	ID         string
	Title      string
	ModulId    string
	Order      int
	Type       int
	TypeString string
	UrlMateri  string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User       CurrentUser
	UserCourse UserCourse
}

type AllMateriModul struct {
	JawabanMateri int `json:"jawaban_materi"`
	Materi        []Domain
}

type CurrentUser struct {
	ID          string
	CurrentTime string
	IsComplate  bool
	CourseId    string
}

type UserCourse struct {
	UserCourseId string
}

type Usecase interface {
	ValidasiMateri(ctx context.Context, domain *Domain) (*Domain, error)
	CreateMateri(ctx context.Context, domain *Domain) (Domain, error)
	UpdateMateri(ctx context.Context, domain *Domain) (Domain, error)
	DeleteMateri(ctx context.Context, domain *Domain) (Domain, error)
	GetOneMateri(ctx context.Context, domain *Domain) (Domain, error)
	GetAllMateri(ctx context.Context, domain *Domain) (AllMateriModul, error)
	UpdateProgressMateri(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	DeleteMateri(ctx context.Context, domain *Domain) (Domain, error)
	CreateMateri(ctx context.Context, domain *Domain) (Domain, error)
	UpdateMateri(ctx context.Context, domain *Domain) (Domain, error)
	GetOnemateri(ctx context.Context, domain *Domain) (Domain, error)
	GetAllMateri(ctx context.Context, domain *Domain) (AllMateriModul, error)
	UpdateProgressMateri(ctx context.Context, domain *Domain) (Domain, error)
	GetCountMateriFinish(ctx context.Context, domain *Domain) (int, error)
	GetCountMateriCourse(ctx context.Context, domain *Domain) (int, error)
}
