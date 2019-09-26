package util

import (
	"github.com/go-yaml/yaml"
)

var YamlUtil *yamlUtil

func init()  {
	YamlUtil = &yamlUtil{}
}
type yamlUtil struct {

}

/**
 * 解析一个yml文件
 * @param string yamlFilePath yml文件路径
 * @param interface parseResult 接受文件的解析结果,需要是一个指针地址
 * @return error 解析的异常信息
 * @author go_developer@163.com
 */
func (y *yamlUtil) ParseYamlFile(yamlFilePath string, parseResult interface{}) error {
	yaml.Parse()
	return nil
}
