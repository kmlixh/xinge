package xinge

import (
	"net/http"
)

// XingeURL 信鸽API地址
var XingeURL = "https://openapi.xg.qq.com/v3/push/app"

// PushURL 修改信鸽的请求URL
func PushURL(url string) {
	XingeURL = url
}

// NewSingleIOSAccountPush 新建一个iOS单账号push请求
func NewSingleIOSAccountPush(account, title, content string, opts ...ReqOption) (*http.Request, error) {
	req := &Request{
		Platform:     PlatformiOS,
		MessageType:  MsgTypeNotify,
		AudienceType: AdTypeAccount,
		AccountList:  []string{account},
		Message: Message{
			Title:   title,
			Content: content,
			IOS: &IOSParams{
				Aps: &Aps{
					Alert: map[string]string{
						"title":   title,
						"content": content,
					},
					Badge: 1,
					Sound: "default",
				},
			},
		},
	}
	return req.RenderOptions(opts...)
}

// NewSingleAndroidAccountPush 新建一个安卓通知栏push请求
func NewSingleAndroidAccountPush(account, title, content string, opts ...ReqOption) (*http.Request, error) {
	req := &Request{
		Platform:     PlatformAndroid,
		MessageType:  MsgTypeNotify,
		AudienceType: AdTypeAccount,
		AccountList:  []string{account},
		Message: Message{
			Title:   title,
			Content: content,
		},
	}
	return req.RenderOptions(opts...)
}
