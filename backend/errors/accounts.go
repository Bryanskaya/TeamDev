package errors

import (
	"errors"
)

var (
	RecordNotFound  = errors.New("Record is not found")
	UnknownError    = errors.New("Unknown error")
	WrongPassword   = errors.New("Wrong password")
	DBAdditionError = errors.New("Eror in addition new record to DB")

	UnknownAccount = errors.New("Account is not found")
	UnknownRole    = errors.New("Role is not found")
	AccessDenied   = errors.New("Access denied")
	AccountExists  = errors.New("Account associated with such login is already exists")
)
