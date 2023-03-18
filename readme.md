# OpenAI Api Go

### Abstract

Go版本的OpenAI的接口封装，可用于快速构建Chat服务，能够短时间能构建自己的机器人。

### Use

克隆代码

```shell
go get -u github.com/zeusguan/openai-go
```

参数详解

```go
type ChatCompletion struct {
	Model            string                 `json:"model"`       // 模型
	Temperature      string                 `json:"temperature"` // 答案的偏移程度
	TopP             int                    `json:"top_p"`       // 答案的偏移程度 top和temperature不能一起用
	N                int                    `json:"n"`           // 输出条数
	Stream           bool                   `json:"stream"`      //  流式输出，可以优化体验
	Stop             string                 `json:"stop"`        //  停止标识，可以diy，如果不填为空则使用默认
	MaxTokens        int                    `json:"maxTokens"`   // 支持的最大Token数，目前版本式4096，输入和输出加起来不能超过4096
	PresencePenalty  int                    `json:"presence_penalty"`  // 增加回答的随性
	FrequencyPenalty int                    `json:"frequency_penalty"` // 增加回答的随性
	LogitBias        map[string]interface{} `json:"logit_bias"`
	User             string                 `json:"user"`
	Messages         []*Message             `json:"messages"` // 问题
}

```

编写调用

```go
import (
	goapi "github.com/zeusguan/openai-go"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getText() {
	openai := goapi.NewOpenAI("this is token")
	openai.CreateImages()    // 图片生成
	openai.ChatCompletions() // 对话模式
	openai.Completions()     // 基于GPT3的模型接口
}
```

