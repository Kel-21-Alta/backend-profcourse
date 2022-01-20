package getAllSummary

import "profcourse/business/summary"

type GetAllSummaryResponse struct {
	CountCourse int `json:"count_course"`
	CountUser int `json:"count_user"`
	CountSpesiazation int `json:"count_spesiazation"`
}

func FromDomain(domain *summary.Domain) *GetAllSummaryResponse {
	return &GetAllSummaryResponse{
		CountCourse:       domain.CountCourse,
		CountUser:         domain.CountUser,
		CountSpesiazation: domain.CountSpesialization,
	}
}
