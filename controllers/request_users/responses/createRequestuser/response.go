package createrequestuser

import "profcourse/business/request_users"

type CreateRequestUser struct {
	Topik             string `json:"topik"`
	CategoryRequestId string `json:"category_request_id"`
	CategoryRequest   string `json:"category_request"`
	Id string `json:"id"`
}

func FromDomain(domain request_users.Domain) *CreateRequestUser {
	return &CreateRequestUser{Topik: domain.Request, CategoryRequestId: domain.CategoryID, CategoryRequest: domain.Category.Title, Id: domain.Id}
}
