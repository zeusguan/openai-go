package openai_go

import "time"

func (openAi *openAi) ChatCompletions(completion *ChatCompletion, timeout time.Duration) ([]*Message, error) {
	h := NewHttp("https://api.openai.com/v1/chat/completions", timeout)
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
