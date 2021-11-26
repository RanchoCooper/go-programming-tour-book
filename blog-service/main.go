package main

import (
	"log"
	"net/http"
    "os"
    "time"

    "go-programming-tour-book/blog-service/global"
    "go-programming-tour-book/blog-service/internal/model"
    "go-programming-tour-book/blog-service/internal/routers"
	"go-programming-tour-book/blog-service/pkg/setting"

    "github.com/gin-gonic/gin"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

func init() {
    var err error
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

    err = setupDBEngine()
    if err != nil {
        log.Fatalf("init.setupDBEngine err: %v", err)
    }

}

func setupSetting() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = settings.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
    global.DatabaseSetting.Password = os.Getenv("MYSQL_PASSWORD")

    return nil
}

func setupDBEngine() error {
    var err error
    global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HTTPPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
