package openai_go

import (
	"strings"
	"time"
)

// AddHumanSetting 添加人设
func (openAi *openAi) AddHumanSetting(content string) {
	openAi.Messages = append(openAi.Messages, &Message{
		Role:    "system",
		Content: content,
	})
}

// CleanHumanSetting 清除人设
func (openAi *openAi) CleanHumanSetting(content string) {
	openAi.Messages = append(openAi.Messages, &Message{
		Role:    "system",
		Content: content,
	})
}

// WeedOutMessage 淘汰最老的信息
func (openAi *openAi) WeedOutMessage() {
	if len(openAi.Messages) == 0 {
		openAi.Messages = []*Message{}
		return
	}
	openAi.Messages = append(openAi.Messages[:1], openAi.Messages[2:]...)
	return
}

// ChatCompletions 聊天交互
func (openAi *openAi) ChatCompletions(completion *ChatCompletion, timeout time.Duration) ([]*Message, error) {
	h := NewHttp("https://api.openai.com/v1/chat/completions", timeout)
	completion.Stream = false // 关闭流式渲染，流式渲染要用特殊的方式实现
	root, err := h.Post(ConvertJson(completion), openAi.getHeaders())
	if err != nil {
		message := "This model's maximum context length is 4096 token"
		if strings.Contains(root.Get("error").Get("message").String(), message) {
			openAi.WeedOutMessage() // 淘汰最老的信息
			return openAi.ChatCompletions(completion, timeout)
		}
		return nil, err
	}
	messages := make([]*Message, 0)
	choices := root.Get("choices").Array()
	if len(choices) > 0 {
		for _, choice := range choices {
			message := choice.Get("message")
			messages = append(messages, &Message{
				Role:    message.Get("role").String(),
				Content: message.Get("content").String(),
			})
		}
	}
	return messages, nil
}
