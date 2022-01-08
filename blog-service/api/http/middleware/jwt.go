package middleware

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"

    "blog-service/api/http/errcode"
    "blog-service/api/http/handle"
    jwt2 "blog-service/api/http/jwt"
)

/**
 * @author Rancho
 * @date 2022/1/8
 */

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
            ecode = errcode.UnauthorizedNeedToken
        } else {
            _, err := jwt2.ParseToken(token)
            if err != nil {
                switch err.(*jwt.ValidationError).Errors {
                case jwt.ValidationErrorExpired:
                    ecode = errcode.UnauthorizedTokenTimeout
                default:
                    ecode = errcode.UnauthorizedTokenError
                }
            }
        }

        if ecode != errcode.Success {
            response := handle.NewResponse(c)
            response.ToErrorResponse(ecode)
            c.Abort()
            return
        }

        c.Next()
    }
}
