package courses

import (
	"context"
	"mime/multipart"
	"time"
)

type Domain struct {
	ID          string
	Title       string
	Description string
	ImgUrl      string
	TeacherId   string
	TeacherName string
	Status      int8
	StatusText  string
	FileImage   *multipart.FileHeader

	CertificateId string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	CreateCourse(ctx context.Context, domain *Domain) (*Domain, error)
}

type Repository interface {
	CreateCourse(ctx context.Context, domain *Domain) (*Domain, error)
}
