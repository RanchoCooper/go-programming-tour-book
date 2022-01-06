package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "blog-service/cmd/http_server"
    "blog-service/config"
    "blog-service/internal/domain.model/service"
    "blog-service/internal/port.adapter/repository"
    "blog-service/util/logger"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    initConfig()
    initRuntime(ctx)
    initServer(ctx, cancel)
}

func initConfig() {
    config.Init()
}

func initRuntime(ctx context.Context) {
    repository.Init(
        repository.WithMySQL(ctx),
        repository.WithRedis(ctx),
    )
    service.Init(ctx)
}

func initServer(ctx context.Context, cancel context.CancelFunc) {
    errCh := make(chan error)
    httpCloseCh := make(chan struct{})
    http_server.Start(ctx, errCh, httpCloseCh)

    // graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    select {
    case <-quit:
        cancel()
        logger.Log.Info(ctx, "Start graceful shutdown")
    case err := <-errCh:
        cancel()
        logger.Log.Error(ctx, fmt.Sprintf("http err:%v", err))
    }
    <-httpCloseCh
    logger.Log.Infof(ctx, "%s HTTP server exit!", config.Config.App.Name)
}
