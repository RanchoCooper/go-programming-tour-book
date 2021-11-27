package app

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "go-programming-tour-book/blog-service/pkg/errcode"
)

/**
 * @author Rancho
 * @date 2021/11/27
 */

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
    return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
    if data == nil {
        data = gin.H{}
    }
    r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
    r.Ctx.JSON(http.StatusOK, gin.H{
        "list": list,
        "pager": Pager{
            Page:      GetPage(r.Ctx),
            PageSize:  GetPageSize(r.Ctx),
            TotalRows: totalRows,
        },
    })
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
    response := gin.H{
        "code": err.Code,
        "msg": err.Msg,
    }
    if len(err.Details) > 0 {
        response["details"] = err.Details
    }
    r.Ctx.JSON(err.StatusCode(), response)
}