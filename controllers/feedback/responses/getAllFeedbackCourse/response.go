package getAllFeedbackCourse

import "profcourse/business/feedback"

type GetAllFeedbackCourseResponse struct {
	CourseId  string   `json:"course_id"`
	RatingAll float32  `json:"rating_all"`
	Review    []Review `json:"review"`
}

type Review struct {
	ID           string `json:"id"`
	NameUser     string `json:"name_user"`
	UrlImage string `json:"url_image"`
	Body         string `json:"body"`
	Rating       float32 `json:"rating"`
}

func FromListDomain(domain feedback.CourseReviews) *GetAllFeedbackCourseResponse {
	var listReview []Review

	for _, domain := range domain.Review{
		listReview = append(listReview, Review{
			ID:       domain.ID,
			NameUser: domain.User.NameUser,
			Body:     domain.Review,
			Rating:   domain.Rating,
			UrlImage: domain.User.UrlImageUser,
		})
	}

	return &GetAllFeedbackCourseResponse{
		CourseId:  domain.CourseId,
		RatingAll: domain.RatingAll,
		Review:    listReview,
	}
}
