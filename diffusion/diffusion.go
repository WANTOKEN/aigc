package diffusion

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// GenerateImage 函数用于调用Stable Diffusion API生成图片
func GenerateImage(text string) (string, error) {
	// 构造请求体
	requestData := GenerateImageRequest{
		Text: text,
	}
	requestDataBytes, _ := json.Marshal(requestData)

	// 发送HTTP POST请求到Stable Diffusion API
	response, err := http.Post("http://example.com/diffusion-api/generate", "application/json", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 解析响应体中的JSON数据
	var responseBody GenerateImageResponse
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return "", err
	}

	// 返回生成的图片地址
	return responseBody.ImageURL, nil
}
