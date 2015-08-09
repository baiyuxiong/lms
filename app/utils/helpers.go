package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/revel/revel"
	"io"
	"math/rand"
	"time"
)

type TrackResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Flow struct {
	Id    string         `json:"id"`
	Name  string         `json:"name"`
	Steps map[string]Step       `json:"steps"`
}

type Step struct {
	Charge string `json:"charge"`
	Type string `json:"type"`
}

const (
	Alnum = iota
	Alpha
	Numeric
)



var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphas = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")

func RandString(stringType, n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	if stringType == Alnum {
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
	} else if stringType == Numeric {
		for i := range b {
			b[i] = numbers[rand.Intn(len(numbers))]
		}
	}else if stringType == Alpha {
		for i := range b {
			b[i] = alphas[rand.Intn(len(alphas))]
		}
	}
	return string(b)
}

func Salt() string {
	return RandString(Alnum, 16)
}

func StringInSlice(a string, list []string) bool {
	//revel.INFO.Println(a,list)
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func FormatNow(format string) string {
	return time.Now().Format(format)
}

// 加密密码,转成md5
func EncryptPassword(salt, password string) string {
	return Md5String(password+salt)
}

func Md5String(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func ValidationErrorToString(e []*revel.ValidationError) string {
	if nil != e {
		return e[0].Message
	}
	return ""
}

func Pagination(limit, offset int, totalCount int64) string {

	pageStart := int64(offset+1)
	pageEnd := int64(offset+limit)
	if pageEnd > totalCount {
		pageEnd = totalCount
	}

	currentPage := offset/limit+1
	prevPage := "<li class=\"prev disabled\"><a href=\"#\"><-上一页</a></li>";
	if currentPage>1 {
		prevPage = "<li class=\"prev\"><a href=\"/admin/ListUser?offset="+fmt.Sprintf("%d", offset-limit)+"\"><-上一页</a></li>";
	}

	nextPage := "<li class=\"next disabled\"><a href=\"#\">下一页-> </a></li>"
	nextOffset := int64(offset + limit)
	if nextOffset < totalCount {
		nextPage = "<li class=\"next\"><a href=\"/admin/ListUser?offset="+fmt.Sprintf("%d", nextOffset)+"\">下一页-> </a></li>"
	}

	pages := "<div class=\"row\"><div class=\"col-md-6\"><div class=\"dataTables_info\" id=\"example_info\">显示 "
	pages += fmt.Sprintf("%d", pageStart) + " 到 " + fmt.Sprintf("%d", pageEnd) + " 条，共 " + fmt.Sprintf("%d", totalCount) + " 条</div></div>"
	pages += "<div class=\"col-md-6\"><div class=\"dataTables_paginate paging_bootstrap\"><ul class=\"pagination\">"
	pages += prevPage
	pages +="<li class=\"active\"><a href=\"#\">"+fmt.Sprintf("%d", currentPage)+"</a></li>"
	pages += nextPage+"</ul></div></div></div>"


	return pages
}


