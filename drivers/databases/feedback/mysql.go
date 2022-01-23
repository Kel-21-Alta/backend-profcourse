package feedback

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/feedback"
)

type FeedbackRepository struct {
	Conn *gorm.DB
}

func (f FeedbackRepository) GetAvegareRatingCourse(ctx context.Context, domain *feedback.CourseReviews) (feedback.CourseReviews, error) {
	var result float32

	err := f.Conn.Table("feedbacks").Select("AVG(rating) as result").Where("course_id = ?",domain.CourseId).Scan(&result).Error

	if err != nil {
		return feedback.CourseReviews{}, err
	}
	return feedback.CourseReviews{RatingAll: result}, err
}

func (f FeedbackRepository) GetAllFeedbackByCourse(ctx context.Context, domain *feedback.CourseReviews) (feedback.CourseReviews, error) {
	var recs []Feedback

	err := f.Conn.Preload("User").Where("course_id = ?", domain.CourseId).Find(&recs).Order("rating desc").Error

	if err != nil {
		return feedback.CourseReviews{}, err
	}

	result := ToCourseReview(recs)

	return result, nil
}

func (f FeedbackRepository) GetOneFeedbackByUserAndCourse(ctx context.Context, domain *feedback.Domain) (feedback.Domain, error) {
	var rec Feedback

	if err := f.Conn.Where("user_id = ?", domain.UserId).Where("course_id = ?", domain.CourseId).First(&rec).Error; err != nil {
		return feedback.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (f FeedbackRepository) UpdateFeedback(ctx context.Context, domain *feedback.Domain) (feedback.Domain, error) {
	var rec Feedback

	if err := f.Conn.First(&rec, "id = ?", domain.ID).Error; err != nil {
		return feedback.Domain{}, err
	}

	rec.Review = domain.Review
	rec.Rating = domain.Rating

	if err := f.Conn.Save(&rec).Error; err != nil {
		return feedback.Domain{}, nil
	}
	return rec.ToDomain(), nil
}

func (f FeedbackRepository) CreateFeedback(ctx context.Context, domain *feedback.Domain) (feedback.Domain, error) {
	var rec = FromDomain(domain)
	var err error

	if f.Conn.Model(&rec).Where("user_id = ?", rec.UserId).Where("course_id = ?", rec.CourseId).Updates(&rec).RowsAffected == 0 {
		err = f.Conn.Create(&rec).Error
	}

	if err != nil {
		return feedback.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) feedback.Repository {
	return &FeedbackRepository{Conn: conn}
}
