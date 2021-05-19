package middleware

import (
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/pkg/app"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		if s, exists := c.GetQuery("token"); exists {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			// FIXME
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
