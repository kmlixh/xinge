# 腾讯信鸽Golang SDK（非官方版本）

[![Go Report Card](https://goreportcard.com/badge/gitee.com/kmlixh/xinge)](https://goreportcard.com/report/gitee.com/kmlixh/xinge)
[![GoDoc](https://godoc.org/gitee.com/kmlixh/xinge?status.svg)](https://godoc.org/gitee.com/kmlixh/xinge)

### 前言
##### 部分代码来源于信鸽官方推荐的[FrontMage](https://github.com/FrontMage/xinge)

    原本想使用FrontMage的库，但是看过源码之后觉得有些地方还是不太好，索性就自己动手改了一下。
    越改越多，发现和大神们的代码基本上不同了，再索性就发布出来，提供给大家！

## 用法

### 安装
`$ go get gitee.com/kmlixh/xinge`

### 全局推送
```go
package xinge

import (
	"testing"
	"time"
	"fmt"
	"gitee.com/kmlixh/xinge"
)

var authorAndroid = xinge.Authorization{AppID:"085f557303c8b", SecretKey:"046cf0c53a1bf6683bb22020a0ed8fec"}
var authoriOS = xinge.Authorization{AppID:"d5089ed7c3200", SecretKey:"d46a1b7d9d5327df90519d758cee8a1d"}
var xgClient =xinge.NewXingeClient(authorAndroid,authoriOS,xinge.NewHttpClient())

//测试推送一条面向全员的Notify消息
func main() {
	//以下两个配置信息都是真实的，但是并未配置正确的客户端，此处仅用来测试服务器的返回是否一致
	msg := xinge.NewPushAllNotifyPushMsg(xinge.PlatformAndroid, "测试的标题", "测试的内容"+time.Now().String())
	resp := xgClient.Push(msg)
	fmt.Println(resp)
}
```
### 用法解析：

#### XgClient

XgClient作为推送操作的执行者，内部封装了信鸽推送的相关方法和逻辑。
