package xinge

//NewTokenRequest
func NewTokenRequest(platform Platform, msgType MessageType, title string, content string, tokens ...string) IPushMessage {
	var tps AudienceType
	lens := len(tokens)
	if lens == 0 {
		return nil
	} else if lens == 1 {
		tps = AdTypeToken
	} else if lens > 1 {
		tps = AdTypeTokenList
	}
	return &PushMsg{
		Platform:     platform,
		AudienceType: tps,
		TokenList:    tokens,
		PushID:       "0",
		MessageType:  msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}

//NewAccountRequest
func NewAccountRequest(platform Platform, msgType MessageType, title string, content string, accounts ...string) IPushMessage {
	var tps AudienceType
	lens := len(accounts)
	if lens == 0 {
		return nil
	} else if lens == 1 {
		tps = AdTypeAccount
	} else if lens > 1 {
		tps = AdTypeAccountList
	}
	return &PushMsg{
		Platform:     platform,
		AudienceType: tps,
		AccountList:  accounts,
		PushID:       "0",
		MessageType:  msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}

//NewTokenNotifyRequest
func NewTokenNotifyRequest(platform Platform, title string, content string, tokens ...string) IPushMessage {
	return NewTokenRequest(platform, MsgTypeNotify, title, content, tokens...)
}

//NewAccountNotifyRequest
func NewAccountNotifyRequest(platform Platform, title string, content string, accounts ...string) IPushMessage {
	return NewAccountRequest(platform, MsgTypeNotify, title, content, accounts...)
}

//NewTagRequest
func NewTagRequest(platform Platform, msgType MessageType, title string, content string, tagOpt TagOpration, tags ...string) IPushMessage {
	if len(tags) == 0 {
		return nil
	} else {
		return &PushMsg{
			Platform:     platform,
			AudienceType: AdTypeTag,
			TagList:      &TagList{tags, tagOpt},
			PushID:       "0",
			MessageType:  msgType,
			Message: Message{
				Title:   title,
				Content: content,
			}}
	}
}

//NewTagNotifyRequest
func NewTagNotifyRequest(platform Platform, title string, content string, tagOpt TagOpration, tags ...string) IPushMessage {
	return NewTagRequest(platform, MsgTypeNotify, title, content, tagOpt, tags...)
}

//NewPushAllRequest
func NewPushAllRequest(platform Platform, msgType MessageType, title string, content string) IPushMessage {
	return &PushMsg{
		Platform:     platform,
		AudienceType: AdTypeAll,
		PushID:       "0",
		MessageType:  msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}

//NewPushAllNotifyRequest
func NewPushAllNotifyRequest(platform Platform, title string, content string) IPushMessage {
	return NewPushAllRequest(platform, MsgTypeNotify, title, content)
}
