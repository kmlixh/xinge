package xinge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// TODO: 信鸽人工服务
// 后端API V3，iOS 发push的时候message_type填message的时候报错

// 基础鉴权为未通过，请检查key[appKey|groupKey]是否与push_type匹配！

// 另外由于iOS和Android是两个不同的appid和key，如果platform填all，即全平台推送的时候怎么填Authorization header？

// CommonRspEnv 信鸽推送接口通用基础返回值的environment字段
type CommonRspEnv string

const (
	// EnvProd 生产环境
	EnvProd CommonRspEnv = "product"
	// EnvDev 测试环境
	EnvDev CommonRspEnv = "dev"
)

// XingeURL 信鸽API地址
var XingeURL = "https://openapi.xg.qq.com/v3/push/app"

// PushURL 修改信鸽的请求URL
func PushURL(url string) {
	XingeURL = url
}

type ReqOption func(*PushRequest) error

// CommonRsp 信鸽推送接口的通用基础返回值
type CommonRsp struct {
	// TODO: doc this
	Seq int64 `json:"seq"`
	// 推送id
	PushID string `json:"push_id"`
	// 错误码
	RetCode int `json:"ret_code"`
	// 用户指定推送环境，仅支持iOS
	Environment CommonRspEnv `json:"environment"`
	// 结果描述
	ErrMsg string `json:"err_msg,omitempty"`
	// 请求正确且有额外数据时，则结果在这个字段中
	Result map[string]string `json:"result,omitempty"`
}

// RequestAudienceType 推送目标
type AudienceType string

const (
	// AdAll 全量推送
	AdTypeAll AudienceType = "all"
	// AdTag 标签推送
	AdTypeTag AudienceType = "tag"
	// AdToken  单设备推送
	AdTypeToken AudienceType = "token"
	// AdTokenList  设备列表推送
	AdTypeTokenList AudienceType = "token_list"
	// AdTypeAccount 单账号推送
	AdTypeAccount AudienceType = "account"
	// AdAccountList 账号列表推送
	AdTypeAccountList AudienceType = "account_list"
)

// Platform push API platform参数
type Platform string

const (
	PlatformAndroid Platform = "android"
	// PlatformiOS 苹果推送
	PlatformiOS Platform = "ios"
)

// MessageType push API message_type参数
type MessageType string

const (
	// MsgTypeNotify 消息类型为通知栏消息
	MsgTypeNotify MessageType = "notify"
	// MsgTypeMessage 消息类型为透传消息(Android)/静默消息(iOS)
	MsgTypeMessage MessageType = "message"
)

