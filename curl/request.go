/**
 * golang版本的curl请求库
 * Request构造类，用于设置请求参数，发起http请求
 * @author go_developer@163.com<张德满>
 */

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

/**
 * http 请求类构造
 * @author go_developer@163.com
 */
type Request struct {
	cli             *http.Client
	req             *http.Request
	Raw             *http.Request
	Method          string
	Url             string
	dialTimeout     time.Duration
	responseTimeOut time.Duration
	Headers         map[string]string
	Cookies         map[string]string
	Queries         map[string]string
	PostData        map[string]interface{}
}

// 创建一个Request实例
func NewRequest() *Request {
	r := &Request{}
	r.dialTimeout = 5
	r.responseTimeOut = 5
	return r
}

/**
 * 设置请求方法
 * @author go_developer@163.com<张德满>
 */
func (req *Request) SetMethod(method string) *Request {
	req.Method = method
	return req
}

/**
 * 设置请求的地址
 * @author go_developer@163.com<张德满>
 */
func (req *Request) SetUrl(url string) *Request {
	req.Url = url
	return req
}

/**
 * 设置请求头
 * @author go_developer@163.com<张德满>
 */
func (req *Request) SetHeaders(headers map[string]string) *Request {
	req.Headers = headers
	return req
}

/**
 * 将用户自定义请求头添加到http.Request实例上
 * @author go_developer@163.com<张德满>
 */
func (req *Request) setHeaders() {
	for k, v := range req.Headers {
		req.req.Header.Set(k, v)
	}
	req.req.Header.Set("CURL-CLIENT", "CURL-CLIENT-GO-TOOL-CURL")
}

/**
 * 设置请求cookies
 * @author go_developer@163.com<张德满>
 */
func (req *Request) SetCookies(cookies map[string]string) *Request {
	req.Cookies = cookies
	return req
}

/**
 * 将用户自定义cookies添加到http.Request实例上
 * @author go_developer@163.com<张德满>
 */
func (req *Request) setCookies() {
	for k, v := range req.Cookies {
		req.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

/**
 * 设置url查询参数
 * @author go_developer@163.com<张德满>
 */
func (req *Request) SetQueries(queries map[string]string) *Request {
	req.Queries = queries
	return req
}

/**
 * 将用户自定义url查询参数添加到http.Request上
 * @author go_developer@163.com<张德满>
 */
func (req *Request) setQueries() {
	q := req.req.URL.Query()
	for k, v := range req.Queries {
		q.Add(k, v)
	}
	req.req.URL.RawQuery = q.Encode()
}

/**
 * 设置post请求的提交数据
 * @author go_developer@163.com<张德满>
 */
func (req *Request) SetPostData(postData map[string]interface{}) *Request {
	req.PostData = postData
	return req
}

// 发起get请求
func (req *Request) Get() (*Response, error) {
	return req.Send(req.Url, http.MethodGet)
}

// 发起Delete请求
func (req *Request) Delete() (*Response, error) {
	return req.Send(req.Url, http.MethodDelete)
}

// 发起Delete请求
func (req *Request) Put() (*Response, error) {
	return req.Send(req.Url, http.MethodPut)
}

// 发起post请求
func (req *Request) Post() (*Response, error) {
	return req.Send(req.Url, http.MethodPost)
}

// 发起put请求
func (req *Request) PUT() (*Response, error) {
	return req.Send(req.Url, http.MethodPut)
}

// 发起patch请求
func (req *Request) PATCH() (*Response, error) {
	return req.Send(req.Url, http.MethodPatch)
}

/*
 * 设置连接超时时间
 * @author go_developer@163.com<张德满>
 */
func (req *Request) SetDialTimeOut(TimeOutSecond int) {
	req.dialTimeout = time.Duration(TimeOutSecond)
}

/*
 * 设置读取超时时间
 * @author go_developer@163.com<张德满>
 */
func (req *Request) SetResponseTimeOut(TimeOutSecond int) {
	req.responseTimeOut = time.Duration(TimeOutSecond)
}

/*
 * 发起请求
 * @author go_developer@163.com<张德满>
 */
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

/**
 * 格式化postData
 * @author go_developer@163.com
 */
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

/**
 * 将任意一个是数据转化为字符串
 * @author go_developer@163.com<张德满>
 */
func (req *Request) valToString(data interface{}) (string, error) {
	var val []byte
	var err error
	if val, err = json.Marshal(data); nil != err {
		return "", errors.New("json序列化失败 " + err.Error())
	}
	return strings.TrimRight(strings.TrimLeft(string(val), "\""), "\""), nil
}

