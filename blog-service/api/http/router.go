package http

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "blog-service/api/http/middleware"
    "blog-service/config"
)

func NewServerRoute() *gin.Engine {
    if config.Config.App.Debug {
        gin.SetMode(gin.DebugMode)
    } else {
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.Default()
    router.Use(middleware.Translations())

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})

    })

    router.GET("/auth", GetAuth)

    example := router.Group("/example")
    {
        example.POST("", CreateExample)
    }

    apiv1 := router.Group("/api/v1")
    {
        tag := NewTag()
        // 创建标签
        apiv1.POST("/tags", tag.Create)
        // 删除指定标签
        apiv1.DELETE("/tags/:id", tag.Delete)
        // 更新指定标签
        apiv1.PUT("/tags/:id", tag.Update)
        // 获取标签列表
        apiv1.GET("/tags", tag.List)

        article := NewArticle()
        // 创建文章
        apiv1.POST("/articles", article.Create)
        // 删除指定文章
        apiv1.DELETE("/articles/:id", article.Delete)
        // 更新指定文章
        apiv1.PUT("/articles/:id", article.Update)
        // 获取指定文章
        apiv1.GET("/articles/:id", article.Get)
        // 获取文章列表
        apiv1.GET("/articles", article.List)
    }

    return router
}
