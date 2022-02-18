package main

import (
	"html/template"
	"io"
	"time"
	"tinfo-go/controller/appctl"
	"tinfo-go/controller/sysctl"
	"tinfo-go/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TplRender is a custom html/template renderer for Echo framework
type TplRender struct {
	templates *template.Template
}

var funcMap template.FuncMap

func init() {
	funcMap = template.FuncMap{"str2html": Str2html, "str2js": Str2js, "date": Date}
}

// Render renders a template document
func (t *TplRender) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	if IS_DEV {
		t.templates, _ = utils.LoadHTMLTmpl("./views", funcMap)
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

// Str2html Convert string to template.HTML type.
func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

// Str2js Convert string to template.JS type.
func Str2js(raw string) template.JS {
	return template.JS(raw)
}

// Date Date
func Date(t time.Time, format string) string {
	return t.Format(format) //"2006-01-02 15:04:05"
}

// initRender for echo
func initRender() *TplRender {
	tpl, _ := utils.LoadHTMLTmpl("./views", funcMap)
	return &TplRender{
		templates: tpl,
	}
}

// Set up the routes of web framework.
func setupRoutes() *echo.Echo {
	appctl.InitCaptchaHandler()
	var r = echo.New()
	r.Renderer = initRender()
	// recovery from 500 error.
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	// custom midware to check the ip security function.
	r.Use(appctl.MID_checkIPCountry)
	// business logic
	r.Static("/app", "./app")
	r.Static("/static", "./static")
	r.Static("/favicon.ico", "./static/favicon.ico")
	// Home pages
	r.GET("/", appctl.IndexGET)
	r.GET("/unauthorized", appctl.HandleUnauthorized)
	// custom session. Create and end session don't need any authentication, so we don't put veryfication here.
	rssid := r.Group("/ssid")
	rssid.GET("/new", appctl.SSIDNew)
	rssid.GET("/check", appctl.SSIDCheck)
	rssid.GET("/close", appctl.SSIDClose)
	// captcha API, no need authentication.
	rcap := r.Group("/captcha")
	rcap.GET("/new", appctl.CaptchaNew)
	rcap.GET("/verify", appctl.CaptchaVer)
	// user management
	radm := r.Group("/admin")
	radm.POST("/login", sysctl.AdminLogin)
	radm.POST("/changepwd", sysctl.AdminChangePWD)
	radm.GET("/logout", sysctl.AdminLogout)
	// message board
	// Ip tools
	ript := r.Group("/iptool")
	ript.GET("/country", appctl.IPTCountry)
	// Others
	//r.NoRoute(Handle404)
	return r
}
