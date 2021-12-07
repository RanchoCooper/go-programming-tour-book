package routers

import (
    "github.com/gin-gonic/gin"

    "go-programming-tour-book/blog-service/global"
    "go-programming-tour-book/blog-service/internal/service"
    "go-programming-tour-book/blog-service/pkg/app"
    "go-programming-tour-book/blog-service/pkg/convert"
    "go-programming-tour-book/blog-service/pkg/errcode"
    "go-programming-tour-book/blog-service/pkg/upload"
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
    response := app.NewResponse(c)
    file, fileHeader, err := c.Request.FormFile("file")
    fileType := convert.StrTo(c.PostForm("type")).MustInt()
    if err != nil {
        errResp := errcode.InvalidParams.WithDetails(err.Error())
        response.ToErrorResponse(errResp)
        return
    }

    if fileHeader == nil || fileType <= 0 {
        response.ToErrorResponse(errcode.InvalidParams)
        return
    }
    svc := service.New(c.Request.Context())
    fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
    if err != nil {
        global.Logger.Errorf("svc.UploadFile err: %v", err)
        errResp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
        response.ToErrorResponse(errResp)
        return
    }

    response.ToResponse(gin.H{
        "file_access_url": fileInfo.AccessUrl,
    })
}
