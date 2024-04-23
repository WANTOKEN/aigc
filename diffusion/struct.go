package diffusion

// 请求Stable Diffusion API生成图片的结构体
type GenerateImageRequest struct {
	Text string `json:"text"`
}

// Stable Diffusion API返回的生成图片的结构体
type GenerateImageResponse struct {
	ImageURL string `json:"image_url"`
}
