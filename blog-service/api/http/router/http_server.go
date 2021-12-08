package router

import (
    "context"
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
