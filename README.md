# 腾讯信鸽Golang SDK（非官方版本）

[![Go Report Card](https://goreportcard.com/badge/gitee.com/kmlixh/xinge)](https://goreportcard.com/report/gitee.com/kmlixh/xinge)
[![GoDoc](https://godoc.org/gitee.com/kmlixh/xinge?status.svg)](https://godoc.org/gitee.com/kmlixh/xinge)

### 前言
##### 部分代码来源于信鸽官方推荐的[FrontMage](https://github.com/FrontMage/xinge)
##### 原本想使用FrontMage的库，但不太符合我懒人的习惯，于是自己动手改造一下，分享给大家！

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
### 类型解析：
```cgo
IAuth   //推送鉴权的接口；
Authorization //实现了IAuth接口，提供鉴权的相关操作
XgClient  //信鸽推送的客户端，推送相关的操作都由XgClient完成
IPushMSg //消息抽象接口
PushMsg   //被推送的消息，抽象和封装了推送消息本身。
Platform //平台；字符串的变体。目前只有两个，Android和iOS。这里不禁有人要问，为什么没有全平台？嗯，信鸽就是没有全平台！！！
Audience //推送目标，或者说推送受众。这里有All,tag,token,token_list等等，详细看官方文档。
PushMsgOption //函数模型，用于快速扩展和修改PushMsg的某个属性
//大概介绍这么多
```

### 用法解析：

#### XgClient相关

XgClient作为推送操作的执行者，内部封装了信鸽推送的相关方法和逻辑。

创建XgClient
```go
func NewXingeClient(androidAuth IAuth, iOSAuth IAuth, client *http.Client) XgClient
func NewXingeClient3() XgClient
func NewXingeClientent2(appId string, secretKey string, platform Platform) XgClient
```

XgClient的用法：
```go
//这个方法用来解析返回的http.Response为通用的返回数据，不多解释
func (xg XgClient) MarshalResp(resp *http.Response) CommonRsp
//推送一条消息
// 这里需要注意的是，这个操作会比对msg中的Platform属性和XgClient自身是否持有相应平台的IAuth
// 如果XgClient没有相应的IAuth，则会抛出异常
func (xg XgClient) Push(msg IPushMsg) CommonRsp
//使用指定的IAuth信息来推送一条消息
func (xg XgClient) PushWithAuthorization(msg IPushMsg, auth IAuth) CommonRsp

//设置Android平台的IAuth信息
func (xg *XgClient) SetAndroidAuth(auth IAuth)
//设置iOS平台的IAuth信息
func (xg *XgClient) SetIOSAuth(auth IAuth)
//直接通过appId,secretKey，以及平台信息来设置某个平台（Android或者iOS）的IAuth
func (xg *XgClient) SetAuth(appId string, secretKey string, platform Platform)

```
#### 如何创建推送消息？

这里提供了以下一些方法来帮助用户快速创建推送消息：
```go
func DefaultPushMsg(platform Platform, msgType MessageType, title string, content string) IPushMsg
func NewAccountNotifyPushMsg(platform Platform, title string, content string, accounts ...string) IPushMsg
func NewAccountPushMsg(platform Platform, msgType MessageType, title string, content string, accounts ...string) IPushMsg
func NewPushAllNotifyPushMsg(platform Platform, title string, content string) IPushMsg
func NewPushAllPushMsg(platform Platform, msgType MessageType, title string, content string) IPushMsg
func NewTagNotifyPushMsg(platform Platform, title string, content string, tagOpt TagOperation, tags ...string) IPushMsg
func NewTagPushMsg(platform Platform, msgType MessageType, title string, content string, tagOpt TagOperation, tags ...string) IPushMsg
func NewTokenNotifyPushMsg(platform Platform, title string, content string, tokens ...string) IPushMsg
func NewTokenPushMsg(platform Platform, msgType MessageType, title string, content string, tokens ...string) IPushMsg
```
包括了所有推送类型的推送方法，需要重点说明的是，token和tokenlist，account和account_list的类型区分，都是依靠末尾的不定长度的参数个数决定的！

例如：只传递一个token进入NewTokenPushMsg函数，则类型为token，如果为多个，则为token_list。account同理！

**重点说明：对于token_list和account_list类型的推送，如果推送数量超过1000，理论上会自动轮询推送，不需要再手动处理！请知悉！**

#### 如何修改推送的消息？
两种方法：

    1.手动修改；
    2.借助PushMsgOption进行修改；
第一种略过，我们说第二种，PushMsg具有一个非常好用的方法：

```go
func (rst *PushMsg) RenderOptions(opts ...PushMsgOption) error
```

这个方法能使用PushMsgOption对象来修改自身的属性。详细请看源码实现！！！

#### 如何创建PushMsgOption？

结合[FrontMage](https://github.com/FrontMage/xinge)大神写一些东西，做了修改和补充后，具有以下规模：

```go
func OptionAccountList(al ...string) PushMsgOption
func OptionAccountListAdd(a string) PushMsgOption
func OptionAccountType(at int) PushMsgOption
func OptionAddAction(k string, v interface{}) PushMsgOption
func OptionAndroidParams(params *AndroidParams) PushMsgOption
func OptionAps(aps *Aps) PushMsgOption
func OptionApsAlert(alert Alert) PushMsgOption
func OptionApsBadage(badge int) PushMsgOption
func OptionApsCategory(category string) PushMsgOption
func OptionApsContentAvailable(contentAvailable int) PushMsgOption
func OptionApsSound(sound string) PushMsgOption
func OptionApsThreadId(threadId string) PushMsgOption
func OptionBuilderID(id int) PushMsgOption
func OptionCleanable(c int) PushMsgOption
func OptionContent(c string) PushMsgOption
func OptionCustomContent(ct map[string]string) PushMsgOption
func OptionCustomContentSet(k, v string) PushMsgOption
func OptionEnvDevelop() PushMsgOption
func OptionEnvProduction() PushMsgOption
func OptionExpireTime(et time.Time) PushMsgOption
func OptionIOSParams(params *IOSParams) PushMsgOption
func OptionIconRes(ir string) PushMsgOption
func OptionIconType(it int) PushMsgOption
func OptionLights(l int) PushMsgOption
func OptionLoopTimes(lt int) PushMsgOption
func OptionMessage(m Message) PushMsgOption
func OptionMessageType(t MessageType) PushMsgOption
func OptionMultiPkg(mp bool) PushMsgOption
func OptionNID(id int) PushMsgOption
func OptionPlatAndroid() PushMsgOption
func OptionPlatIos() PushMsgOption
func OptionPlatform(p Platform) PushMsgOption
func OptionPushID(pid string) PushMsgOption
func OptionRing(ring int) PushMsgOption
func OptionRingRaw(rr string) PushMsgOption
func OptionSendTime(st time.Time) PushMsgOption
func OptionSeq(s int64) PushMsgOption
func OptionSmallIcon(si int) PushMsgOption
func OptionStatTag(st string) PushMsgOption
func OptionStyleID(s int) PushMsgOption
func OptionTagList(op TagOperation, tags ...string) PushMsgOption
func OptionTagListOpt2(tl TagList) PushMsgOption
func OptionTitle(t string) PushMsgOption
func OptionTokenList(tl ...string) PushMsgOption
func OptionTokenListAdd(t string) PushMsgOption
func OptionVibrate(v int) PushMsgOption
```
基本上已经包含了所有可以修改的部分，具体每一个函数用来修改那一块，可以看源码，源码中都有注释！！！

**特别说明：**

修改和扩充了Alert对象，给Alert增加了一些操作方法。详细的请涉及到苹果推送的同学自己尝试(没有苹果开发者证书，这块没法做测试，sorry)
```go
type Alert map[string]interface{}
func (alert Alert) Set(key string, value interface{})
func (alert Alert) SetActionLocKey(data string)
func (alert Alert) SetBody(content string)
func (alert Alert) SetLaunchImage(data string)
func (alert Alert) SetLocArgs(data []string)
func (alert Alert) SetLocKey(data string)
func (alert Alert) SetTitle(title string)
func (alert Alert) SetTitleLocArgs(data []string)
func (alert Alert) SetTitleLocKey(data string)
```
几个或许有用处的方法：

```go
func DefaultApsAlert(title string, body string) Alert

func DefaultIOSParams(title string, content string) *IOSParams

func DefaultAndroidParams() *AndroidParams

func DefaultPushMsg(platform Platform, msgType MessageType, title string, content string) IPushMsg

```
相应的，如果想看特别详细的代码文档，点击文档顶部的godoc图标，或者是下面的图标

[![Go Report Card](https://goreportcard.com/badge/gitee.com/kmlixh/xinge)](https://goreportcard.com/report/gitee.com/kmlixh/xinge)
[![GoDoc](https://godoc.org/gitee.com/kmlixh/xinge?status.svg)](https://godoc.org/gitee.com/kmlixh/xinge)

### 捐助作者

<img src="https://gitee.com/kmlixh/xinge/raw/master/img/wechat.png" width="200"  align=center />
&nbsp;&nbsp;&nbsp;
<img src="https://gitee.com/kmlixh/xinge/raw/master/img/alipay.png" width="200" align=center />

