package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

type HttpClient struct {
	once   *sync.Once
	client *http.Client
}

var DefaultClient = &HttpClient{
	new(sync.Once),
	new(http.Client),
}

func init() {
	DefaultClient.create()
}

// DoReq 发送HTTP请求
func (c *HttpClient) DoReq(method, url string, reqBody io.Reader, header map[string]string) (rspBody []byte, err error) {
	fmt.Printf("method:%s, url:%s, reqBody:%v, header:%v\n", method, url, reqBody, header)
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=utf8")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rspBody, err = ioutil.ReadAll(resp.Body)
	fmt.Printf("rspbody:%s\n", rspBody)
	return
}

func (c *HttpClient) create() *http.Client {
	c.once.Do(func() {
		c.client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   5 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   5 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   100,
			},
			Timeout: 10 * time.Second,
		}
	})
	return c.client
}
