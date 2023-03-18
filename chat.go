package openai_go

import "time"

func mergeMessage(message []*Message, chat *ChatCompletion) {
	chat.Messages = append(chat.Messages, message...)
}

// ChatCompletions 聊天交互
func (openAi *openAi) ChatCompletions(completion *ChatCompletion, timeout time.Duration) ([]*Message, error) {
	h := NewHttp("https://api.openai.com/v1/chat/completions", timeout)
	completion.Stream = false // 关闭流式渲染，流式渲染要用特殊的方式实现
	root, err := h.Post(ConvertJson(completion), openAi.getHeaders())
	if err != nil {
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
