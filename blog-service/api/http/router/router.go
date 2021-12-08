package router

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "go-programming-tour-book/blog-service/api/http/handle"
    "go-programming-tour-book/blog-service/api/http/middleware"
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
    r.POST("/file/file", handle.UploadFile)
    r.StaticFile("/static", config.Config.App.UploadSavePath)

    // auth
    r.POST("/auth", handle.GetAuth)

    apiV1 := r.Group("/api/v1")
    {
        // tags
        apiV1.POST("/tags", handle.CreateTag)
        apiV1.DELETE("/tags/:id", handle.DeleteTag)
        apiV1.PUT("/tags/:id", handle.UpdateTag)
        apiV1.PATCH("/tags/:id/state", handle.UpdateTag)
        apiV1.GET("/tags", handle.ListTag)
        apiV1.GET("/tags/:id", handle.GetTag)

        // articles
        apiV1.POST("/articles", handle.CreateTag)
        apiV1.DELETE("/articles/:id", handle.DeleteTag)
        apiV1.PUT("/articles/:id", handle.UpdateTag)
        apiV1.PATCH("/articles/:id/state", handle.UpdateTag)
        apiV1.GET("/articles/:id", handle.ListTag)
        apiV1.GET("/articles", handle.GetTag)
    }

    return r
}
