package openai_go

import (
	"testing"
	"time"
)

func TestOpenAi_ChatCompletions(t *testing.T) {
	openApi := NewOpenAI("")
	messages, err := openApi.ChatCompletions(&ChatCompletion{
		Model: "gpt-3.5-turbo",
		Messages: []*Message{
			{
				Role:    "user",
				Content: "who are you?",
			},
		},
	}, 3*time.Second)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log(ConvertJson(messages))

}

func TestOpenAi_CreateImages(t *testing.T) {
	openApi := NewOpenAI("")
	messages, err := openApi.CreateImages(&Image{
		Prompt: "美国短毛猫",
		N:      1,
		Size:   "1024x1024",
	}, 3*time.Second)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log(ConvertJson(messages))
}
