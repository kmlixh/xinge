package xinge

//NewTokenRequest 新建token类型的
func NewTokenPushMsg(platform Platform, msgType MsgType, title string, content string, tokens ...string) IPushMsg {
	var tps AudienceType
	lens := len(tokens)
	if lens == 0 {
		return nil
	} else if lens == 1 {
		tps = AudiTypeToken
	} else if lens > 1 {
		tps = AudiTypeTokenList
	}
	return &PushMsg{
		Platform:     platform,
		AudienceType: tps,
		TokenList:    tokens,
		PushID:       "0",
		MsgType:      msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}

//NewAccountPushMsg 基于account的推送
func NewAccountPushMsg(platform Platform, msgType MsgType, title string, content string, accounts ...string) IPushMsg {
	var tps AudienceType
	lens := len(accounts)
	if lens == 0 {
		return nil
	} else if lens == 1 {
		tps = AudiTypeAccount
	} else if lens > 1 {
		tps = AudiTypeAccountList
	}
	return &PushMsg{
		Platform:     platform,
		AudienceType: tps,
		AccountList:  accounts,
		PushID:       "0",
		MsgType:      msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}

//NewTokenNotifyPushMsg 基于token的notify类型的推送
func NewTokenNotifyPushMsg(platform Platform, title string, content string, tokens ...string) IPushMsg {
	return NewTokenPushMsg(platform, MsgTypeOfNotify, title, content, tokens...)
}

//NewAccountNotifyPushMsg 基于account的notify类型的推送
func NewAccountNotifyPushMsg(platform Platform, title string, content string, accounts ...string) IPushMsg {
	return NewAccountPushMsg(platform, MsgTypeOfNotify, title, content, accounts...)
}

//NewTagPushMsg tag类型的推送
func NewTagPushMsg(platform Platform, msgType MsgType, title string, content string, tagOpt TagOperation, tags ...string) IPushMsg {
	if len(tags) == 0 {
		return nil
	} else {
		return &PushMsg{
			Platform:     platform,
			AudienceType: AudiTypeTag,
			TagList:      &TagList{tags, tagOpt},
			PushID:       "0",
			MsgType:      msgType,
			Message: Message{
				Title:   title,
				Content: content,
			}}
	}
}

//NewTagNotifyPushMsg 基于tag的notify类型的推送
func NewTagNotifyPushMsg(platform Platform, title string, content string, tagOpt TagOperation, tags ...string) IPushMsg {
	return NewTagPushMsg(platform, MsgTypeOfNotify, title, content, tagOpt, tags...)
}

//NewPushAllPushMsg 全员推送
func NewPushAllPushMsg(platform Platform, msgType MsgType, title string, content string) IPushMsg {
	return &PushMsg{
		Platform:     platform,
		AudienceType: AudiTypeAll,
		PushID:       "0",
		MsgType:      msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}

//NewPushAllNotifyPushMsg 基于notify类型的全员推送
func NewPushAllNotifyPushMsg(platform Platform, title string, content string) IPushMsg {
	return NewPushAllPushMsg(platform, MsgTypeOfNotify, title, content)
}
