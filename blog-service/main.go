package main

import (
    "net/http"
    "time"

    "github.com/RanchoCooper/go-programming-tour-book/blog-service/internal/routers"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

func main() {
    router := routers.NewRouter()

    s := &http.Server{
        Addr: ":8888",
        Handler: router,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    err := s.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
