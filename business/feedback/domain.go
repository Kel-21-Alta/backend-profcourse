package feedback

import (
	"golang.org/x/net/context"
	"time"
)

type Domain struct {
	ID        string
	UserId    string
	CourseId  string
	Review    string
	Rating    float32
	CreatedAt time.Time
	UpdateAt  time.Time
}

type Usecase interface {
	CreateFeedback(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	GetOneFeedbackByUserAndCourse(ctx context.Context, domain *Domain) (Domain, error)
	UpdateFeedback(ctx context.Context, domain *Domain) (Domain, error)
	CreateFeedback(ctx context.Context, domain *Domain) (Domain, error)
}
