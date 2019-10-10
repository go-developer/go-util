package unit

import (
	"go-util/util"
	"testing"
)

/**
 * 测试map转结构体的功能
 * @author go_developer@163.com
 */
func TestToStruct(t *testing.T)  {
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
	var r  result
	if err := util.MapUtil.ToStruct(mapData, &r); nil != err {
		t.Fatalf("toStruct测试未通过, 错误信息 [ %s ]", err.Error())
	}
}
