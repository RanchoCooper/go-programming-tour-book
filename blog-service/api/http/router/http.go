package router

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"

    "go-programming-tour-book/blog-service/api/http/errcode"
    "go-programming-tour-book/blog-service/config"
    "go-programming-tour-book/blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2021/11/27
 */

type Response struct {
    Ctx *gin.Context
}

type Pager struct {
    Page      int   `json:"page"`
    PageSize  int   `json:"page_size"`
    TotalRows int64 `json:"total_rows"`
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

func (r *Response) ToResponseList(list interface{}, totalRows int64) {
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
        "msg":  err.Msg,
    }
    if len(err.Details) > 0 {
        response["details"] = err.Details
    }
    r.Ctx.JSON(err.StatusCode(), response)
}

func NewHTTPServer() {
    gin.SetMode(config.Config.Server.RunMode)
    r := NewRouter()

    s := &http.Server{
        Addr:           ":" + config.Config.Server.HTTPPort,
        Handler:        r,
        ReadTimeout:    time.Duration(config.Config.Server.ReadTimeout) * time.Second,
        WriteTimeout:   time.Duration(config.Config.Server.WriteTimeout) * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    // test logger
    logger.Log.Infof(context.Background(), "%s: go-programming-tour-book/%s", "rancho", "blog-service")

    err := s.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
