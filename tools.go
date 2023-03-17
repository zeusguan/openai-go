package openai_go

import "encoding/json"

func ConvertJson(obj interface{}) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(marshal)
}
