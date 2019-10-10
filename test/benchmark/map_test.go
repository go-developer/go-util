package benchmark

import (
	"go-util/util"
	"testing"
)

/**
 * map转struct的基准性能测试
 * @author go_developer@163.com
 */
func BenchmarkToStruct(b *testing.B)  {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	b.StartTimer() //重新开始时间
	mapData := map[string]string{
		"name": "zhangdeman",
		"age": "23",
		"height": "180",
	}
	type result struct{
		Name string `json:"name"`
		Age string `json:"age"`
		Height string `json:"height"`
	}

	for i := 0; i < b.N; i++ {
		var r  result
		util.MapUtil.ToStruct(mapData, &r)
	}

}
