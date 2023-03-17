package openai_go

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// BaseCompletion 基础完成对象
type BaseCompletion struct {
	Model            string                 `json:"model"`       // 模型
	Temperature      string                 `json:"temperature"` // 答案的偏移程度
	TopP             int                    `json:"top_p"`       // 答案的偏移程度 top和temperature不能一起用
	N                int                    `json:"n"`           // 输出条数
	Stream           bool                   `json:"stream"`      //  流式输出，可以优化体验
	Stop             string                 `json:"stop"`
	MaxTokens        int                    `json:"maxTokens"` // 支持的最大Token数，目前版本式4096，输入和输出加起来不能超过4096
	PresencePenalty  int                    `json:"presence_penalty"`
	FrequencyPenalty int                    `json:"frequency_penalty"`
	LogitBias        map[string]interface{} `json:"logit_bias"`
	User             string                 `json:"user"`
}

type ChatCompletion struct {
	BaseCompletion
	Messages []*Message `json:"messages"` // 问题
}

type Completion struct {
	BaseCompletion
	Prompt   string `json:"prompt"`
	Logprobs int    `json:"logprobs"`
	Suffix   string `json:"suffix"`
	Echo     bool   `json:"echo"`
	BestOf   int    `json:"best_of"`
}

// Image 图片生成对象
type Image struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_Format"`
}
