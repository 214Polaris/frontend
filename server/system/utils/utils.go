package utils

import (
	"crypto/sha512"
	"fmt"
	"time"
)

// 分页
func Page(Limit, Page int) (limit, offset int) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

// 排序
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}
	return sort
}

// 获取时间
const TimeLayout = "2006-01-02 15:04:05"

var (
	Local = time.FixedZone("CST", 8*3600)
)

func GetNow() string {
	now := time.Now().In(Local).Format(TimeLayout)
	return now
}

func TimeFormat(s string) string {
	result, err := time.ParseInLocation(TimeLayout, s, time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return result.In(Local).Format(TimeLayout)
}

// 密码加密
func Md5(password string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(password)))
}
