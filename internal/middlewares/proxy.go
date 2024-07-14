package middlewares

import (
	"log"
	"net/http/httputil"
	"net/url"
	"strings"
	"web/gin/initializers"

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
	if strings.HasPrefix(path, "/rcstorage/") {
		target, _ := url.Parse(initializers.RCSTORAGE)
		return "/rcstorage", target
	} else if strings.HasPrefix(path, "/rcauth/") {
		target, _ := url.Parse(initializers.RCAUTH)
		return "/rcauth", target
	} else if strings.HasPrefix(path, "/rctracker/") {
		target, _ := url.Parse(initializers.RCTRACKER)
		return "/rctracker", target
	} else if strings.HasPrefix(path, "/rcnotifications/") {
		target, _ := url.Parse(initializers.RCNOTIFICATIONS)
		return "/rcnotifications", target
	} else {
		return path, nil
	}
}
