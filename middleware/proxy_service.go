package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// Proxy 代理判断
func Proxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		isTrue := viper.GetBool("is_true")
		urls := viper.GetString("wx_url")
		if isTrue && urls != "" {
			req, err := http.NewRequestWithContext(c, c.Request.Method, urls+c.Request.RequestURI, c.Request.Body)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}
			defer req.Body.Close()
			req.Header = c.Request.Header
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}
			// header 也带过来
			for k := range resp.Header {
				for j := range resp.Header[k] {
					c.Header(k, resp.Header[k][j])
				}
			}
			extraHeaders := make(map[string]string)
			extraHeaders["Proxy-Type"] = "Proxy"
			c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, extraHeaders)
			c.Abort()
			return
		}
		c.Next()
	}
}