// PushRequest push API 必要参数
type PushRequest struct {
	// 受众类型，见AudienceType类型
	AudienceType `json:"audience_type"`
	// 推送平台，见Platform类型
	Platform `json:"platform"`
	// 消息内容
	Message `json:"message"`
	// 消息类型，见MessageType类型
	MessageType `json:"message_type"`

	// 当AudienceType == AdTag时必填
	TagList *TagList `json:"tag_list,omitempty"`
	// 当AudienceType == AdToken 或 AdTokenList 时必填的参数，
	// 当AdToken时即使传了多个token，也只有第一个会被推送
	// 当AdTokenList时，最多支持1000个token，同时push_id第一次请求时必须填0
	// 系统会返回一个push_id = 123(例)，后续推送如果push_id填写123(例)
	// 则会使用跟123相同的文案推送
	TokenList []string `json:"token_list,omitempty"`
	// 当AudienceType == AdTypeAccount 或 AdAccountList 时必填的参数，
	// 当AdAccount时即使传了多个token，也只有第一个会被推送
	// 当AdAccountList时，最多支持1000个token，同时push_id第一次请求时必须填0
	// 系统会返回一个push_id = 123(例)，后续推送如果push_id填写123(例)
	// 则会使用跟123相同的文案推送
	AccountList []string `json:"account_list,omitempty"`

	// 	消息离线存储时间（单位为秒）
	// 最长存储时间3天，若设置为0，则默认值（3天）
	// 建议取值区间[600, 86400x3]
	// 第三方通道离线保存消息不同厂商标准不同
	ExpireTime int `json:"expire_time,omitempty"`

	// 	指定推送时间
	// 格式为yyyy-MM-DD HH:MM:SS
	// 若小于服务器当前时间，则会立即推送
	// 仅全量推送和标签推送支持此字段
	SendTime string `json:"send_time,omitempty"`

	// 	多包名推送
	// 当app存在多个不同渠道包（例如应用宝、豌豆荚等），推送时如果是希望手机上安装任何一个渠道的app都能收到消息那么该值需要设置为true
	MultiPkg bool `json:"multi_pkg,omitempty"`

	// 	循环任务重复次数
	// 仅支持全推、标签推
	// 建议取值[1, 15]
	LoopTimes int `json:"loop_times,omitempty"`

	// 	用户指定推送环境，仅限iOS平台推送使用
	// product： 推送生产环境
	// dev： 推送开发环境
	Environment CommonRspEnv `json:"environment,omitempty"`

	// 	统计标签，用于聚合统计
	// 使用场景(示例)：
	// 现在有一个活动id：active_picture_123,需要给10000个设备通过单推接口（或者列表推送等推送形式）下发消息，同时设置该字段为active_picture_123
	// 推送完成之后可以使用v3统计查询接口，根据该标签active_picture_123 查询这10000个设备的实发、抵达、展示、点击数据
	StatTag string `json:"stat_tag,omitempty"`

	// 	接口调用时，在应答包中信鸽会回射该字段，可用于异步请求
	// 使用场景：异步服务中可以通过该字段找到server端返回的对应应答包
	Seq int64 `json:"seq,omitempty"`

	// 单账号推送时可选
	// 	 账号类型，参考后面账号说明。
	//  必须与账号绑定时设定的账号类型一致
	AccountType int `json:"account_type,omitempty"`

	// 账号列表推送、设备列表推送时必需
	// 账号列表推送和设备列表推送时，第一次推送该值填0，系统会创建对应的推送任务，
	// 并且返回对应的pushid：123，后续推送push_id 填123(同一个文案）
	// 表示使用与123 id 对应的文案进行推送。(注：文案的有效时间由前面的expire_time 字段决定）
	PushID string `json:"push_id,omitempty"`

	nextIndex int `json:"-"` //如果推送的account,token大于1000,需要轮询推送
}
type IPushRequest interface {
	RenderOptions(opts ...ReqOption) error
	clone(options ...ReqOption) IPushRequest
	nextRequest() IPushRequest
	toHttpRequest(auther Auther) (request *http.Request, err error)
	IsPlatform(platform Platform) bool
}

func (rst *PushRequest) RenderOptions(opts ...ReqOption) error {
	for _, opt := range opts {
		err := opt(rst)
		return err
	}
	return nil
}
func (rst *PushRequest) clone(options ...ReqOption) IPushRequest {
	request := PushRequest{
		AudienceType: rst.AudienceType,
		Platform:     rst.Platform,
		Message:      rst.Message,
		MessageType:  rst.MessageType,
		TagList:      rst.TagList,
		TokenList:    rst.TokenList,
		AccountList:  rst.AccountList,
		ExpireTime:   rst.ExpireTime,
		SendTime:     rst.SendTime,
		MultiPkg:     rst.MultiPkg,
		LoopTimes:    rst.LoopTimes,
		Environment:  rst.Environment,
		StatTag:      rst.StatTag,
		Seq:          rst.Seq,
		AccountType:  rst.AccountType,
		PushID:       rst.PushID}
	err := request.RenderOptions(options...)
	if err != nil {
		return nil
	}
	return &request
}
func (rst *PushRequest) nextRequest() IPushRequest {
	var request IPushRequest
	if rst.AudienceType == AdTypeAccountList {
		request = rst.clone(sliceAccountList)
	} else if rst.AudienceType == AdTypeTokenList {
		request = rst.clone(sliceTokenList)
	} else if rst.nextIndex == 0 {
		rst.nextIndex = -1
		return rst
	} else {
		request = nil
	}
	return request
}
func (rst *PushRequest) IsPlatform(platform Platform) bool {
	return rst.Platform == platform
}

