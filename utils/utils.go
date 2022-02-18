// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/gomarkdown/markdown"
	"github.com/phuslu/iploc"
)

// MD5Bytes encodes string to MD5 bytes.
func MD5Bytes(str string) []byte {
	m := md5.New()
	m.Write([]byte(str))
	return m.Sum(nil)
}

// MD5 encodes string to MD5 hex value.
func MD5(str string) string {
	return hex.EncodeToString(MD5Bytes(str))
}

// SHA1 encodes string to SHA1 hex value.
func SHA1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// ShortSHA1 truncates SHA1 string length to at most 10.
func ShortSHA1(sha1 string) string {
	if len(sha1) > 10 {
		return sha1[:10]
	}
	return sha1
}

// SHA256 encodes string to SHA256 hex value.
func SHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Convert IP address to country, eg: CN\US\HK\TW
func GetIPCountry(ipaddr string) string {
	b_str := iploc.Country(net.ParseIP(ipaddr))
	return string(b_str)
	//return iploc.Country(net.IP{119, 23, 83, 65})
}

// convert markdown text to HTML
func MarkDown2HTML(input string) string {
	md := []byte(input)
	output := markdown.ToHTML(md, nil, nil)
	return string(output)
}

// read markdown from input file, return HTML string
func MarkDownFile2HTML(input_path string) string {
	file, err := os.Open(input_path)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(file)
	output := markdown.ToHTML(b, nil, nil)
	return string(output)
}

// BasicAuthDecode decodes username and password portions of HTTP Basic Authentication
// from encoded content.
func BasicAuthDecode(encoded string) (string, string, error) {
	s, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", "", err
	}

	auth := strings.SplitN(string(s), ":", 2)
	return auth[0], auth[1], nil
}

