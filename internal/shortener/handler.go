package shortener

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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

	existedShortenUrl, err := GetUrlByOriginUrl(&data)
	fmt.Println(existedShortenUrl, err)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if existedShortenUrl != nil {
		c.JSON(200, gin.H{
			"shorten_url":    os.Getenv("APP_URL") + "/" + existedShortenUrl.ShortenURL,
			"original_url":   existedShortenUrl.OriginalURL,
			"redirect_count": existedShortenUrl.RedirectCount,
		})
		return
	}

	shortenUrlData, err := SaveShortenUrl(data)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"shorten_url":    os.Getenv("APP_URL") + "/" + shortenUrlData.ShortenURL,
		"original_url":   shortenUrlData.OriginalURL,
		"redirect_count": shortenUrlData.RedirectCount,
	})
}

func RedirectToUrl(c *gin.Context) {
	urlToRedirect, err := GetUrlByShortenUrl(c.Param("url"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.Redirect(302, *urlToRedirect)
}

func GetStatisticByUrl(c *gin.Context) {
	fmt.Println("get statistic by url")
}
