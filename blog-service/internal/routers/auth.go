package routers

import (
    "fmt"

    "github.com/gin-gonic/gin"

    "go-programming-tour-book/blog-service/global"
    "go-programming-tour-book/blog-service/internal/service"
    "go-programming-tour-book/blog-service/pkg/app"
    "go-programming-tour-book/blog-service/pkg/errcode"
)

/**
 * @author Rancho
 * @date 2021/12/7
 */

func GetAuth(c *gin.Context) {
    param := service.AuthRequest{}
    response := app.NewResponse(c)
    valid, errs := app.BindAndValid(c, &param)
    fmt.Println(param.AppKey)
    fmt.Println(param.AppSecret)
    if !valid {
        global.Logger.Errorf("app.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    svc := service.New(c.Request.Context())
    err := svc.CheckAuth(&param)
    if err != nil {
        global.Logger.Errorf("svc.CheckAuth err: %v", err)
        response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
        return
    }

    token, err := app.GenerateToken(param.AppKey, param.AppSecret)
    if err != nil {
        global.Logger.Errorf("app.GenerateToken err: %v", err)
        response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
        return
    }

    response.ToResponse(gin.H{
        "token": token,
    })
}
