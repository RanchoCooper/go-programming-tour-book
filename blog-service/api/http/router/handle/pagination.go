package handle

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"

    "go-programming-tour-book/blog-service/config"
)

/**
 * @author Rancho
 * @date 2021/11/27
 */

func GetPage(c *gin.Context) int {
    page := cast.ToInt(c.Query("page"))
    if page <= 0 {
        return 1
    }

    return 0
}

func GetPageSize(c *gin.Context) int {
    pageSize := cast.ToInt(c.Query("page_size"))
    if pageSize <= 0 {
        return config.Config.App.DefaultPageSize
    }
    if pageSize > config.Config.App.MaxPageSize {
        return config.Config.App.MaxPageSize
    }

    return pageSize
}

func GetPageOffset(page, pageSize int) int {
    result := 0
    if page > 0 {
        result = (page - 1) * pageSize
    }

    return result
}
