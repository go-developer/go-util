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
	"errors"
	"net"
	"os"
	"time"
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

// GetServerIP 获取部署服务的服务器IP
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/05 20:36:57
func (pu *projectUtil) GetServerIP() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String(), nil
					}
				}
			}
		}
	}

	return "", errors.New("获取服务器IP失败")
}

// GetTraceID 生成traceID, 规则 时间 + serverIP + 随机字符串
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 00:57:14
func (pu *projectUtil) GetTraceID() string {
	serverIP, _ := pu.GetServerIP()
	return time.Now().Format("20060102150405") + "-" + serverIP + "-" + StringUtil.MD5(StringUtil.GenRandomString("", 32))
}
