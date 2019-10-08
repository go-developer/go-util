package benchmark

import (
	"go-util/util"
	"testing"
)

/**
 * 测试字符串按照指定分隔符转化为小驼峰的性能
 * @author go_developer@163.com
 */
func BenchmarkToHump(b *testing.B) {

	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作，例如读取文件数据，数据库连接之类的，
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		util.StringUtil.ToHump("zhang_de_man", "_")
	}
}

/**
 * 字符串首字母大写
 * @author go_developer@163.com
 */
func BenchmarkCapitalize(b *testing.B) {

	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作，例如读取文件数据，数据库连接之类的，
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		util.StringUtil.Capitalize("zhangdeman")
	}
}

