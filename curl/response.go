// Package curl ...
//
// File : response.go
//
// Decs : 响应信息文件
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:07:20
package curl

import (
	"io/ioutil"
	"net/http"
)

// Response 响应结果的数据结构
//
// Author : go_developer@163.com<张德满>
type Response struct {
	Raw     *http.Response
	Headers map[string]string
	Body    []byte
	Err     error
	Code    int
}

// NewResponse 获取相应信息的结构体
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:08:40
func NewResponse() *Response {
	return &Response{
		Headers: make(map[string]string),
	}
}

// IsOk 判断响应状态是否为200
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:09:00
func (resp *Response) IsOk() bool {
	return http.StatusOK == resp.Raw.StatusCode
}

// parseHeaders 解析响应header
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:09:21
func (resp *Response) parseHeaders() {
	resp.Code = resp.Raw.StatusCode
	for k, v := range resp.Raw.Header {
		resp.Headers[k] = v[0]
	}
}

// parseBody 解析响应body
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/06/26 01:09:38
func (resp *Response) parseBody() {
	resp.Body, resp.Err = ioutil.ReadAll(resp.Raw.Body)
}