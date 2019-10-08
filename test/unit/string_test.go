package unit

import (
	"go-util/util"
	"testing"
)

/**
 * 测试字符串按照指定分隔符转化为小驼峰的性能
 * @author go_developer@163.com
 */
func TestToHump(t *testing.T) {
	var testTable = []struct {
		In   string
		Flag string
		Out  string
	}{
		{"zhangdeman", "_", "zhangdeman"},     //不存在指定分隔符的case
		{"_zhangdeman", "_", "Zhangdeman"},    //以分隔符开头的字符串
		{"____zhangdeman", "_", "Zhangdeman"}, //开头连续若干个分隔符
		{"zhang_de_man", "_", "zhangDeMan"},   //常见的分隔方式
	}
	for _, testCase := range testTable {
		result := util.StringUtil.ToHump(testCase.In, testCase.Flag)
		if testCase.Out != result {
			t.Fatalf("toHump方法执行异常, in = %s, flag = %s, except = %s, real = %s", testCase.In, testCase.Flag, testCase.Out, result)
		}

	}
}
