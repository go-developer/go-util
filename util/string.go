// Package util ...
//
// File : string.go
//
// Decs : 字符串操作工具集
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:36:23
package util

import (
	"math/rand"
	"strings"
	"time"
)

// StringUtil 字符串工具类实例
//
// Author : go_developer@163.com<张德满>
var StringUtil *stringUtil

func init()  {
	StringUtil = &stringUtil{}
}

// stringUtil 字符串操作工具类
//
// Author : go_developer@163.com<张德满>
type stringUtil struct {

}

// GenRandomString 生成指定长度的随机字符串
//
// source 不为空时,将以此作为字符池
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:29:41
func (s *stringUtil) GenRandomString(source string, length uint) string {
	if length == 0 {
		return ""
	}
	if len(source) == 0 {
		//字符串为空，默认字符源为如下:
		source = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	strByte := []byte(source)
	var genStrByte = make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < int(length); i++ {
		genStrByte = append(genStrByte, strByte[r.Intn(len(strByte))])
	}
	return string(genStrByte)
}

// 将按照split分割的字符串转为小驼峰
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:30:55
func (s *stringUtil) ToHump(str string, split string) string {
	strArr := strings.Split(str, split)
	out := strArr[0]
	for i := 1; i < len(strArr); i++ {
		out = out + s.Capitalize(strArr[i])
	}
	return out
}

// 实现首字母大写
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:31:09
func (s *stringUtil) Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	strArr := strings.Split(str, "")
	return strings.ToUpper(strArr[0]) + strings.Join(strArr[1:], "")
}

