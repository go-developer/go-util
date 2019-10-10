package util

import (
	"encoding/json"
	"errors"
)

var MapUtil *mapUtil

func init()  {
	MapUtil = &mapUtil{}
}

/**
 * map操作的相关方法集合
 * @author go_developer@163.com
 */
type mapUtil struct {

}

/**
 * map转为结构体的方法
 * @param mapData 原始的map数据
 * @param 解析后的数据结构体,应该是一个结构体指针
 * @return error map转结构体的错误信息
 */
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
