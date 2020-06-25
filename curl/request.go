// Package curl ...
//
// File : request.go
//
// Decs :  Request构造类，用于设置请求参数，发起http请求
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:46:43
package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

// Request 请求类构造
//
// Author : go_developer@163.com<张德满>
type Request struct {
	cli             *http.Client
	req             *http.Request
	Raw             *http.Request
	Method          string
	URL             string
	dialTimeout     time.Duration
	responseTimeOut time.Duration
	Headers         map[string]string
	Cookies         map[string]string
	Queries         map[string]string
	PostData        map[string]interface{}
}

// NewRequest 创建一个Request实例
//
// Author : go_developer@163.com<张德满>
func NewRequest() *Request {
	r := &Request{}
	r.dialTimeout = 5
	r.responseTimeOut = 5
	return r
}

// SetMethod 设置请求方法
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:53:31
func (req *Request) SetMethod(method string) *Request {
	req.Method = method
	return req
}

// SetURL 设置请求的地址
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:54:02
func (req *Request) SetURL(url string) *Request {
	req.URL = url
	return req
}

// SetHeaders 设置请求头
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:54:41
func (req *Request) SetHeaders(headers map[string]string) *Request {
	req.Headers = headers
	return req
}

// setHeaders 将用户自定义请求头添加到http.Request实例上
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:55:10
func (req *Request) setHeaders() {
	for k, v := range req.Headers {
		req.req.Header.Set(k, v)
	}
	req.req.Header.Set("CURL-CLIENT", "CURL-CLIENT-GO-TOOL-CURL")
}

// SetCookies 设置请求cookies
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:55:54
func (req *Request) SetCookies(cookies map[string]string) *Request {
	req.Cookies = cookies
	return req
}

// setCookies 将用户自定义cookies添加到http.Request实例上
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:56:14
func (req *Request) setCookies() {
	for k, v := range req.Cookies {
		req.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

// SetQueries 设置url查询参数
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:56:42
func (req *Request) SetQueries(queries map[string]string) *Request {
	req.Queries = queries
	return req
}

// setQueries 将用户自定义url查询参数添加到http.Request上
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:57:00
func (req *Request) setQueries() {
	q := req.req.URL.Query()
	for k, v := range req.Queries {
		q.Add(k, v)
	}
	req.req.URL.RawQuery = q.Encode()
}

// SetPostData 设置post请求的提交数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:57:48
func (req *Request) SetPostData(postData map[string]interface{}) *Request {
	req.PostData = postData
	return req
}

// Get 发起get请求
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 00:58:17
func (req *Request) Get() (*Response, error) {
	return req.Send(req.URL, http.MethodGet)
}

// Delete 发起Delete请求
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:01:04
func (req *Request) Delete() (*Response, error) {
	return req.Send(req.URL, http.MethodDelete)
}

// Put 发起Delete请求
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:01:28
func (req *Request) Put() (*Response, error) {
	return req.Send(req.Url, http.MethodPut)
}

// Post 发起post请求
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:01:41
func (req *Request) Post() (*Response, error) {
	return req.Send(req.URL, http.MethodPost)
}

// PUT 发起put请求
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:01:58
func (req *Request) PUT() (*Response, error) {
	return req.Send(req.URL, http.MethodPut)
}

// PATCH 发起patch请求
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:02:21
func (req *Request) PATCH() (*Response, error) {
	return req.Send(req.URL, http.MethodPatch)
}

// SetDialTimeOut 设置连接超时时间
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:02:55
func (req *Request) SetDialTimeOut(TimeOutSecond int) {
	req.dialTimeout = time.Duration(TimeOutSecond)
}

// SetResponseTimeOut 设置读取超时时间
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:03:12
func (req *Request) SetResponseTimeOut(TimeOutSecond int) {
	req.responseTimeOut = time.Duration(TimeOutSecond)
}

// Send 发起请求
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:03:27
func (req *Request) Send(url string, method string) (*Response, error) {
	// 检测请求url是否填了
	if url == "" {
		return nil, errors.New("未设置请求接口")
	}
	// 检测请求方式是否填了
	if method == "" {
		return nil, errors.New("未指定请求方法")
	}
	// 初始化Response对象
	response := NewResponse()
	// 初始化http.Client对象
	req.cli = &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*req.dialTimeout)
				if err != nil {
					return nil, err
				}
				err = conn.SetDeadline(time.Now().Add(time.Second * req.dialTimeout))
				return conn, err
			},
			ResponseHeaderTimeout: time.Second * req.responseTimeOut,
		},
	}
	// 加载用户自定义的post数据到http.Request
	var payload io.Reader
	if strings.ToUpper(method) == http.MethodPost && nil != req.PostData {
		if jData, err := req.formatPostData(); err != nil {
			return nil, err
		} else {
			payload = bytes.NewReader([]byte(jData))
		}
	} else {
		payload = nil
	}

	var  err error
	if req.req, err = http.NewRequest(method, url, payload); err != nil {
		return nil, err
	}

	req.setHeaders()
	req.setCookies()
	req.setQueries()

	req.Raw = req.req

	if resp, err := req.cli.Do(req.req); err != nil {
		return nil, err
	} else {
		response.Raw = resp
	}

	defer response.Raw.Body.Close()

	response.parseHeaders()
	response.parseBody()

	return response, nil
}

// formatPostData 格式化postData
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:04:29
func (req *Request) formatPostData() (string, error) {
	paramStr := ""
	for key, val := range req.PostData {
		strVal, err := req.valToString(val)
		if nil != err {
			return "", nil
		}
		paramStr = paramStr + key + "=" + strVal + "&"
	}
	paramStr = strings.TrimRight(paramStr, "&")
	return paramStr, nil
}

// valToString 将任意一个是数据转化为字符串
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:04:51
func (req *Request) valToString(data interface{}) (string, error) {
	var val []byte
	var err error
	if val, err = json.Marshal(data); nil != err {
		return "", errors.New("json序列化失败 " + err.Error())
	}
	return strings.TrimRight(strings.TrimLeft(string(val), "\""), "\""), nil
}

