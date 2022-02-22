/*
see doc in /docs/
*/
package main

import (
	"fmt"
	"os"
	"tinfo-go/model"
	"tinfo-go/utils"
)

func install() {
	// For some initiation work. Currently Dont
	utils.LogD("Website initiating... installing database")
	var conf = initConfig()
	model.AddNewPref("site_name", conf.SITE_NAME)
	model.AddNewPref("site_version", conf.SITE_VERSION)
	model.AddNewPref("site_desc", conf.SITE_DESC)
	model.AddNewPref("site_port", conf.SITE_PORT)
	model.AddNewUser(conf.ADMIN_UNAME, conf.ADMIN_PWD, conf.ADMIN_EMAIL, "admin", conf.ADMIN_AVATAR, "")
}

func deleledb() {
	utils.LogD("Website delete.")
}

var helpStr = `--install: install.
--test: test only.
--reset:WARNNING: this will delete everything!
--help: show this help
--geoip [ip] looking for contry code of an IP
--md2html [.md file] translate markdown file to HTML
-p port choose witch port to listen to.
`

func main() {
	// set log level, if we were in release mod, we just set it to LOG_LEVEL_RELEASE
	utils.SetLogLevel(utils.LOG_LEVEL_DEBUG)
	//TODO: make conf be configurable here.
	var conf = initConfig()
	model.InitDatabase()
	// Start the http server.
	var app = setupRoutes()
	var port = conf.SITE_PORT
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--install":
			install()
			os.Exit(0)
		case "--test":
			P_TestData()
			os.Exit(0)
		case "--reset":
			deleledb()
			os.Exit(0)
		case "--help":
			fmt.Print(helpStr)
			os.Exit(0)
		case "--geoip":
			ip := os.Args[2]
			fmt.Printf("%s\n", utils.GetIPCountry(ip))
			os.Exit(0)
		case "--md2html":
			inputpath := os.Args[2]
			fmt.Printf("%s\n", utils.MarkDownFile2HTML(inputpath))
			os.Exit(0)
		case "-p":
			port = os.Args[2]
		case "--pool":
			fmt.Printf("starting test the mining pool...")
		}
	}
	utils.LogD("Starting server at:" + port)
	app.Logger.Fatal(app.Start(":" + port))
}
