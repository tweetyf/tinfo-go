package appctl

import (
	"net/http"
	"time"
	"tinfo-go/model"
	"tinfo-go/utils"

	"github.com/labstack/echo/v4"
)

// ======================================SESSION ===========================
//New Session
type Ssid struct {
	Ssid   string        `json:"ssid"`
	Sskey  string        `json:"sskey"`
	Expire time.Duration `json:"expire"`
}

func SSIDNew(c echo.Context) error {
	cssid, _ := utils.RandomString(16)
	csskey, _ := utils.RandomString(32)
	model.SSID_CACHE.Set(cssid, cssid, 0)
	return c.JSON(http.StatusOK, Ssid{
		cssid,
		csskey,
		model.SSID_TIMEOUT / time.Millisecond,
	})
}

// Check if it is a valid SSID, for internal use only!
func SSIDcheck(ssid string, sskey string) bool {
	if (ssid == "") || (sskey == "") {
		return false
	}
	sskey2, found := model.SSID_CACHE.Get(ssid)
	return found && (sskey == sskey2)
}

//Check if it is a valid SSID
func SSIDCheck(c echo.Context) error {
	ssid := c.QueryParam("ssid")
	sskey := c.QueryParam("sskey")
	if SSIDcheck(ssid, sskey) {
		return c.JSON(http.StatusOK, Status{
			Status: "matched",
		})
	} else {
		return c.JSON(http.StatusOK, Status{
			Status: "mismatched",
		})
	}
}

// Write off the SSID
func SSIDClose(c echo.Context) error {
	ssid := c.QueryParam("ssid")
	sskey := c.QueryParam("sskey")
	if SSIDcheck(ssid, sskey) {
		model.SSID_CACHE.Delete(ssid)
		return c.JSON(http.StatusOK, Status{
			Status: "closed",
		})
	} else {
		return c.JSON(http.StatusOK, Status{
			Status: "mismatched",
		})
	}
}
