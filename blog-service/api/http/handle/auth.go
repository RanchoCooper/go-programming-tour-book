package handle

import (
    "github.com/gin-gonic/gin"

    "blog-service/api/http/dto"
    "blog-service/api/http/errcode"
    "blog-service/api/http/jwt"
    "blog-service/api/http/validator"
    "blog-service/internal/domain.model/service"
    "blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2022/1/7
 */

func GetAuth(c *gin.Context) {
    param := &dto.AuthRequest{}
    response := NewResponse(c)
    valid, errs := validator.BindAndValid(c, param, c.ShouldBindJSON)
    if !valid {
        logger.Log.Errorf(c, "getAuth.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    err := service.Service.AuthService.CheckAuth(c, param)
    if err != nil {
        logger.Log.Errorf(c, "CheckAuth err: %v", err)
        response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
        return
    }

    token, err := jwt.GenerateToken(param.AppKey, param.AppSecret)
    if err != nil {
        logger.Log.Errorf(c, "GenerateToken err: %v", err)
        response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
        return
    }

    response.ToResponse(gin.H{
        "token": token,
    })
}
