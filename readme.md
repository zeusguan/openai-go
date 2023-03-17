# OpenAI Api Go

### ToDo List

1.支持设定人设。

2.基于GPT实现代码补全插件

### Usage

克隆代码

```shell
go get -u github.com/zeusguan/openai-go
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

