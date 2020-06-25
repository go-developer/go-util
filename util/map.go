package util

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MapUtil map操作工具集
//
// Author : go_developer@163.com<张德满>
var MapUtil *mapUtil

func init()  {
	MapUtil = &mapUtil{}
}

// mapUtil map操作的相关方法集合
//
// Author : go_developer@163.com<张德满>
type mapUtil struct {

}

// ToStruct map转为结构体的方法
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:16:36
func (m *mapUtil) ToStruct(mapData interface{}, structData interface{}) error {
	var err error
	var jsonByteData []byte
	if jsonByteData, err = json.Marshal(mapData); nil != err {
		return errors.New("map数据json序列化失败 ==> " + err.Error())
	}
	if err = json.Unmarshal(jsonByteData, structData); nil != err {
		return errors.New("数据应设置结构体失败 ==> " + err.Error())
	}
	return nil
}

// 安全的从map中获取一个key的值
//
// 需要是接受结果数据类型的数据指针
//
// 当读取的key不存在或者数据类型不存在, 会抛异常
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:20:39
func (m *mapUtil) SafeGetValue(data map[string]interface{}, key string, value interface{}) error {
	if nil == data {
		return errors.New("map为nil")
	}
	var (
		val interface{}
		exist bool
	)
	if val, exist = data[key]; !exist {
		return fmt.Errorf("key : %s 不存在")
	}

	return ConvertAssign(value, val)
}