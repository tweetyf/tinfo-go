package sysctl

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Status struct {
	Status string `json:"status"`
}

// ======================================ADMIN ===========================
func AdminLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, Status{Status: "ok"})
}
func AdminChangePWD(c echo.Context) error {
	return c.JSON(http.StatusOK, Status{Status: "ok"})
}
func AdminLogout(c echo.Context) error {
	return c.JSON(http.StatusOK, Status{Status: "ok"})
}
