package xinge

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

/*
TODO List
1.自定义Client,包裹请求体,隐藏http.Client,处理结果返回:使返回结果更加可视化;
2.优化token_list和account_list两种类型的业务逻辑,目前并不支持超过一千个token或者一千个account的情况
3.改造Aps的数据结构,改造Alert的数据结构
4.优化和改造业务逻辑(简化和优化业务展现流程)

*/

//XgClient 用来推送消息，或者设置Tag的信鸽客户端
type XgClient struct {
	Android Authorization
	IOS     Authorization
	Client  *http.Client
}

//Push 推送消息
func (client XgClient) Push(rst IPushMsg) CommonRsp {
	var commonRsp CommonRsp
	temp := rst.nextRequest()
	pushId := "0"
	for ; temp != nil; temp = rst.nextRequest() {
		temp.RenderOptions(OptionPushID(pushId))
		var httpRequest *http.Request
		if rst.IsPlatform(PlatformAndroid) {
			httpRequest, _ = temp.toHttpRequest(client.Android)
		} else {
			httpRequest, _ = temp.toHttpRequest(client.IOS)
		}

		if httpRequest != nil {
			resp, _ := client.Client.Do(httpRequest)
			commonRsp = client.MarshalResp(resp)
			if len(commonRsp.PushID) > 0 {
				pushId = commonRsp.PushID
			}
		}
	}

	return commonRsp
}

//MarshalResp 解析返回
func (client XgClient) MarshalResp(resp *http.Response) CommonRsp {
	body, _ := ioutil.ReadAll(resp.Body)

	r := CommonRsp{}
	json.Unmarshal(body, &r)
	resp.Body.Close()
	return r
}

//NewHttpClient 创建一个默认的http.Client
func NewHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     30 * time.Second,
			DisableCompression:  false,
			// 默认开启了keep-alive
			DisableKeepAlives: false,
		},
	}
}
