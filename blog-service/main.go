package main

import (
    "github.com/gin-gonic/gin"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

func main() {
   r := gin.Default()

   r.GET("/ping", func(context *gin.Context) {
       context.JSON(200, gin.H{"message": "pong"})
   })

   r.Run(":8888")
}
