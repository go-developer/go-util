package curl

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Raw     *http.Response
	Headers map[string]string
	Body    []byte
	Err     error
	Code    int
}

/**
 * 获取相应信息的结构体
 * @author go_developer@163.com<张德满>
 */
func NewResponse() *Response {
	return &Response{
		Headers: make(map[string]string),
	}
}

/**
 * 判断响应状态是否为200
 * @author go_developer@163.com<张德满>
 */
func (resp *Response) IsOk() bool {
	return http.StatusOK == resp.Raw.StatusCode
}

/**
 * 解析响应header
 * @author go_developer@163.com<张德满>
 */
func (resp *Response) parseHeaders() {
	resp.Code = resp.Raw.StatusCode
	for k, v := range resp.Raw.Header {
		resp.Headers[k] = v[0]
	}
}

/**
 * 解析响应body
 * @author go_developer@163.com<张德满>
 */
func (resp *Response) parseBody() {
	resp.Body, resp.Err = ioutil.ReadAll(resp.Raw.Body)
}