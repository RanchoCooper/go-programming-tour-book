package main

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"

    http2 "go-programming-tour-book/blog-service/api/http"
    "go-programming-tour-book/blog-service/api/http/router"
    "go-programming-tour-book/blog-service/config"
    "go-programming-tour-book/blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

// @title 博客系统
// @version 1.0
// @description Go语言编程之旅
func main() {
    gin.SetMode(config.Config.Server.RunMode)
    router := router.NewRouter()

    s := &http.Server{
        Addr:           ":" + config.Config.Server.HTTPPort,
        Handler:        router,
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

func setupValidator() error {
    binding.Validator = http2.NewCustomValidator()
    return nil
}
