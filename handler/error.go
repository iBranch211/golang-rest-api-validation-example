package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// APIError example
type APIError struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp int    `json:"timestamp"`
}

func ErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	code := he.Code
	message := he.Message
	if m, ok := he.Message.(string); ok {
		message =
			echo.Map{
				"status":    code,
				"message":   m,
				"path":      c.Request().RequestURI,
				"timestamp": time.Now().Unix(),
			}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			c.JSON(code, message)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func ResourceNotFoundException(resourceName string, fieldName string, fieldValue string) error {
	msg := fmt.Sprintf("%s not found with %s : %s", resourceName, fieldName, fieldValue)
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

func BadRequestException(msg string) error {
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func ConflictException(resourceName string, fieldName string, fieldValue string) error {
	msg := fmt.Sprintf("%s with %s : %s already exists", resourceName, fieldName, fieldValue)
	return echo.NewHTTPError(http.StatusConflict, msg)
}
