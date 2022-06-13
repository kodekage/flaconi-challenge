package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
)

var v = validator.New()

func requestHandler(context echo.Context) error {
	var requestBody []map[string]interface{}
	queries := context.QueryParams()

	if len(queries) == 0 {
		err := errors.New("nesting parameters are missing in request")
		return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	}

	if nestingParams, ok := queries["nesting_params"]; !ok {
		err := errors.New("nesting_params missing")
		return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	} else {
		if err := context.Bind(&requestBody); err != nil {
			return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
		}

		if err := v.Var(requestBody, "required,dive"); err != nil {
			return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
		}

		result := inputJsonParser(requestBody, nestingParams)
		outputResult(result)

		return context.JSON(http.StatusOK, "OK")
	}
}
