// Package util ...
//
// File : project.go
//
// Decs : 项目相关工具集
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:49:52
package util

import (
	"os"
)

// ProjectUtil 项目相关工具集
//
// Author : go_developer@163.com<张德满>
var ProjectUtil *projectUtil

type projectUtil struct {
}

// 获取项目目录
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 20:53:06
func (pu *projectUtil) GetCurrentPath() (string, error) {
	return os.Getwd()
}
