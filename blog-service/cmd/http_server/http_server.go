package http_server

import (
    "context"
    "fmt"
    "net/http"

    http2 "blog-service/api/http/handle"
    "blog-service/config"
    "blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2022/1/5
 */

func Start(ctx context.Context, errChan chan error, httpCloseCh chan struct{}) {
    // init server
    srv := &http.Server{
        Addr:    config.Config.HTTPServer.Addr,
        Handler: http2.NewServerRoute(),
    }

    // run server
    go func() {
        logger.Log.Info(ctx, fmt.Sprintf("%s HTTP server is starting on %s", config.Config.App.Name, config.Config.HTTPServer.Addr))
        errChan <- srv.ListenAndServe()
    }()

    // watch the ctx exit
    go func() {
        <-ctx.Done()
        if err := srv.Shutdown(ctx); err != nil {
            logger.Log.Info(ctx, fmt.Sprintf("httpServer shutdown:%v", err))
        }
        httpCloseCh <- struct{}{}
    }()
}