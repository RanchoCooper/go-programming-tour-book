package api

import (
	"fmt"

	"github.com/RanchoCooper/go-programming-tour-book/blog-service/global"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/internal/service"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/pkg/app"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {

}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	fmt.Println("result")
	fmt.Println(fileInfo)

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
