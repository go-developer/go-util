package util

import (
	"errors"
	"time"
)

var TimeUtil *timeUtil

func init() {
	TimeUtil = &timeUtil{}
}

/**
 * 时间工具助手
 * @author go_developer@163.com
 */
type timeUtil struct {
}

/**
 * 获取当前时间戳的格式化时间
 * @author go_developer@163.com
 */
func (t *timeUtil) GetCurrentFormatTime() string {
	return t.GetFormatTime(time.Now().Unix())
}

/**
 * 获取时间戳的格式化时间
 * @param int64 second 要格式化的时间戳，单位秒
 * @author go_developer@163.com
 */
func (t *timeUtil) GetFormatTime(second int64) string {
	return time.Unix(second, 0).Format("2006-01-02 15:04:05")
}

/**
 * 根据格式化的时间获取对应的时间戳
 * @param string formatTime 格式化的时间，eg: 2006-01-02 15:04:05
 * @return int64 unixTime 转化后的时间戳，若转化过程中出现异常,辞职为0
 * @return error err 时间戳转化过程中的异常信息,存在两种可能: 1. 获取本地时区异常 2. 格式化时间戳解析异常
 * @author go_developer@163.com<张德满>
 */
func (t *timeUtil) GetUnixTime(formatTime string) (unixTime int64, err error) {

	var (
		timeLayout = "2006-01-02 15:04:05" //转化所需模板
		loc        *time.Location //时区
		timeInfo   time.Time      //时间信息
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
