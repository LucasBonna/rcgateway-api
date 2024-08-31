package middlewares

import (
	"log"
	"net/http/httputil"
	"net/url"
	"rc/gateway/initializers"
	"strings"

	"github.com/gin-gonic/gin"
)

func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		prefix, target := getPrefix(path)

		if target != nil {
			log.Printf("Redirecting %s to %s\n", path, target.String())

			// Ajuste do caminho apenas uma vez
			newURL := strings.TrimPrefix(path, prefix)
			c.Request.URL.Path = newURL

			// Criação do proxy reverso
			proxy := httputil.NewSingleHostReverseProxy(target)
			proxy.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func getPrefix(path string) (string, *url.URL) {
	if strings.HasPrefix(path, "/ehcrawler/") {
		target, _ := url.Parse(initializers.EHCRAWLER)
		return "/ehcrawler", target
	} else {
		return path, nil
	}
}
