## tinfo-go

This project is a implemention of Tinfo, this project is wirten in golang.

There is 2 main part of this program:
1. front-end. see tinfo-ui project
2. backends, mostly json apis


/*
This is a growing project. copyright, @ 2019. Jay Liu.

## CLI API:

please use --help to see related documents

## HTTP APIs:

### Session Management:
----------------

HTTP GET /ssid/new
	Description: Generate a new SSID
	Params: None
	Return: json: {ssid": ssid,	"sskey": sskey, "expire": SSID_TIMEOUT"}

HTTP GET /ssid/check
	Description: Check if ssid is valid
	Params: ssid, sskey
	Return: json; if success: {"status": "matched",} if fail: {"status": "mismatched",}

HTTP GET /ssid/close
	Description: Writeoff the session
	Params: ssid, sskey
	Return: json; if success: {"status": "closed",} if fail: {"status": "mismatched",}

### Captcha Management
----------------

HTTP GET /captcha/new
	Description: Get a  captcha image
	Params: cid
	Return: if provide cid, A png format picture, otherwise will return a new cid and captcha address and timeout

HTTP GET /captcha/verify
	Description: verify if captcha match
	Params: cid, cstr
	Return: json; if matched {"status": "matched",} other {"status": "captcha mismatch!",}

### User Management
----------------

HTTP POST /admin/login
	Description: Check if admin can login
	Params: ssid, cstr, username, password
	Return: json; if sucess {"status":"ok" } else {"status": "denied"}

HTTP POST /admin/changepwd
	Description: change password for admin
	Params: ssid, username, newpwd
	Return: json; ok {"status":"ok"} else {"status": "denied"}

HTTP GET /admin/logout
	Description: logout admin
	Params: ssid, username
	Return: json; ok {"status":"ok"} else {"status": "denied"}

### IP tools
-----------------
HTTP GET /iptool/country
	Description: look for IP's geo location, country name
	Params: ip
	Return: json; ok {"country":"US"}

###
*/