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
type XingeClient struct {
	AndroidAuth Auther
	IOSAuth     Auther
	Client      *http.Client
}

func (client XingeClient) Push(rst *PushRequest) CommonRsp {
	var commonRsp CommonRsp
	temp := rst.nextRequest()
	pushId := "0"
	for ; temp != nil; temp = rst.nextRequest() {
		temp.RenderOptions(PushIDOpt(pushId))
		var httpRequest *http.Request
		if rst.Platform == PlatformAndroid {
			httpRequest, _ = temp.toHttpRequest(client.AndroidAuth)
		} else {
			httpRequest, _ = temp.toHttpRequest(client.IOSAuth)
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
func (client XingeClient) MarshalResp(resp *http.Response) CommonRsp {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	r := CommonRsp{}
	json.Unmarshal(body, &r)
	return r
}

// New 创建一个新的默认http客户端
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
