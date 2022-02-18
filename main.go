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
	model.AddNewUser(ADMIN_UNAME, ADMIN_PWD, ADMIN_EMAIL, "admin", ADMIN_AVATAR, "")
}

func test() {
	// test user
	model.AddNewUser("test1", "Fuckpwd", "test@example.com", "dev", "http://ava.tar", "descrityion1")
	model.AddNewUser("test2", "Fuck", "test@example.com", "dev", "http://ava.tar", "descrityion1")
	users := model.GetAllUsers()
	utils.LogD("%v", users)
	for _, u := range users {
		u.Udesc = "new desc here!"
		model.UpdateUser(u)
	}
	utils.LogD("%v", model.GetAllUsers())
	for _, u := range users {
		model.DeleteUser(u.Uid)
	}
	utils.LogD("%v", model.GetAllUsers())
	// test preference
	model.AddNewPref("pref1", "this is content of pref1")
	model.AddNewPref("pref2", "this is content of pref2")
	model.AddNewPref("pref5", 234)
	model.UpdatePref("pref5", 7890)
	utils.LogD("%v", model.GetAllPrefs())
	utils.LogD("pref5: %v", model.GetPref("pref5"))
	model.DeletePref("pref5")
	model.DeletePref("pref1")
	model.DeletePref("pref2")
	utils.LogD("%v", model.GetAllPrefs())
	// test channel
	model.AddNewChannel("channel1", "http://cpiurl", "https://urlforcover", "descrption", "issuerID1")
	model.AddNewChannel("channel2", "http://cpiurl", "https://urlforcover", "descrption", "issuerID2")
	c := model.GetAllChannels()
	utils.LogD("%v", c)
	for _, u := range c {
		u.ChanDesc = "new  Channel desc here!"
		model.UpdateChannel(u)
	}
	utils.LogD("%v", model.GetAllChannels())
	for _, u := range c {
		model.DeleteChannel(u.ChanId)
	}
	utils.LogD("%v", model.GetAllChannels())
	// test posts
	// test adding inde page
	model.AddNewIndepage("about", "Aboutpage title", "## This is H2", "issuerID1")
	model.AddNewIndepage("news", "News title", "# This is H1", "issuerID2")
	pg1 := model.GetIndepages()
	utils.LogD("%v", pg1)
	for _, u := range pg1 {
		u.Ptitle = "new  page title here!"
		model.UpdateIndepage(u)
	}
	utils.LogD("%v", model.GetIndepages())
	pg1 = model.GetIndepages()
	for _, u := range pg1 {
		model.DeleteIndepage(u.Pname)
	}
	utils.LogD("%v", model.GetIndepages())
	// Redis test
	utils.LogD("%v", model.GetCacheRedis("test catch", "none"))
	model.SetCacheRedis("test catch", "this is update for cache!!!", 0)
	utils.LogD("%v", model.GetCacheRedis("test catch", "none"))
	model.DelCacheRedis("test catch")
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
	var r = setupRoutes()
	model.InitDatabase()
	// Start the http server.
	var port = SITE_PORT
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--install":
			install()
			os.Exit(0)
		case "--test":
			test()
			os.Exit(0)
		case "--reset":
			deleledb()
			os.Exit(0)
		case "--help":
			fmt.Printf(helpStr)
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
	r.Logger.Fatal(r.Start(":" + port))
}
