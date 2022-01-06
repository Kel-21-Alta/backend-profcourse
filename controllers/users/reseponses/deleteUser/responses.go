package deleteUser

import "profcourse/business/users"

type DeleteUserResponse struct {
	Message string
}


func FromDomain(domain users.Domain) *DeleteUserResponse {
	return &DeleteUserResponse{Message: domain.Message}
}