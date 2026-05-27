package shortener_transport_http

import "github.com/gin-gonic/gin"

func (s *ShortenerHTTPHandler) HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialUrl, err := s.shortenerService.GetInitialUrl(c.Request.Context(), shortUrl)
	if err != nil {
		c.JSON(404, gin.H{"error": "short URL not found"})
		return
	}

	c.Redirect(302, initialUrl)
}
