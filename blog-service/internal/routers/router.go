package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    _ "go-programming-tour-book/blog-service/docs"
    "go-programming-tour-book/blog-service/global"
    "go-programming-tour-book/blog-service/internal/middleware"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

func NewRouter() *gin.Engine {
    tag := NewTag()
    article := NewArticle()
    upload := NewUpload()

    r := gin.Default()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(middleware.Translations())

    // swagger doc
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // static file
    r.POST("/upload/file", upload.UploadFile)
    r.StaticFile("/static", global.AppSetting.UploadSavePath)

    // auth
    r.POST("/auth", GetAuth)

    apiV1 := r.Group("/api/v1")
    {
        // tags
        apiV1.POST("/tags", tag.Create)
        apiV1.DELETE("/tags/:id", tag.Delete)
        apiV1.PUT("/tags/:id", tag.Update)
        apiV1.PATCH("/tags/:id/state", tag.Update)
        apiV1.GET("/tags", tag.List)
        apiV1.GET("/tags/:id", tag.Get)

        // articles
        apiV1.POST("/articles", article.Create)
        apiV1.DELETE("/articles/:id", article.Delete)
        apiV1.PUT("/articles/:id", article.Update)
        apiV1.PATCH("/articles/:id/state", article.Update)
        apiV1.GET("/articles/:id", article.List)
        apiV1.GET("/articles", article.Get)
    }

    return r
}
