package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CorsMiddlewareOptions struct {
	AllowOrigin  *string
	AllowMethods *string
	AllowHeaders *string
	ContentType  *string
}

// CorsMiddleware Cors Middleware
func CorsMiddleware(options ...CorsMiddlewareOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
			return
		}

		allowOrigin := "*"
		allowMethods := "GET,POST,PUT,PATCH,DELETE,OPTIONS"
		allowHeaders := "accept,authorization,content-type,origin"
		contentType := "application/json"

		if len(options) > 0 {
			o := options[0]
			if o.AllowOrigin != nil {
				allowOrigin = *o.AllowOrigin
			}
			if o.AllowMethods != nil {
				allowMethods = *o.AllowMethods
			}
			if o.AllowHeaders != nil {
				allowHeaders = *o.AllowHeaders
			}
			if o.ContentType != nil {
				contentType = *o.ContentType
			}
		}

		c.Header("Access-Control-Allow-Origin", allowOrigin)
		c.Header("Access-Control-Allow-Methods", allowMethods)
		c.Header("Access-Control-Allow-Headers", allowHeaders)
		c.Header("Allow", allowMethods)
		c.Header("Content-Type", contentType)
		c.AbortWithStatus(http.StatusOK)
	}
}

// LoggerMiddleware Logger Middleware
func LoggerMiddleware() gin.HandlerFunc {
	if gin.IsDebugging() {
		return gin.Logger()
	}

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Ignore 200-299 or 304
		if param.StatusCode/100 == 2 || param.StatusCode == http.StatusNotModified {
			return ""
		}

		// Standardized access log format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

/**/
// Helpers
/**/

// NewMiddlewareExtender concat middleware
func NewMiddlewareExtender(baseMiddleware ...gin.HandlerFunc) func(handlerFunc ...gin.HandlerFunc) []gin.HandlerFunc {
	return func(middlewares ...gin.HandlerFunc) []gin.HandlerFunc {
		return append(baseMiddleware, middlewares...)
	}
}
