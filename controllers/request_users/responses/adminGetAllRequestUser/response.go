package adminGetAllRequestUser

import "profcourse/business/request_users"

type GetAllRequestUser struct {
	RequestId string `json:"request_id"`
	Category string `json:"category"`
	CategoryId string `json:"category_id"`
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	Body string `json:"body"`
}

func FromListDomain(domains []request_users.Domain) []GetAllRequestUser {
	var list []GetAllRequestUser

	for _, domain := range domains {
		list = append(list, GetAllRequestUser{
			RequestId: domain.Id,
			Category:  domain.Category.Title,
			Body:      domain.Request,
			CategoryId: domain.CategoryID,
			UserId: domain.UserId,
			UserName: domain.User.Name,
		})
	}
	return  list
}
