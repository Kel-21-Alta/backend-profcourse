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

type DataError struct {
	Message string `json:"message"`
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
		Data:    DataError{Message: err.Error()},
	}
	return c.JSON(status, response)
}

func CheckStatus(err error) int {
	if ALREADY_DELETE == err || ALREADY_REGISTERED_SPESIALIZATION == err || SPESIALIZATION_ID_EMPTY == err || CATEGORY_EMPTY== err || FEEDBACK_EMPTY == err || REVIEW_EMPTY == err || RATING_EMPTY == err || LAST_MODUL_EMPTY == err || LAST_MATERI_EMPTY == err || err == PROGRESS_EMPTY || err == PILIHAN_QUIZ_EMPTY || err == PILIHAN_QUIZ_MINUS || err == PERTANYAAN_QUIZ_EMPTY || err == JAWABAN_QUIZ_EMPTY || err == TYPE_MATERI_EMPTY || err == ORDER_MATERI_EMPTY || err == EMPTY_FILE_MATERI || err == EMPTY_MODUL_ID || err == COURSES_SPESIALIZATION_EMPTY || err == BAD_REQUEST || err == EMPTY_COURSE || err == ALREADY_REGISTERED_COURSE || err == ORDER_MODUL_EMPTY || err == INVALID_PARAMS || err == TITLE_EMPTY || err == DESC_EMPTY || err == INVALID_FILE || err == EMPTY_EMAIL || err == EMPTY_NAME || err == INVALID_EMAIL || err == EMAIL_UNIQUE || err == PASSWORD_EMPTY {
		return http.StatusBadRequest
	}
	if err == FORBIDDIN_USER || err == WRONG_PASSWORD || err == WRONG_EMAIL {
		return http.StatusForbidden
	}
	if err == ID_EMPTY || err == EMPTY_USER {
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
