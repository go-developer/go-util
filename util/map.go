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

/**
 * 获取一个key类型为int的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetUintKeyList(data map[uint]interface{}) []uint {
	keyList := make([]uint, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int8的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetUint8KeyList(data map[uint8]interface{}) []uint8 {
	keyList := make([]uint8, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int16的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetUint16KeyList(data map[uint16]interface{}) []uint16 {
	keyList := make([]uint16, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int8的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetUint32KeyList(data map[uint32]interface{}) []uint32 {
	keyList := make([]uint32, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int8的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetUint64KeyList(data map[uint64]interface{}) []uint64 {
	keyList := make([]uint64, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetIntKeyList(data map[int]interface{}) []int {
	keyList := make([]int, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int8的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetInt8KeyList(data map[int8]interface{}) []int8 {
	keyList := make([]int8, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int16的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetInt16KeyList(data map[int16]interface{}) []int16 {
	keyList := make([]int16, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int8的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetInt32KeyList(data map[int32]interface{}) []int32 {
	keyList := make([]int32, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为int8的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetInt64KeyList(data map[int64]interface{}) []int64 {
	keyList := make([]int64, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为float32的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetFloat32KeyList(data map[float32]interface{}) []float32 {
	keyList := make([]float32, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为float32的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetFloat64KeyList(data map[float64]interface{}) []float64 {
	keyList := make([]float64, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}

/**
 * 获取一个key类型为string的map的全部key
 * @param data 输入的map
 * @author go_developer@163.com
 */
func (m *mapUtil) GetStringKeyList(data map[string]interface{}) []string {
	keyList := make([]string, len(data), len(data))
	for key, _ := range data{
		keyList = append(keyList, key)
	}
	return keyList
}


