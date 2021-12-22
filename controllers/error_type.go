package controller

import "errors"

var EMPTY_NAME = errors.New("Name Empty")
var EMPTY_EMAIL = errors.New("Email empty")
var INVALID_EMAIL = errors.New("Email Tidak Valid")
var EMAIL_UNIQUE = errors.New("Email telah digunakan")
