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
	"encoding/binary"
	"errors"
	"net"
	"os"
	"time"

	"github.com/go-developer/snowflake"
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

// IP2Long ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 21:35:43
func (pu *projectUtil) IP2Long(ipStr string) uint64 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint64(ip.To4())
}

// Long2IP ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 21:36:43
func (pu *projectUtil) Long2IP(ipLong uint64) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint64(ipByte, ipLong)
	return net.IP(ipByte).String()
}

// GenerateID 生成ID
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 22:00:50
func (pu *projectUtil) GenerateID(ipLong uint64) uint64 {
	var (
		ip  string
		err error
		id  int64
		iw  *snowflake.IDWorker
	)
	if ip, err = pu.GetServerIP(); nil != err {
		return 0
	}
	if iw, err = snowflake.NewIDWorker(int64(pu.IP2Long(ip))); nil != err {
		return 0
	}
	if id, err = iw.NextID(); nil != err {
		return 0
	}
	return uint64(id)
}
