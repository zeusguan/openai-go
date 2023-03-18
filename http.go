package openai_go

import (
	"errors"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
	"time"
)

type Http struct {
	Url     string        `json:"url"`
	Timeout time.Duration `json:"timeout"`
}

// NewHttp 创建HTTP请求对象
func NewHttp(url string, timeout time.Duration) *Http {
	return &Http{
		Url:     url,
		Timeout: timeout,
	}
}

// Post 发起Post请求
func (h *Http) Post(data string, headers map[string]string) (*gjson.Result, error) {
	request, err := http.NewRequest("POST", h.Url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	// 绑定headers
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	client := http.Client{
		Timeout: h.Timeout,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	content := string(body)
	if !gjson.Valid(content) {
		return nil, errors.New("返回结果非JSON，body:" + content)
	}
	root := gjson.Parse(content)
	// 状态码
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(root.String())
	}
	return &root, nil
}

func (h *Http) Get(headers map[string]string) (*gjson.Result, error) {
	request, err := http.NewRequest("GET", h.Url, nil)
	if err != nil {
		return nil, err
	}
	// 绑定headers
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	client := http.Client{
		Timeout: h.Timeout,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	content := string(body)
	if !gjson.Valid(content) {
		return nil, errors.New("返回结果非JSON，body:" + content)
	}
	root := gjson.Parse(content)
	// 状态码
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(root.String())
	}
	return &root, nil
}