func (rst PushRequest) toHttpRequest(auther Auther) (request *http.Request, err error) {

	bodyBytes, err := json.Marshal(rst)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bodyBytes))
	request, err = http.NewRequest("POST", XingeURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	auther.Auth(request)
	return
}
func sliceAccountList(rst *PushRequest) error {
	lens := len(rst.AccountList)
	if lens > rst.nextIndex { //大于1000需要二次提交,如果此处长度大于nextIndex,则说明还有(account或者token没有推送)
		end := rst.nextIndex + 1000
		if end > lens {
			end = lens - rst.nextIndex
		}
		rst.AccountList = rst.AccountList[rst.nextIndex:end]
		return nil
	} else {
		return errors.New("index out of range!")
	}
}
func sliceTokenList(rst *PushRequest) error {
	lens := len(rst.TokenList)
	if lens > rst.nextIndex { //大于1000需要二次提交,如果此处长度大于nextIndex,则说明还有(account或者token没有推送)
		end := rst.nextIndex + 1000
		if end > lens {
			end = lens - rst.nextIndex
		}
		rst.TokenList = rst.TokenList[rst.nextIndex:end]
		rst.nextIndex = end
		return nil
	} else {
		return errors.New("index out of range!")
	}
}

// TagOpration 标签推送参数的逻辑操作符
type TagOpration string

const (
	// OprationAnd 推送tag1且tag2
	OprationAnd TagOpration = "AND"
	// OprationOr 推送tag1或tag2
	OprationOr TagOpration = "OR"
)

// TagList 标签推送参数
type TagList struct {
	// 标签
	Tags []string `json:"tags"`
	// 标签逻辑操作符
	TagOpration `json:"op"`
}

// Message 消息体
type Message struct {
	Title      string   `json:"title,omitempty"`
	Content    string   `json:"content,omitempty"`
	AcceptTime []string `json:"accept_time,omitempty"`

	Android *AndroidParams `json:"android,omitempty"`

	IOS *IOSParams `json:"ios,omitempty"`
}

// AndroidParams 安卓push参数
type AndroidParams struct {
	// TODO: doc these
	NID           int                     `json:"n_id,omitempty"`
	BuilderID     int                     `json:"builder_id,omitempty"`
	Ring          int                     `json:"ring,omitempty"`
	RingRaw       string                  `json:"ring_raw,omitempty"`
	Vibrate       int                     `json:"vibrate,omitempty"`
	Lights        int                     `json:"lights,omitempty"`
	Clearable     int                     `json:"clearable,omitempty"`
	IconType      int                     `json:"icon_type,omitempty"`
	IconRes       string                  `json:"icon_res,omitempty"`
	StyleID       int                     `json:"style_id,omitempty"`
	SmallIcon     int                     `json:"small_icon,omitempty"`
	Action        *map[string]interface{} `json:"action,omitempty"`
	CustomContent *map[string]string      `json:"custom_content,omitempty"`
}

// IOSParams iOS push参数
type IOSParams struct {
	Aps    *Aps              `json:"aps,omitempty"`
	Custom map[string]string `json:"custom,omitempty"`
}

// Aps 通知栏iOS消息的aps字段，详情请参照苹果文档
type Aps struct {
	Alert            map[string]string `json:"alert,omitempty"`
	Badge            int               `json:"badge,omitempty"`
	Category         string            `json:"category,omitempty"`
	ContentAvailable int               `json:"content-available,omitempty"`
	Sound            string            `json:"sound,omitempty"`
}

// TODO: error type constant
