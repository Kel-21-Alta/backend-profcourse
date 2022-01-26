package request

import "profcourse/business/users_spesializations"

type RegisterSpesializationRequest struct {
	SpesializationID string `json:"spesialization_id"`
}

func (r RegisterSpesializationRequest) ToDomain() *users_spesializations.Domain {
	return &users_spesializations.Domain{SpesializationID: r.SpesializationID}
}


