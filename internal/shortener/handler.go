package shortener

import "github.com/gin-gonic/gin"

func GenerateUrl(c *gin.Context) {
	c.JSON(200, gin.H{
		"url":     c.PostForm("url"),
		"message": "hello world man",
	})
}
