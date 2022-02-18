package appctl

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Status struct {
	Status string `json:"status"`
}

// Home page of web site
func IndexGET(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/app/home/")
}

//==================================== Other Handlers ===========================
type SContent struct {
	Title   string `json:"title"`
	Header  string `json:"header"`
	Message string `json:"message"`
}

func Handle404(c echo.Context) error {
	cnt := new(SContent)
	cnt.Title = "oops"
	cnt.Header = "404"
	cnt.Message = "beep.."
	return c.JSON(http.StatusNotFound, &cnt)
}

func HandleMove(c echo.Context) error {

	return c.JSON(http.StatusMovedPermanently, SContent{
		Title:   "oops!",
		Header:  "Moved",
		Message: "Oops!, this page has been moved.",
	})
}

func HandleUnauthorized(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, SContent{
		Title:   "oops!",
		Header:  "Access Denied!",
		Message: "DING!",
	})
}
