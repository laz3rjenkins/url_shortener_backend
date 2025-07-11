package shortener

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"minq-backend/storage"
	"net/http"
	"os"
)

func GenerateUrl(c *gin.Context) {
	var shorten Shorten

	repository := NewRepository(storage.DB)

	if err := c.ShouldBind(&shorten); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := ShortenUrlAttributes{
		OriginalURL:   shorten.Url,
		ShortenURL:    GenerateShortenString(),
		RedirectCount: 0,
	}

	existedShortenUrl, err := repository.GetByOriginalURL(c, data.OriginalURL)
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

	err = repository.SaveShortenURL(c, data)
	if err := repository.SaveShortenURL(c, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"shorten_url":    os.Getenv("APP_URL") + "/" + data.ShortenURL,
		"original_url":   data.OriginalURL,
		"redirect_count": data.RedirectCount,
	})
}

func RedirectToUrl(c *gin.Context) {
	repository := NewRepository(storage.DB)

	urlToRedirect, err := repository.GetByShortenURL(c, c.Param("url"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	if urlToRedirect == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shorten url not found"})
		return
	}

	c.Redirect(302, *urlToRedirect)
}

func GetStatisticByUrl(c *gin.Context) {
	fmt.Println("get statistic by url")
}
