package unit

import (
	"go-util/util"
	"testing"
)

/**
 * 测试字符串按照指定分隔符转化为小驼峰的功能测试
 * @author go_developer@163.com
 */
func TestToHump(t *testing.T) {
	var testTable = []struct {
		In   string
		Flag string
		Out  string
	}{
		{"zhangdeman", "_", "zhangdeman"},           //不存在指定分隔符的case
		{"_zhangdeman", "_", "Zhangdeman"},          //以分隔符开头的字符串
		{"____zhangdeman", "_", "Zhangdeman"},       //开头连续若干个分隔符
		{"zhang_de_man", "_", "zhangDeMan"},         //常见的分隔方式
		{"zhang___de___man", "_", "zhangDeMan"},     //中间连续若干个分隔符的case
		{"zhang_ __de__ _man", "_", "zhang De Man"}, //分隔符之间只包含空格的case
	}
	for _, testCase := range testTable {
		result := util.StringUtil.ToHump(testCase.In, testCase.Flag)
		if testCase.Out != result {
			t.Fatalf("toHump方法执行异常, in = %s, flag = %s, except = %s, real = %s", testCase.In, testCase.Flag, testCase.Out, result)
		}
	}
}

/**
 * 字符串首字母大写功能测试
 * @author go_developer@163.com
 */
func TestCapitalize(t *testing.T) {
	var testTable = []struct {
		In  string
		Out string
	}{
		{"zhangdeman", "Zhangdeman"},   //普通字符串
		{"$zhangdeman", "$zhangdeman"}, //特殊字符开头的开头的字符串
		{" zhangdeman", " zhangdeman"}, //空格开头的字符串
		{"", ""},                       //空字符串
	}
	for _, testCase := range testTable {
		result := util.StringUtil.Capitalize(testCase.In)
		if testCase.Out != result {
			t.Fatalf("Capitalize方法执行异常, in = %s, except = %s, real = %s", testCase.In, testCase.Out, result)
		}
	}
}
