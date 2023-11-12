package utils

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

// 判断是否为有效的邮箱
func IsEmail(email string) bool {
	// 判断是否为邮箱
	pattern := "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	if !Match(pattern, email) {
		return false
	}
	return true
}

// 判断是否为有效的手机号
func IsPhone(phone string) bool {
	// 手机号，11位数字，1开头，2-10位为0-9的数字，11位为1-9
	pattern := "^1[0-9]{9}[1-9]$"
	if !Match(pattern, phone) {
		return false
	}
	return true
}

// 判断是否为有效的账号
func IsAccount(account string) bool {
	// 账号，10-20位，字母开头，允许字母数字下划线
	pattern := "^[a-zA-Z][a-zA-Z0-9_]{9,19}$"
	if !Match(pattern, account) {
		return false
	}
	return true
}

// 判断是否为有效的密码
func IsPassword(password string) bool {
	//  10-20位，字母开头，允许字母数字下划线
	pattern := "^[a-zA-Z][a-zA-Z0-9_]{9,19}$"
	if !Match(pattern, password) {
		return false
	}
	return true
}

func Match(pattern string, str string) bool {
	return Regexp(pattern).MatchString(str)
}

func Regexp(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func ValidUserAuthority(context *gin.Context) (userId int32, err error) {
	id := context.Param("id")
	userId, ok := context.MustGet("userId").(int32)
	if ok {
		if id != fmt.Sprintf("%v", userId) {
			return 0, fmt.Errorf("你无权修改此账号的信息")
		}
	} else {
		return 0, fmt.Errorf("获取身份信息失败，请重新登录")
	}
	return userId, nil
}
