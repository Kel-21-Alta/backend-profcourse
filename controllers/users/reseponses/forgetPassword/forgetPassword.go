package forgetPassword

type ForgetPasswordResponse struct {
	Message string `json:"message"`
}

func GenerateResponses() *ForgetPasswordResponse {
	return &ForgetPasswordResponse{Message: "Berhasil mengirim email lupa password"}
}
