package main

import (
	"tinfo-go/model"
	"tinfo-go/utils"
)

func P_TestData() {
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
