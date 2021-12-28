package controller

import "errors"

var EMPTY_NAME = errors.New("Nama kosong")
var EMPTY_EMAIL = errors.New("Email kosong")
var INVALID_EMAIL = errors.New("Email tidak valid")
var EMAIL_UNIQUE = errors.New("Email telah digunakan")
var PASSWORD_EMPTY = errors.New("Kata sandi kosong")
var WRONG_PASSWORD = errors.New("Kata sandi salah")
var WRONG_EMAIL = errors.New("Email salah")
var FORBIDDIN_USER = errors.New("User tidak diizinkan")
var ID_EMPTY = errors.New("User tidak dikenali")
var EMPTY_USER = errors.New("User tidak terdaftar")
var TITLE_EMPTY = errors.New("Judul kosong")
var DESC_EMPTY = errors.New("Deskripsi kosong")
var FILE_IMAGE_EMPTY = errors.New("File gambar kosong")
var INVALID_FILE = errors.New("File tidak valid")
var INVALID_PARAMS = errors.New("Parameter tidak valid")
var BAD_REQUEST = errors.New("Request tidak valid")
var EMPTY_COURSE = errors.New("Kusus tidak ada")
var ALREADY_REGISTERED_COURSE = errors.New("User sudah mendaftar kursus")
