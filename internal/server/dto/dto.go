package dto

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/texm/dokku-go"
	"net/http"
	"reflect"
	"regexp"
)

// allow alphanumeric, underscores, and hyphens
func appNameCharsValidator() func(level validator.FieldLevel) bool {
	r, err := regexp.Compile("\\w[\\w\\-]*")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to compile regexp")
	}
	return func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() != reflect.String {
			return false
		}
		str := fl.Field().String()
		return r.FindString(str) == str
	}
}

type requestValidator struct {
	validator *validator.Validate
}

func (rv *requestValidator) Validate(i interface{}) error {
	return rv.validator.Struct(i)
}

func NewRequestValidator() *requestValidator {
	v := validator.New()

	err := v.RegisterValidation("appNameChars", appNameCharsValidator())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to register validator")
	}
	v.RegisterAlias("appName", "appNameChars,min=4,max=32")
	v.RegisterAlias("processName", "appNameChars,min=2,max=32")

	return &requestValidator{validator: v}
}

type RequestError struct {
	err              error
	isBinding        bool
	isInvalidFormat  bool
	isInvalidData    bool
	validationErrors []validator.FieldError
}

func (r *RequestError) String() string {
	return fmt.Sprintf("%+v", r.validationErrors)
}

func (r *RequestError) ToHTTP() *echo.HTTPError {
	err := echo.NewHTTPError(http.StatusBadRequest).SetInternal(r.err)

	if r.isBinding {
		err.Message = echo.Map{"type": "binding"}
	} else if r.isInvalidFormat {
		err.Message = echo.Map{"type": "format"}
	} else if r.isInvalidData {
		fields := map[string]string{}
		for _, fe := range r.validationErrors {
			fields[fe.Field()] = fe.ActualTag()
		}
		err.Message = echo.Map{
			"type":   "validation",
			"fields": fields,
		}
	} else {
		err = echo.ErrInternalServerError
	}

	return err
}

func BindRequest(c echo.Context, r any) *RequestError {
	if err := c.Bind(r); err != nil {
		return &RequestError{
			err:       err,
			isBinding: true,
		}
	}
	if err := c.Validate(r); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			return &RequestError{
				err:              err,
				isInvalidData:    true,
				validationErrors: errors,
			}
		}
		return &RequestError{
			err:             err,
			isInvalidFormat: true,
		}
	}
	return nil
}

func MaybeConvertDokkuError(err error) *echo.HTTPError {
	if errors.Is(err, dokku.InvalidAppError) {
		return echo.NewHTTPError(http.StatusBadRequest, "no such app")
	}
	if errors.Is(err, dokku.NameTakenError) {
		return echo.NewHTTPError(http.StatusBadRequest, "duplicate app name")
	}
	return nil
}
