# util说明

## string.go ： 定义一系列字符串操作的的用方法
> GenRandomString 生成指定长度的随机字符串
> Capitalize 字符串的首字母大写
> ToHump 将按照指定字符分割的字符串转换为小驼峰

## time.go : 定义了一系列关于时间操作的常用方法
> GetCurrentFormatTime 获取当前的格式化时间
> GetFormatTime 获取指定时间戳的格式化时间
> GetUnixTime 根据格式化时间，将其转换为unix时间戳

## file.go : 定义了一系列文件操作方法
> GetFileInfo 获取指定文件的信息
> IsDir 判断给定文件是否为目录
> GetDirFileList 递归获取一个目录下所有文件

## type.go : 定义文件类型转换方法,摘自 database/sql 包,将内部不可访问方法，可以访问
> ConvertAssign 数据类型转换