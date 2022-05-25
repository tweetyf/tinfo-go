package appctl

import (
	"fmt"
	"image/color"
	"image/png"
	"net/http"
	"strings"
	"time"
	"tinfo-go/app"
	"tinfo-go/model"
	"tinfo-go/utils"

	"github.com/afocus/captcha"
	"github.com/labstack/echo/v4"
)

// ======================================CAPTCHA ===========================
var cap *captcha.Captcha

func InitCaptchaHandler() {
	cap = captcha.New()

	/*
		if err := cap.SetFont("./static/comic.ttf"); err != nil {
			panic(err.Error())
		}
	*/
	fontContenrs, err := app.Static.ReadFile("static/comic.ttf")
	if err != nil {
		panic(err.Error())
	}
	err = cap.AddFontFromBytes(fontContenrs)
	if err != nil {
		panic(err.Error())
	}
	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
}

func CaptchaNew(c echo.Context) error {
	a := c.QueryParam("a")
	if a == "fetch" {
		return captchaGen(c)
	} else {
		cid, _ := utils.RandomString(16)
		cstr, _ := utils.RandomString(6)
		var cnt struct {
			Cid     string        `json:"cid"`
			Capaddr string        `json:"cap_addr"`
			Expire  time.Duration `json:"expire"`
		}
		cnt.Cid = cid
		cnt.Capaddr = c.Request().RequestURI + "?a=fetch&cid=" + cid
		cnt.Expire = model.CACHE_TIMEOUT / time.Millisecond
		//put ssid and string into highspeed cache.
		model.CAPTCHA_CACHE.Set(cid, cstr, 0)
		return c.JSON(http.StatusOK, &cnt)
	}
}

// Genererate a new captcha. Need cid
func captchaGen(c echo.Context) error {
	cid := c.QueryParam("cid")
	_cstr, found := model.CAPTCHA_CACHE.Get(cid)
	cstr := fmt.Sprintf("%v", _cstr)
	if found {
		img := cap.CreateCustom(cstr)
		utils.LogD("CAPTCHA new str: " + cstr + " ssid: " + cid)
		return png.Encode(c.Response().Writer, img)
	} else {
		return c.JSON(http.StatusOK, Status{
			Status: "invalid cid.",
		})
	}
}

// verify Captcha: need cstr and cid
func CaptchaVer(c echo.Context) error {
	cstr := c.QueryParam("cstr")
	cid := c.QueryParam("cid")
	utils.LogD("CAPTCHA verify str: " + cstr + " cid: " + cid)
	if CaptchaVerStr(cstr, cid) {
		// captcha image only be used once, when transmitting finished, just delete it.
		model.CAPTCHA_CACHE.Delete(cid)
		return c.JSON(http.StatusOK, Status{
			Status: "matched",
		})
	} else {
		return c.JSON(http.StatusNotFound, Status{
			Status: "captcha mismatch!",
		})
	}
}

//verify Captcha from memory
func CaptchaVerStr(cstr string, cid string) bool {
	cstr2, found := model.CAPTCHA_CACHE.Get(cid)
	cstr22 := fmt.Sprintf("%v", cstr2)
	return found && (strings.ToLower(cstr) == strings.ToLower(cstr22))
}
