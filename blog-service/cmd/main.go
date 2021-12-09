package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"
    "time"

    "go-programming-tour-book/blog-service/api/http/router"
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
    ctx := context.Background()
    go router.NewHTTPServer(ctx)

    // graceful shutdown
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    logger.Log.Info(ctx, "shutdown server...")
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    router.Shutdown(ctx)
    logger.Log.Info(ctx, "server exit")
}
