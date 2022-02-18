package appctl

import (
	"net/http"
	"tinfo-go/utils"

	"github.com/labstack/echo/v4"
)

// ======================================MiddleWares ===========================
// Example: Ban all ips from China
func MID_checkIPCountry(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check the request ip origin, of they are from some country them deny request.
		ipc := utils.GetIPCountry(c.RealIP())
		if ipc == "CN" {
			utils.LogD("IP rule vialation: deny service for host %v from: %v", c.RealIP(), ipc)
			c.Error(nil)
		} else {
			return next(c)
		}
		return nil
	}

}

// ============================Other tools =================
// look for country code from given IP address.
func IPTCountry(c echo.Context) error {
	cip := c.QueryParam("ip")
	res := utils.GetIPCountry(cip)
	var cnt struct {
		Ip      string `json:"ip"`
		Country string `json:"country"`
		Status  string `json:"status"`
	}
	cnt.Country = res
	cnt.Ip = cip
	cnt.Status = "ok"
	utils.LogD("IPTCountry: " + cip + " country: " + res)
	return c.JSON(http.StatusOK, &cnt)
}
