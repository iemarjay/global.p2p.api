package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type appValidator struct {
	validator *validator.Validate
}

func NewAppValidator(validator *validator.Validate) *appValidator {
	return &appValidator{validator: validator}
}

func (av *appValidator) Validate(i interface{}) error {
	if err := av.validator.Struct(i); err != nil {
		//for _, err := range err.(validator.ValidationErrors) {
		//
		//	fmt.Println(err)
		//	fmt.Println(err.Namespace())
		//	fmt.Println(err.Field())
		//	fmt.Println(err.StructNamespace())
		//	fmt.Println(err.StructField())
		//	fmt.Println(err.Tag())
		//	fmt.Println(err.ActualTag())
		//	fmt.Println(err.Kind())
		//	fmt.Println(err.Type())
		//	fmt.Print(err.Value())
		//	fmt.Println( "value")
		//	fmt.Print(err.Param())
		//	fmt.Println("param")
		//}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

