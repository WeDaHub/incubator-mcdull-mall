package utils

import (
	"math/rand"
	"net/http"
	"time"
)

func RandomNumberStr(lenth int) string {
	chars := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	tmpArr := make([]byte, 0, lenth)
	rand.Seed(time.Now().Unix())
	for i := 0; i < lenth; i++ {
		n := rand.Intn(len(chars))
		tmpArr = append(tmpArr, chars[n])
	}
	return string(tmpArr)
}

func PhoneMark(phone string) string {
	return phone[0:3] + "****" + phone[7:]
}

// ReadUserIP 获取客户端IP地址
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
