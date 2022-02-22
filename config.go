package main

import (
	"encoding/json"
	"io/ioutil"
	"tinfo-go/utils"
)

// tokens
var TOKEN_DEFAULT_LIFESPAN int64 = 3600 * 24 * 1
var TOKEN_SESSION_MAXAGE int = 3600 * 24 * 1
var TOKEN_SESSION_SECRET []byte = []byte("e5ba34c907f46f6d8dae11ece24004f48")

// sessions
var SESSION_SECRET = "83aae5ba34c907f46f6d8d"
var SESSION_NAME = "8966"

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
