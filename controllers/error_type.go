package controller

import "errors"

var EMPTY_NAME = errors.New("Name kosong")
var EMPTY_EMAIL = errors.New("Email kosong")
var INVALID_EMAIL = errors.New("Email tidak valid")
var EMAIL_UNIQUE = errors.New("Email telah digunakan")
var PASSWORD_EMPTY = errors.New("Password kosong")
var WRONG_PASSWORD = errors.New("Password salah")
var WRONG_EMAIL = errors.New("Email salah")
var FORBIDDIN_USER = errors.New("User tidak diizinkan")
var ID_EMPTY = errors.New("User tidak dikenali")
