package router

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"

    "go-programming-tour-book/blog-service/config"
    "go-programming-tour-book/blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/8
 */

var server http.Server

func NewHTTPServer(ctx context.Context) {
    gin.SetMode(config.Config.Server.RunMode)
    r := NewRouter()

    server = http.Server{
        Addr:           ":" + config.Config.Server.HTTPPort,
        Handler:        r,
        ReadTimeout:    time.Duration(config.Config.Server.ReadTimeout) * time.Second,
        WriteTimeout:   time.Duration(config.Config.Server.WriteTimeout) * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    // test logger
    logger.Log.Infof(ctx, "%s: go-programming-tour-book/%s", "rancho", "blog-service")

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}

func Shutdown(ctx context.Context) {
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("server forced shutdown, err: ", err.Error())
    }
}
