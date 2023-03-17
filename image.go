package openai_go

import "time"

func (openAi *openAi) CreateImages(completion *Image, timeout time.Duration) ([]string, error) {
	h := NewHttp("https://api.openai.com/v1/images/generations\n\n", timeout)
	root, err := h.Post(ConvertJson(completion), openAi.getHeaders())
	if err != nil {
		return nil, err
	}
	urls := make([]string, 0)
	dataList := root.Get("data").Array()
	if len(dataList) > 0 {
		for _, data := range dataList {
			urls = append(urls, data.Get("user").String())
		}
	}
	return urls, nil
}
