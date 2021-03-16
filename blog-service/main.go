package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/RanchoCooper/go-programming-tour-book/blog-service/configs"
	_ "github.com/RanchoCooper/go-programming-tour-book/blog-service/docs"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/global"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/internal/model"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/internal/routers"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/pkg/logger"
)

func init() {
	// init setting
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	// init database
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	// init logger
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func setupSetting() error {
	setting, err := configs.NewSetting()
	if err != nil {
		return nil
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return nil
	}
	fmt.Printf("ServerSetting: %v\nAppSetting: %v\nDatabaseSetting: %v\n", global.ServerSetting, global.AppSetting, global.DatabaseSetting)
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(*global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	global.Logger.Infof(context.Background(), "%s: go-programming-tour-book/%s", "rancho", "blog-service")
	return nil
}
