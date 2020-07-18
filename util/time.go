// Package util ...
//
// File : time.go
//
// Decs : 时间操作工具集
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:35:42
package util

import (
	"errors"
	"time"
)

// TimeUtil 实践操作工具集实例
//
// Author : go_developer@163.com<张德满>
var TimeUtil *timeUtil

func init() {
	TimeUtil = &timeUtil{}
}

// timeUtil 时间工具助手
//
// Author : go_developer@163.com<张德满>
type timeUtil struct {
}

// GetCurrentFormatTime 获取当前时间戳的格式化时间
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:32:47
func (t *timeUtil) GetCurrentFormatTime() string {
	return t.GetFormatTime(time.Now().Unix())
}

// GetFormatTime 获取时间戳的格式化时间
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:33:03
func (t *timeUtil) GetFormatTime(second int64) string {
	return time.Unix(second, 0).Format("2006-01-02 15:04:05")
}

// GetUnixTime 根据格式化的时间获取对应的时间戳
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:34:09
func (t *timeUtil) GetUnixTime(formatTime string) (unixTime int64, err error) {

	var (
		timeLayout = "2006-01-02 15:04:05" //转化所需模板
		loc        *time.Location          //时区
		timeInfo   time.Time               //时间信息
	)

	//获取时区失败
	if loc, err = time.LoadLocation("Local"); nil != err {
		return 0, errors.New("获取时区失败 ==> " + err.Error())
	}
	if timeInfo, err = time.ParseInLocation(timeLayout, formatTime, loc); nil != err {
		return 0, errors.New("解析时间失败 ==> " + err.Error())
	}
	unixTime = timeInfo.Unix()
	return unixTime, nil
}

// GetFormatCurrentNanoTime 获取格式化的纳秒时间戳
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 00:22:39
func (t *timeUtil) GetFormatCurrentNanoTime(nano int64) string {
	second := nano / 1e9
	leave := nano - second*1e9
	return time.Unix(second, leave).String()
}
