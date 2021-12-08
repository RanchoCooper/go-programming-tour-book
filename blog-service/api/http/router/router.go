package router

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "go-programming-tour-book/blog-service/api/http/middleware"
    handle2 "go-programming-tour-book/blog-service/api/http/router/handle"
    "go-programming-tour-book/blog-service/config"
    _ "go-programming-tour-book/blog-service/docs"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

func NewRouter() *gin.Engine {

    r := gin.Default()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(middleware.Translations())

    // swagger doc
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // static file
    r.POST("/file/file", handle2.UploadFile)
    r.StaticFile("/static", config.Config.App.UploadSavePath)

    // auth
    r.POST("/auth", handle2.GetAuth)

    apiV1 := r.Group("/api/v1")
    {
        // tags
        apiV1.POST("/tags", handle2.CreateTag)
        apiV1.DELETE("/tags/:id", handle2.DeleteTag)
        apiV1.PUT("/tags/:id", handle2.UpdateTag)
        apiV1.PATCH("/tags/:id/state", handle2.UpdateTag)
        apiV1.GET("/tags", handle2.ListTag)
        apiV1.GET("/tags/:id", handle2.GetTag)

        // articles
        apiV1.POST("/articles", handle2.CreateTag)
        apiV1.DELETE("/articles/:id", handle2.DeleteTag)
        apiV1.PUT("/articles/:id", handle2.UpdateTag)
        apiV1.PATCH("/articles/:id/state", handle2.UpdateTag)
        apiV1.GET("/articles/:id", handle2.ListTag)
        apiV1.GET("/articles", handle2.GetTag)
    }

    return r
}
