package xinge

//DefaultPushMsg 创建基础的推送消息
func DefaultPushMsg(platform Platform, msgType MessageType, title string, content string) IPushMsg {
	msg := &PushMsg{
		Platform:    platform,
		PushID:      "0",
		MessageType: msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
	var optParams PushMsgOption
	if platform == PlatformAndroid {
		optParams = OptionAndroidParams(DefaultAndroidParams())
	} else {
		optParams = OptionIOSParams(DefaultIOSParams(title, content))
	}
	msg.RenderOptions(optParams)
	return msg
}

//NewTokenPushMsg 新建token类型的
func NewTokenPushMsg(platform Platform, msgType MessageType, title string, content string, tokens ...string) IPushMsg {
	msg := DefaultPushMsg(platform, msgType, title, content)
	if len(tokens) > 0 {
		msg.RenderOptions(OptionTokenList(tokens...))
	}
	return msg
}

//NewAccountPushMsg 基于account的推送
func NewAccountPushMsg(platform Platform, msgType MessageType, title string, content string, accounts ...string) IPushMsg {

	msg := DefaultPushMsg(platform, msgType, title, content)
	if len(accounts) == 0 {
		msg.RenderOptions(OptionAccountList(accounts...))
	}
	return msg
}

//NewTokenNotifyPushMsg 基于token的notify类型的推送
func NewTokenNotifyPushMsg(platform Platform, title string, content string, tokens ...string) IPushMsg {
	return NewTokenPushMsg(platform, MessageTypeOfNotify, title, content, tokens...)
}

//NewAccountNotifyPushMsg 基于account的notify类型的推送
func NewAccountNotifyPushMsg(platform Platform, title string, content string, accounts ...string) IPushMsg {
	return NewAccountPushMsg(platform, MessageTypeOfNotify, title, content, accounts...)
}

//NewTagPushMsg tag类型的推送
func NewTagPushMsg(platform Platform, msgType MessageType, title string, content string, tagOpt TagOperation, tags ...string) IPushMsg {
	if len(tags) == 0 {
		return nil
	}
	return &PushMsg{
		Platform:     platform,
		AudienceType: AudienceTypeTag,
		TagList:      &TagList{tags, tagOpt},
		PushID:       "0",
		MessageType:  msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}

//NewTagNotifyPushMsg 基于tag的notify类型的推送
func NewTagNotifyPushMsg(platform Platform, title string, content string, tagOpt TagOperation, tags ...string) IPushMsg {
	return NewTagPushMsg(platform, MessageTypeOfNotify, title, content, tagOpt, tags...)
}

//NewPushAllPushMsg 全员推送
func NewPushAllPushMsg(platform Platform, msgType MessageType, title string, content string) IPushMsg {
	return &PushMsg{
		Platform:     platform,
		AudienceType: AudienceTypeAll,
		PushID:       "0",
		MessageType:  msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}

//NewPushAllNotifyPushMsg 基于notify类型的全员推送
func NewPushAllNotifyPushMsg(platform Platform, title string, content string) IPushMsg {
	return NewPushAllPushMsg(platform, MessageTypeOfNotify, title, content)
}
