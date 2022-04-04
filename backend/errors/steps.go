package errors

import (
	"errors"
)

var (
	UnknownStep = errors.New("Step is not found")
	StepExists = errors.New("Step is already exists")
)