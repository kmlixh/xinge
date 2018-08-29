package xinge

import (
	"testing"
	"time"
)

var authorAndroid = Authorization{AppID: "085f557303c8b", SecretKey: "046cf0c53a1bf6683bb22020a0ed8fec"}
var authoriOS = Authorization{AppID: "d5089ed7c3200", SecretKey: "d46a1b7d9d5327df90519d758cee8a1d"}
var xgClient = XgClient{Android: authorAndroid, IOS: authoriOS, Client: NewHttpClient()}

//测试推送一条面向全员的Notify消息
func TestXgClient_Push(t *testing.T) {
	//以下两个配置信息都是真实的，但是并未配置正确的客户端，此处仅用来测试服务器的返回是否一致
	msg := NewPushAllNotifyPushMsg(PlatformAndroid, "测试的标题", "测试的内容"+time.Now().String())
	resp := xgClient.Push(msg)
	if resp.RetCode != 0 {
		t.FailNow()
	}
}
