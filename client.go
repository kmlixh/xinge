package xinge

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
TODO List
1.自定义Client,包裹请求体,隐藏http.client,处理结果返回:使返回结果更加可视化;
2.优化token_list和account_list两种类型的业务逻辑,目前并不支持超过一千个token或者一千个account的情况
3.改造Aps的数据结构,改造Alert的数据结构
4.优化和改造业务逻辑(简化和优化业务展现流程)

*/

//IAuth 授权工具接口
type IAuth interface {
	authRequest(req *http.Request)
}

// Authorization 用来添加请求Authorization
type Authorization struct {
	AppID     string
	SecretKey string
}

// authRequest 添加一些默认的请求头
func (a Authorization) authRequest(req *http.Request) {
	req.Header.Add("Authorization", MakeAuthHeader(a.AppID, a.SecretKey))
	req.Header.Add("Content-Type", "application/json")
}

//MakeAuthHeader 生成信鸽推送鉴权串
func MakeAuthHeader(appID, secretKey string) string {
	base64Str := base64.StdEncoding.EncodeToString(
		[]byte(
			fmt.Sprintf("%s:%s", appID, secretKey),
		),
	)
	return fmt.Sprintf("Basic %s", base64Str)
}

//MakeAuthoraztion 构造一个鉴权
func MakeAuthoraztion(appId string, secretKey string) IAuth {
	return &Authorization{appId, secretKey}
}

//XgClient 用来推送消息，或者设置Tag的信鸽客户端
type XgClient struct {
	android IAuth
	iOS     IAuth
	client  *http.Client
}

//SetAndroidAuth 设置默认的Android平台推送鉴权
func (xg *XgClient) SetAndroidAuth(auth IAuth) {
	xg.android = auth
}

//SetIOSAuth 设置默认的iOS平台推送鉴权
func (xg *XgClient) SetIOSAuth(auth IAuth) {
	xg.iOS = auth
}
func (xg *XgClient) SetAuth(appId string, secretKey string, platform Platform) {
	if platform == PlatformAndroid {
		xg.android = Authorization{appId, secretKey}
	} else {
		xg.iOS = Authorization{appId, secretKey}

	}
}

//Push 推送消息
func (xg XgClient) Push(msg IPushMsg) CommonRsp {
	if msg.equalsPlatform(PlatformAndroid) {
		return xg.PushWithAuthorization(msg, xg.android)
	}
	return xg.PushWithAuthorization(msg, xg.iOS)
}

//PushWithAuthorization 使用自定义的Authorization信息进行推送
func (xg XgClient) PushWithAuthorization(msg IPushMsg, auth IAuth) CommonRsp {
	if auth == nil {
		panic(errors.New("authorization could not be nil"))
	}
	var commonRsp CommonRsp
	temp := msg.nextRequest()
	pushId := "0"
	for ; temp != nil; temp = msg.nextRequest() {
		temp.RenderOptions(OptionPushID(pushId))
		var httpRequest *http.Request
		httpRequest, _ = temp.toHttpRequest(auth)
		if httpRequest != nil {
			resp, _ := xg.client.Do(httpRequest)
			commonRsp = xg.MarshalResp(resp)
			if len(commonRsp.PushID) > 0 {
				pushId = commonRsp.PushID
			}
		}
	}
	return commonRsp
}

//NewXingeClient3 无参数构建一个XgClient对象
func NewXingeClient3() XgClient {
	return NewXingeClient(nil, nil, NewHttpClient())
}

//NewXingeClient 构建一个完整的XgClient对象
func NewXingeClient(androidAuth IAuth, iOSAuth IAuth, client *http.Client) XgClient {
	return XgClient{androidAuth, iOSAuth, client}
}

//NewXingeClientent2 构建一个单平台的XgClient对象
func NewXingeClientent2(appId string, secretKey string, platform Platform) XgClient {
	xg := NewXingeClient3()
	xg.SetAuth(appId, secretKey, platform)
	return xg
}

//MarshalResp 解析返回
func (client XgClient) MarshalResp(resp *http.Response) CommonRsp {
	body, _ := ioutil.ReadAll(resp.Body)

	r := CommonRsp{}
	json.Unmarshal(body, &r)
	resp.Body.Close()
	return r
}

//NewHttpClient 创建一个默认的http.client
func NewHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     30 * time.Second,
			DisableCompression:  false,
			// 默认开启了keep-alive
			DisableKeepAlives: true,
		},
	}
}
