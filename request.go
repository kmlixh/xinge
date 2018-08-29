package xinge

//NewTokenRequest
func NewTokenRequest(platform Platform, msgType MsgType, title string, content string, tokens ...string) IPushMsg {
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

//NewAccountRequest
func NewAccountRequest(platform Platform, msgType MsgType, title string, content string, accounts ...string) IPushMsg {
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

//NewTokenNotifyRequest
func NewTokenNotifyRequest(platform Platform, title string, content string, tokens ...string) IPushMsg {
	return NewTokenRequest(platform, MsgTypeOfNotify, title, content, tokens...)
}

//NewAccountNotifyRequest
func NewAccountNotifyRequest(platform Platform, title string, content string, accounts ...string) IPushMsg {
	return NewAccountRequest(platform, MsgTypeOfNotify, title, content, accounts...)
}

//NewTagRequest
func NewTagRequest(platform Platform, msgType MsgType, title string, content string, tagOpt TagOperation, tags ...string) IPushMsg {
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

//NewTagNotifyRequest
func NewTagNotifyRequest(platform Platform, title string, content string, tagOpt TagOperation, tags ...string) IPushMsg {
	return NewTagRequest(platform, MsgTypeOfNotify, title, content, tagOpt, tags...)
}

//NewPushAllRequest
func NewPushAllRequest(platform Platform, msgType MsgType, title string, content string) IPushMsg {
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

//NewPushAllNotifyRequest
func NewPushAllNotifyRequest(platform Platform, title string, content string) IPushMsg {
	return NewPushAllRequest(platform, MsgTypeOfNotify, title, content)
}
