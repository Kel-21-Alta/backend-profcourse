package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseSuccess(c echo.Context, status int, params interface{}) error {
	response := BaseResponse{}
	response.Data = params
	response.Message = "Success"
	response.Code = status
	return c.JSON(status, response)
}

func NewResponseError(c echo.Context, err error) error {
	status := CheckStatus(err)
	response := BaseResponse{
		Code:    status,
		Message: err.Error(),
		Data:    nil,
	}
	return c.JSON(status, response)
}

func CheckStatus(err error) int {
	if err == EMPTY_EMAIL || err == EMPTY_NAME || err == INVALID_EMAIL || err == EMAIL_UNIQUE {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}