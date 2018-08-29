# 腾讯信鸽Golang SDK（非官方版本）

[![Go Report Card](https://goreportcard.com/badge/gitee.com/kmlixh/xinge)](https://goreportcard.com/report/gitee.com/kmlixh/xinge)
[![GoDoc](https://godoc.org/gitee.com/kmlixh/xinge?status.svg)](https://godoc.org/gitee.com/kmlixh/xinge)

### 前言

`原本想使用信鸽官方推荐的 [FrontMage大神](https://github.com/FrontMage/xinge) 编写的库，但是发现有些地方还是不太好用，然后就自己动手改了一下，越改越多，慢慢发现和大神们的代码基本上不同了，索性就单独发布出来，提供给大家！`

## 用法

### 安装
`$ go get gitee.com/kmlixh/xinge`

### 全局推送
```go
package xinge

import (
	"testing"
	"time"
)

var authorAndroid = Authorization{AppID:"085f557303c8b", SecretKey:"046cf0c53a1bf6683bb22020a0ed8fec"}
var authoriOS = Authorization{AppID:"d5089ed7c3200", SecretKey:"d46a1b7d9d5327df90519d758cee8a1d"}
var xgClient = XgClient{Android:authorAndroid, IOS:authoriOS, Client:NewHttpClient()}

//测试推送一条面向全员的Notify消息
func TestXgClient_Push(t *testing.T) {
	//以下两个配置信息都是真实的，但是并未配置正确的客户端，此处仅用来测试服务器的返回是否一致
	msg := NewPushAllNotifyPushMsg(PlatformAndroid, "测试的标题", "测试的内容"+time.Now().String())
	resp := xgClient.Push(msg)
	if resp.RetCode != 0 {
		t.FailNow()
	}
}
```
上述代码复制自example_test.go

#### 基本用法解析：


### 苹果单账号push
```go
import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "fmt"
    "github.com/FrontMage/xinge/req"
    "github.com/FrontMage/xinge/auth"
)

func main() {
    auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
    pushReq, _ := req.NewSingleIOSAccountPush("account", "title", "content")
    auther.Auth(pushReq)

    c := &http.Client{}
    rsp, _ := c.Do(pushReq)
    defer rsp.Body.Close()
    body, _ := ioutil.ReadAll(rsp.Body)

    r := &xinge.CommonRsp{}
    json.Unmarshal(body, r)
    fmt.Printf("%+v", r)
}
```

### 安卓多账号push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
    req.Platform(xinge.PlatformAndroid),
    req.AudienceType(xinge.AdAccountList),
    req.MessageType(xinge.MsgTypeNotify),
    req.AccountList([]string{"10000031", "10000034"}),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
```

### iOS多账号push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
	req.Platform(xinge.PlatformiOS),
	req.EnvDev(),
    req.AudienceType(xinge.AdAccountList),
    req.MessageType(xinge.MsgTypeNotify),
    req.AccountList([]string{"10000031", "10000034"}),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
```

### 单设备push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
    req.Platform(xinge.PlatformiOS),
    req.EnvDev(),
    req.AudienceType(xinge.AdToken),
    req.MessageType(xinge.MsgTypeNotify),
    req.TokenList([]string{"10000031", "10000034"}),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
if r.RetCode != 0 {
    t.Errorf("Failed rsp=%+v", r)
}
```

### 多设备push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
    req.Platform(xinge.PlatformiOS),
    req.EnvDev(),
    req.AudienceType(xinge.AdTokenList),
    req.MessageType(xinge.MsgTypeNotify),
    req.TokenList([]string{"10000031", "10000034"}),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
if r.RetCode != 0 {
    t.Errorf("Failed rsp=%+v", r)
}
```

### 标签push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
    req.Platform(xinge.PlatformiOS),
    req.EnvDev(),
    req.AudienceType(xinge.AdTag),
    req.MessageType(xinge.MsgTypeNotify),
    req.TagList(&xinge.TagList{
        Tags:      []string{"new", "active"},
        Operation: xinge.TagListOpAnd,
    }),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
if r.RetCode != 0 {
    t.Errorf("Failed rsp=%+v", r)
}
```

## 贡献代码指南
目前的设计是通过`ReqOpt`函数来扩展各种请求参数，尽量请保持代码风格一致，使用`gofmt`来格式化代码。

贡献代码时可先从项目中的`TODO`开始，同时也欢迎提交新feature的PR和bug issue。
