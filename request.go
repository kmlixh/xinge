package xinge

func NewSingleTokenRequest(platform Platform, msgType MessageType, token string, title string, content string) IPushRequest {
	return &PushRequest{
		Platform:     platform,
		AudienceType: AdTypeToken,
		TokenList:    []string{token},
		PushID:       "0",
		MessageType:  msgType,
		Message: Message{
			Title:   title,
			Content: content,
		}}
}
