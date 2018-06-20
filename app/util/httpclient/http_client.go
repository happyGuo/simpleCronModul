package httpclient

// http-client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ResponseWrapper struct {
	StatusCode int
	Body       string
	Header     http.Header
}

func Get(url string, timeout int) (ResponseWrapper,error ){
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseWrapper{}, err
	}

	return request(req, timeout)
}

func PostParams(url string, params string, timeout int) (ResponseWrapper, error) {
	buf := bytes.NewBufferString(params)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return ResponseWrapper{}, err
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	return request(req, timeout)
}


func request(req *http.Request, timeout int) (ResponseWrapper,error){
	wrapper := ResponseWrapper{StatusCode: 0, Body: "", Header: make(http.Header)}
	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = time.Duration(timeout) * time.Second
	}
	setRequestHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		wrapper.Body = fmt.Sprintf("执行HTTP请求错误-%s", err.Error())
		return wrapper,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wrapper.Body = fmt.Sprintf("读取HTTP请求返回值失败-%s", err.Error())
		return wrapper,err
	}
	wrapper.StatusCode = resp.StatusCode
	wrapper.Body = string(body)
	wrapper.Header = resp.Header

	return wrapper,nil
}

//setRequestHeader 设置请求头
func setRequestHeader(req *http.Request) {
	req.Header.Set("User-Agent", "golang/zycron")
	req.Header.Set("Connection", "keep-alive")
}

func createRequestError(err error) ResponseWrapper {
	errorMessage := fmt.Sprintf("创建HTTP请求错误-%s", err.Error())
	return ResponseWrapper{0, errorMessage, make(http.Header)}
}
