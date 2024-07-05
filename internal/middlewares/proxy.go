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
		_, target := getPrefix(path)

		if target != nil {
			log.Printf("Redirecting %s to %s\n", path, target.String())
			proxy := httputil.NewSingleHostReverseProxy(target)
			proxy.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func getPrefix(path string) (string, *url.URL) {
	if strings.HasPrefix(path, "/rcauth/") {
		target, _ := url.Parse(initializers.RCAUTH + path[7:])
		return "/rcauth/", target
	} else if strings.HasPrefix(path, "/rcstorage/") {
		target, _ := url.Parse(initializers.RCSTORAGE + path[10:])
		return "/rcstorage/", target
	} else {
		target, _ := url.Parse(initializers.RCGATEWAY + path)
		return path, target
	}
}