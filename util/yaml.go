// Package util ...
//
// File : yaml.go
//
// Decs : yml 工具集
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:49:37
package util

import (
	yaml "gopkg.in/yaml.v2"
)

// YamlUtil yaml 工具集
//
// Author : go_developer@163.com<张德满>
var YamlUtil *yamlUtil

func init()  {
	YamlUtil = &yamlUtil{}
}
type yamlUtil struct {

}

// ParseYamlFile 解析一个yml文件
//
// parseResult 接受文件的解析结果,需要是一个指针地址
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:38:54
func (y *yamlUtil) ParseYamlFile(yamlFilePath string, parseResult interface{}) error {
	var (
		err error
		ymlByte []byte
	)
	if ymlByte, err = FileUtil.ReadFile(yamlFilePath); nil != err {
		return err
	}
	return yaml.Unmarshal(ymlByte, parseResult)
}
