package main

import (
	"net/http"

	DF "aigc/diffusion"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Gin引擎
	r := gin.Default()

	// 定义API端点来接收请求并调用Stable Diffusion生成图片
	r.POST("/generate-image", func(c *gin.Context) {
		// 在这里调用Stable Diffusion API生成图片，并返回结果
		// 解析请求体中的JSON数据
		var requestBody DF.GenerateImageRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 调用生成图片的函数
		imageURL, err := DF.GenerateImage(requestBody.Text)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 返回生成的图片地址给客户端
		c.JSON(http.StatusOK, gin.H{"image_url": imageURL})

	})

	// 启动HTTP服务器
	r.Run(":8080")
}
