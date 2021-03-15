package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/RanchoCooper/go-programming-tour-book/blog-service/configs"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/global"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/internal/routers"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

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
