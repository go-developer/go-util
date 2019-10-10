package unit

import (
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
	var result struct{
		Name string `json:"name"`
		Age string `json:"age"`
		Height string `json:"height"`
	}
	util.MapUtil
}
