package util

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// FileUtil 文件操作工具集
//
// Author : go_developer@163.com<张德满>
var FileUtil *fileUtil

func init() {
	FileUtil = &fileUtil{}
}

// DirFile 获取文件夹下的文件列表
//
// Author : go_developer@163.com<张德满>
type DirFile struct {
	FullFilePath string    //文件的全路径
	FileType     string    //文件类型
	FileName     string    //文件名，不包含路径
	IsDir        bool      //是否是目录
	FileList     []DirFile //当文件为目录，递归查询目录下的文件
}

// fileUtil 文件操作助手
//
// Author : go_developer@163.com<张德满>
type fileUtil struct {
}

// GetFileInfo 获取文件信息， 此方法也可判断文件是否存在
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:13:34
func (f *fileUtil) GetFileInfo(filePath string) (os.FileInfo, error) {
	var (
		err      error
		fileInfo os.FileInfo
	)
	if fileInfo, err = os.Stat(filePath); nil != err {
		return nil, errors.New("获取文件信息异常 ==> " + err.Error())
	}
	return fileInfo, nil
}

// IsDir 获取一个文件路径是否为目录
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:14:38
func (f *fileUtil) IsDir(filePath string) (bool, error) {
	fileInfo, err := f.GetFileInfo(filePath)
	if nil != err {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// GetDirFileList 递归获取目录下文件,因为是递归实现,请慎重评估depth深度
//
// 询过程中的异常信息,两种场景: 1. 目录不存在 2. 传进来的路径不是一个目录 3. 路径为空
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:15:13
func (f *fileUtil) GetDirFileList(dirPath string, depth uint) ([]DirFile, error) {
	if depth == 0 {
		return make([]DirFile, 0), nil
	}

	dirPath = strings.Trim(dirPath, " ")
	if len(dirPath) == 0 {
		return nil, errors.New("目录路径不能为空")
	}

	if !strings.HasSuffix(dirPath, "/") {
		//目录路径的结尾不是 "/",补充进去，保证路径格式的统一
		dirPath = dirPath + "/"
	}

	var (
		err error
		fileInfo os.FileInfo
		dirFileList = make([]DirFile, 0)
	)

	if fileInfo, err = f.GetFileInfo(dirPath); nil != err {
		return nil, err
	}

	if !fileInfo.IsDir() {
		return nil, errors.New(dirPath + " 不是一个目录")
	}

	fullFileList, err := ioutil.ReadDir(dirPath)

	if nil != err {
		return nil, errors.New("读取目录下文件失败")
	}

	for _, itemFile := range fullFileList {
		fullFilePath := dirPath + itemFile.Name()
		fileNameArr := strings.Split(itemFile.Name(), ".")
		tmpDirFile := DirFile{
			FullFilePath: fullFilePath,
			FileType:     fileNameArr[len(fileNameArr) - 1],
			FileName:     itemFile.Name(),
			IsDir:        false,
			FileList:     make([]DirFile, 0),
		}
		tmpFileInfo, _ := f.GetFileInfo(tmpDirFile.FullFilePath)
		if tmpFileInfo.IsDir() {
			tmpDirFile.IsDir = true
			tmpDirFile.FileType = "dir"
			tmpDirFile.FileList, _ = f.GetDirFileList(dirPath + tmpDirFile.FileName + "/", depth - 1)
		}
		dirFileList = append(dirFileList, tmpDirFile)
	}

	return dirFileList, nil
}

// ReadFile 读取文件的内容
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:15:51
func (f *fileUtil) ReadFile(filePath string) ([]byte, error) {
	var (
		err error
		fileHandle *os.File
	)
	if _, err = f.GetFileInfo(filePath); nil != err {
		return nil, err
	}

	if fileHandle, err = os.Open(filePath); nil != err {
		return nil, errors.New("文件打开失败 ==> "+err.Error())
	}

	defer fileHandle.Close()
	return ioutil.ReadAll(fileHandle)
}
