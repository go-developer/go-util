package wrapper

import "fmt"

/**
 * 定义string的包装类型
 * @author zhangdeman001@ke.com
 */
type String string

/**
 * 获取字符串的值
 * @author go_developer@163.com
 */
func (s *String) GetValue() string {
	return fmt.Sprintf("%s", *s)
}

/**
 * 在字符串末尾追加字符串
 * @param string appendStr 要追加的字符串
 * @author go_developer@163.com
 */
func (s *String) Append(appendStr string) *String {
	var str = String(s.GetValue() + appendStr)
	*s = str
	return s
}

/**
 * 获取字符串的长度
 * @author go_developer@163.com
 */
func (s *String) Length() int {
	return len(s.GetValue())
}

/**
 * 判断字符串是否相等
 * @author go_developer@163.com
 */
func (s *String) Equal(str string) bool {
	return s.GetValue() == str
}
