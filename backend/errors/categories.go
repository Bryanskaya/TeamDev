package errors

import (
	"errors"
)

var (
	UnknownCategory = errors.New("Category is not found")
	CategoryExists  = errors.New("Category associated is already exists")
)
