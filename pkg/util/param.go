package util

import (
	"github.com/go-playground/validator/v10" //nolint
	"github.com/labstack/echo/v4"            //nolint
)

// ParsingParameter will parsing request to struct
func ParsingParameter(ctx echo.Context, i interface{}) error { //nolint
	err := ctx.Bind(i)
	if err != nil {
		return &ParsingError{err.Error()}
	}
	return err
}

// ValidateParameter will validate request
func ValidateParameter(ctx echo.Context, i interface{}) (err error) { //nolint
	validate := validator.New() //nolint
	err = validate.Struct(i)

	return
}
