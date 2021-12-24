package changePassword

type ChangePasswordResponse struct {
	Message string `json:"message"`
}

func GenerateMessage() *ChangePasswordResponse {
	return &ChangePasswordResponse{Message: "Password user berhasil diubah"}
}