// BasicAuthEncode encodes username and password in HTTP Basic Authentication format.
func BasicAuthEncode(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

const alphanum = "0123456789abcdefghijklmnopqrstuvwxyz"

// RandomString returns generated random string in given length of characters.
// It also returns possible error during generation.
func RandomString(n int) (string, error) {
	buffer := make([]byte, n)
	max := big.NewInt(int64(len(alphanum)))

	for i := 0; i < n; i++ {
		index, err := randomInt(max)
		if err != nil {
			return "", err
		}

		buffer[i] = alphanum[index]
	}

	return string(buffer), nil
}
func RandomNum(max *big.Int) (uint, error) {
	return randomInt(max)
}

func randomInt(max *big.Int) (uint, error) {
	rand, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return uint(rand.Int64()), nil
}

// HashEmail hashes email address to MD5 string.
// https://en.gravatar.com/site/implement/hash/
func HashEmail(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	h := md5.New()
	h.Write([]byte(email))
	return hex.EncodeToString(h.Sum(nil))
}

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func computeTimeDiff(diff int64) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "now"
	case diff < 2:
		diff = 0
		diffStr = "1 second"
	case diff < 1*Minute:
		diffStr = fmt.Sprintf("%d seconds", diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = "1 minute"
	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d minutes", diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = "1 hour"
	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d hours", diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = "1 day"
	case diff < 1*Week:
		diffStr = fmt.Sprintf("%d days", diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = "1 week"
	case diff < 1*Month:
		diffStr = fmt.Sprintf("%d weeks", diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = "1 month"
	case diff < 1*Year:
		diffStr = fmt.Sprintf("%d months", diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = "1 year"
	default:
		diffStr = fmt.Sprintf("%d years", diff/Year)
		diff = 0
	}
	return diff, diffStr
}

// Check if a time value is greater than time.Now()
func IsTimeExpired(ptime int64) bool {
	now := time.Now()
	diff := now.Unix() - ptime
	LogD("IsTimeExpired? %v diff: %v\n", (diff >= 0), diff)
	return diff >= 0
}

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiff(diff)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

// Subtract deals with subtraction of all types of number.
func Subtract(left interface{}, right interface{}) interface{} {
	var rleft, rright int64
	var fleft, fright float64
	var isInt bool = true
	switch left.(type) {
	case int:
		rleft = int64(left.(int))
	case int8:
		rleft = int64(left.(int8))
	case int16:
		rleft = int64(left.(int16))
	case int32:
		rleft = int64(left.(int32))
	case int64:
		rleft = left.(int64)
	case float32:
		fleft = float64(left.(float32))
		isInt = false
	case float64:
		fleft = left.(float64)
		isInt = false
	}

	switch right.(type) {
	case int:
		rright = int64(right.(int))
	case int8:
		rright = int64(right.(int8))
	case int16:
		rright = int64(right.(int16))
	case int32:
		rright = int64(right.(int32))
	case int64:
		rright = right.(int64)
	case float32:
		fright = float64(left.(float32))
		isInt = false
	case float64:
		fleft = left.(float64)
		isInt = false
	}

	if isInt {
		return rleft - rright
	} else {
		return fleft + float64(rleft) - (fright + float64(rright))
	}
}

// EllipsisString returns a truncated short string,
// it appends '...' in the end of the length of string is too large.
func EllipsisString(str string, length int) string {
	if len(str) < length {
		return str
	}
	return str[:length-3] + "..."
}

// TruncateString returns a truncated string with given limit,
// it returns input string if length is not reached limit.
func TruncateString(str string, limit int) string {
	if len(str) < limit {
		return str
	}
	return str[:limit]
}

// IsLetter reports whether the rune is a letter (category L).
// https://github.com/golang/go/blob/master/src/go/scanner/scanner.go#L257
func IsLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}

// IsTextFile returns true if file content format is plain text or empty.
func IsTextFile(data []byte) bool {
	if len(data) == 0 {
		return true
	}
	return strings.Contains(http.DetectContentType(data), "text/")
}

func IsImageFile(data []byte) bool {
	return strings.Contains(http.DetectContentType(data), "image/")
}

func IsPDFFile(data []byte) bool {
	return strings.Contains(http.DetectContentType(data), "application/pdf")
}

func IsVideoFile(data []byte) bool {
	return strings.Contains(http.DetectContentType(data), "video/")
}

const (
	Byte  = 1
	KByte = Byte * 1024
	MByte = KByte * 1024
	GByte = MByte * 1024
	TByte = GByte * 1024
	PByte = TByte * 1024
	EByte = PByte * 1024
)

var bytesSizeTable = map[string]uint64{
	"b":  Byte,
	"kb": KByte,
	"mb": MByte,
	"gb": GByte,
	"tb": TByte,
	"pb": PByte,
	"eb": EByte,
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+" %s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}

func GenNewPasswd(pwd string, salt uint) string {
	return SHA256(pwd + fmt.Sprint(salt))
}

func GenNewSalt() uint {
	r, _ := RandomNum(big.NewInt(int64(3141592653)))
	return r
}

func ComparePasswd(pwd1 string, pwd2 string, salt uint) bool {
	b1 := SHA256(pwd1 + fmt.Sprint(salt))
	b2 := SHA256(pwd2 + fmt.Sprint(salt))
	return b1 == b2
}

// generate an SHA-1 hashed string
func GenUid() string {
	r, _ := RandomString(40)
	return SHA1(r)
}

// Generate an 40 letters random string
func GenPid() string {
	r, _ := RandomString(40)
	return r
}

// eg: views/index.html ,views/subtmpl/subtmpl.html
// tmpl:=LoadTmpl("./views",funcs)
// tmpl.ExecuteTemplate(yourWriter, "index.html", yourData)、tmpl.ExecuteTemplate(yourWriter, "subtmpl/subtmpl.html", yourData)
func LoadHTMLTmpl(root string, funcs template.FuncMap) (*template.Template, error) {
	tmpl := template.New("LoadTmpl")
	if funcs != nil {
		tmpl.Funcs(funcs)
	}
	rln := len(root)
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		pln := len(path)
		// file ends with .html
		if !fi.IsDir() && pln > 4 && path[pln-5:] == ".html" {
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			name := strings.ReplaceAll(path, "\\", "/")
			_, err = tmpl.New(name[rln-1:]).Parse(string(buf))
			if err != nil {
				return err
			}
		}
		return nil
	})
	return tmpl, err
}
