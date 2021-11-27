package routers

import (
    "go-programming-tour-book/blog-service/pkg/app"
    "go-programming-tour-book/blog-service/pkg/errcode"

    "github.com/gin-gonic/gin"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Tag struct {

}

func NewTag() Tag {
    return Tag{}
}

func (t Tag) Get(c *gin.Context) {
    app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

func (t Tag) List(c *gin.Context) {

}

func (t Tag) Create(c *gin.Context) {

}

func (t Tag) Update(c *gin.Context) {

}

func (t Tag) Delete(c *gin.Context) {

}
