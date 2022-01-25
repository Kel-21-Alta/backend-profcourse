package requests

import "profcourse/business/request_users"

type UpdateRequestUser struct {
	Id string `json:"id"`
	CategoryRequestId string `json:"category_request_id"`
	Topik             string `json:"topik"`
}

func (c *UpdateRequestUser) ToDomain() request_users.Domain {
	return request_users.Domain{CategoryID: c.CategoryRequestId, Request: c.Topik}
}

