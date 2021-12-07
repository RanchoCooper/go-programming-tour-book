package router

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"

    "go-programming-tour-book/blog-service/api/http"
    "go-programming-tour-book/blog-service/api/http/errcode"
)

/**
 * @author Rancho
 * @date 2021/11/30
 */

type Upload struct {
}

func NewUpload() Upload {
    return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
    response := http.NewResponse(c)
    _, fileHeader, err := c.Request.FormFile("file")
    fileType := cast.ToInt(c.PostForm("type"))
    if err != nil {
        errResp := errcode.InvalidParams.WithDetails(err.Error())
        response.ToErrorResponse(errResp)
        return
    }

    if fileHeader == nil || fileType <= 0 {
        response.ToErrorResponse(errcode.InvalidParams)
        return
    }
}
