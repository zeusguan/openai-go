package openai_go

type openAi struct {
	Token string `json:"token"`
}

// NewOpenAI 新建一个OpenAI对象
func NewOpenAI(token string) *openAi {
	return &openAi{
		Token: token,
	}
}

// getHeaders 获取请求头
func (openAi *openAi) getHeaders() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + openAi.Token,
	}
}
