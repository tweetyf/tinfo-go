package main

// site info
var SITE_NAME = "Tinfo-go"
var SITE_VERSION = "0.0.1"
var SITE_DES = "A better future."
var SITE_PORT = "8080"

// admins user info
var ADMIN_UNAME = "admin"
var ADMIN_PWD = "admin"
var ADMIN_EMAIL = "admin@example.com"
var ADMIN_AVATAR = "/static/images/h/145848.png"

// tokens
var TOKEN_DEFAULT_LIFESPAN int64 = 3600 * 24 * 1
var TOKEN_SESSION_MAXAGE int = 3600 * 24 * 1
var TOKEN_SESSION_SECRET []byte = []byte("e5ba34c907f46f6d8dae11ece24004f48")

// sessions
var SESSION_SECRET = "83aae5ba34c907f46f6d8d"
var SESSION_NAME = "8966"

var IS_DEV = true
