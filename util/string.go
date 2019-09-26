package util

import (
	"math/rand"
	"strings"
	"time"
)

var StringUtil *stringUtil

func init()  {
	StringUtil = &stringUtil{}
}

/**
 * 字符串操作工具类
 * @author go_developer@163.com
 */
type stringUtil struct {

}

/**
 * 生成指定长度的随机字符串
 * @param string source 从哪些字符里选取字符
 * @param uint length 生成的随机字符串长度
 * @return string 生成的随机字符串
 * @author go_developer@163.com
 */
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

/**
 * 将按照split分割的字符串转为小驼峰
 * @param string str 原字符串
 * @param string split 字符串的分隔符
 * @author go_developere@163.com
 */
func (s *stringUtil) ToHump(str string, split string) string {
	strArr := strings.Split(str, split)
	out := strArr[0]
	for i := 1; i < len(strArr); i++ {
		out = out + s.Capitalize(strArr[i])
	}
	return out
}

/**
 * 实现首字母大写
 * @param string str 要处理的字符串
 * @return string 首字母大写后的字符串
 * @author zhangdeman001@ke.com
 */
func (s *stringUtil) Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	strArr := strings.Split(str, "")
	return strings.ToUpper(strArr[0]) + strings.Join(strArr[1:len(strArr)], "")
}

