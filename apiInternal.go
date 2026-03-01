package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func iApiVersion(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"version": "1.0.0"})
}
