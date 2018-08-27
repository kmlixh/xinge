package xinge

func NewSingleTokenRequest(platform Platform, token string, title string, content string) IPushRequest {
	return &PushRequest{
		Platform:     platform,
		AudienceType: AdTypeToken,
		TokenList:    []string{token},
		PushID:       "0",
		MessageType:  MsgTypeMessage,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}
