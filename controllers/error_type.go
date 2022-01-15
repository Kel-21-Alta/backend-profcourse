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
var TITLE_EMPTY = errors.New("Judul tidak boleh kosong")
var DESC_EMPTY = errors.New("Deskripsi kosong")
var IMAGE_EMPTY = errors.New("File gambar kosong")
var INVALID_FILE = errors.New("File tidak valid")
var INVALID_PARAMS = errors.New("Parameter tidak valid")
var BAD_REQUEST = errors.New("Request tidak valid")
var EMPTY_COURSE = errors.New("Kusus tidak boleh kosong")
var ALREADY_REGISTERED_COURSE = errors.New("User sudah mendaftar kursus")
var ORDER_MODUL_EMPTY = errors.New("Order modul tidak ada")
var ORDER_MATERI_EMPTY = errors.New("Order materi tidak ada")
var COURSES_SPESIALIZATION_EMPTY = errors.New("Spesialisasi wajib memiliki kurus")
var SPESIALIZATION_NOT_FOUND = errors.New("Spesialisasi tidak ditemukan")
var EMPTY_MODUL_ID = errors.New("Modul id tidak boleh kosong")
var EMPTY_FILE_MATERI = errors.New("File materi tidak boleh kosong")
var TYPE_MATERI_EMPTY = errors.New("Tipe materi tidak boleh kosong")
var TYPE_MATERI_WRONG = errors.New("Tipe materi salah atau tidak valid")
var ID_MATERI_EMPTY = errors.New("ID materi tidak boleh kosng")
