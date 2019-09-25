package util

import (
	"errors"
	"os"
)

var FileUtil *fileUtil

func init() {
	FileUtil = &fileUtil{}
}

/**
 * 文件操作助手
 * @author go_developer@163.com
 */
type fileUtil struct {
}

/**
 * 获取文件信息， 此方法也可判断文件是否存在
 * @param string filePath 文件路径
 * @return *os.FileInfo 文件的信息
 * @return error 获取文件信息是的异常信息
 */
func (f *fileUtil) GetFileInfo(filePath string) (*os.FileInfo, error) {
	var (
		err      error
		fileInfo os.FileInfo
	)
	if fileInfo, err = os.Stat(filePath); nil != err {
		return nil, errors.New("获取文件信息异常 ==> " + err.Error())
	}
	return &fileInfo, nil
}
