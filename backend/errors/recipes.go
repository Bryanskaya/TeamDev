package errors

import (
	"errors"
)

var (
	AccessDeleteDenied = errors.New("Access to delete is denied")

	UnknownRecipe = errors.New("Recipe is not found")
)