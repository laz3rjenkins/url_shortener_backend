package shortener

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateUrl(c *gin.Context) {
	var shorten Shorten

	if err := c.ShouldBind(&shorten); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := ShortenUrlAttributes{
		OriginalURL:   shorten.Url,
		ShortenURL:    GenerateShortenString(),
		RedirectCount: 0,
	}

	if err := SaveShortenUrl(data); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"url":     c.PostForm("url"),
		"message": "hello world man",
	})
}

func RedirectToUrl(c *gin.Context) {
	//c.Redirect(302, c.Param("url"))
	fmt.Println("redirect to url")
}

func GetStatisticByUrl(c *gin.Context) {
	fmt.Println("get statistic by url")
}
