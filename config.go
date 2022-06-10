package main

import (
	"encoding/json"
	"io/ioutil"
	"tinfo-go/utils"
)

var IS_DEV = true

type SiteConf struct {
	SITE_NAME    string
	SITE_VERSION string
	SITE_DESC    string
	SITE_PORT    string
	ADMIN_UNAME  string
	ADMIN_PWD    string
	ADMIN_EMAIL  string
	ADMIN_AVATAR string
}

func initConfig() *SiteConf {
	//load  from json file
	var conf = SiteConf{}
	content, err := ioutil.ReadFile("./site_conf.json")
	if err != nil {
		utils.LogFatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &conf)
	if err != nil {
		utils.LogFatal("Error during Unmarshal(): ", err)
	}
	utils.LogD("Website conf loaded: %v", conf)
	return &conf
}
