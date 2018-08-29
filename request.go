package xinge

func NewTokenRequest(platform Platform, msgType MessageType, title string, content string, tokens ...string) IPushRequest {
	var tps AudienceType
	lens := len(tokens)
	if lens == 0 {
		return nil
	} else if lens == 1 {
		tps = AdTypeToken
	} else if lens > 1 {
		tps = AdTypeTokenList
	}
	return &PushRequest{
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
func NewAccountRequest(platform Platform, msgType MessageType, title string, content string, accounts ...string) IPushRequest {
	var tps AudienceType
	lens := len(accounts)
	if lens == 0 {
		return nil
	} else if lens == 1 {
		tps = AdTypeAccount
	} else if lens > 1 {
		tps = AdTypeAccountList
	}
	return &PushRequest{
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
func NewTokenNotifyRequest(platform Platform, title string, content string, tokens ...string) IPushRequest {
	return NewTokenRequest(platform, MsgTypeNotify, title, content, tokens...)
}
func NewAccountNotifyRequest(platform Platform, title string, content string, accounts ...string) IPushRequest {
	return NewAccountRequest(platform, MsgTypeNotify, title, content, accounts...)
}
func NewTagRequest(platform Platform, msgType MessageType, title string, content string, tagOpt TagOpration, tags ...string) IPushRequest {
	if len(tags) == 0 {
		return nil
	} else {
		return &PushRequest{
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
func NewTagNotifyRequest(platform Platform, title string, content string, tagOpt TagOpration, tags ...string) IPushRequest {
	return NewTagRequest(platform, MsgTypeNotify, title, content, tagOpt, tags...)
}
func NewPushAllRequest(platform Platform, msgType MessageType, title string, content string) IPushRequest {
	return &PushRequest{
		Platform:     platform,
		AudienceType: AdTypeAll,
		PushID:       "0",
		MessageType:  msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}
func NewPushAllNotifyRequest(platform Platform, title string, content string) IPushRequest {
	return NewPushAllRequest(platform, MsgTypeNotify, title, content)
}
