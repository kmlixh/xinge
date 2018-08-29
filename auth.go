package xinge

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// Authorization 用来添加请求Authorization
type Authorization struct {
	AppID     string
	SecretKey string
}

// Auth 添加一些默认的请求头
func (a *Authorization) Auth(req *http.Request) {
	// TODO: 全平台发送时如何填写Auth header
	req.Header.Add("Authorization", makeAuthHeader(a.AppID, a.SecretKey))
	req.Header.Add("Content-Type", "application/json")
}

// makeAuthHeader 根据appid和secretKey拼接并base64encode
func makeAuthHeader(appID, secretKey string) string {
	base64Str := base64.StdEncoding.EncodeToString(
		[]byte(
			fmt.Sprintf("%s:%s", appID, secretKey),
		),
	)
	return fmt.Sprintf("Basic %s", base64Str)
}
