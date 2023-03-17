package openai_go

import "time"

func (openAi *openAi) Completions(completion *Completion, timeout time.Duration) ([]string, error) {
	h := NewHttp("https://api.openai.com/v1/completions", timeout)
	root, err := h.Post(ConvertJson(completion), openAi.getHeaders())
	if err != nil {
		return nil, err
	}
	messages := make([]string, 0)
	choices := root.Get("choices").Array()
	if len(choices) > 0 {
		for _, choice := range choices {
			message := choice.Get("text").String()
			messages = append(messages, message)
		}
	}
	return messages, nil
}
