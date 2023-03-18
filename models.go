package openai_go

import "time"

// ModelList 获取模型列表
func (openAi *openAi) ModelList(timeout time.Duration) ([]*Model, error) {
	h := NewHttp("https://api.openai.com/v1/models", timeout)
	root, err := h.Get(openAi.getHeaders()) // 返回结果式gjson结构
	if err != nil {
		return nil, err
	}
	data := root.Get("data").Array() // 获取数据对象
	models := make([]*Model, 0)
	for _, model := range data {
		models = append(models, &Model{
			Id:         model.Get("id").String(),
			Object:     model.Get("model").String(),
			OwnedBy:    model.Get("owned_by").String(),
			Permission: model.Get("permission").Value(),
		})
	}
	return models, nil
}
