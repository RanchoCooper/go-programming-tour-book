package handle

import (
    "fmt"

    "github.com/gin-gonic/gin"

    "go-programming-tour-book/blog-service/api/http/DTO"
    "go-programming-tour-book/blog-service/api/http/errcode"
    "go-programming-tour-book/blog-service/internal/port.adapter/service"
    "go-programming-tour-book/blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/7
 */

func GetAuth(c *gin.Context) {
    param := DTO.AuthRequest{}
    response := NewResponse(c)
    valid, errs := BindAndValid(c, &param)
    fmt.Println(param.AppKey)
    fmt.Println(param.AppSecret)
    if !valid {
        logger.Log.Errorf(c, "app.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    err := service.CheckAuth(&param)
    if err != nil {
        logger.Log.Errorf(c, "svc.CheckAuth err: %v", err)
        response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
        return
    }

    token, err := GenerateToken(param.AppKey, param.AppSecret)
    if err != nil {
        logger.Log.Errorf(c, "app.GenerateToken err: %v", err)
        response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
        return
    }

    response.ToResponse(gin.H{
        "token": token,
    })
}
