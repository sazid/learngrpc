package service

import "errors"

var ErrAlreadyExists = errors.New("record already exists")
var ErrNotFound = errors.New("record not found")
