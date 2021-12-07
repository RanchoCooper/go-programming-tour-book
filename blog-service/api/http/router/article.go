package router

import (
    "github.com/gin-gonic/gin"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Article struct {
}

func NewArticle() Article {
    return Article{}
}

func (a Article) Get(c *gin.Context) {

}

func (a Article) List(c *gin.Context) {

}

func (a Article) Create(c *gin.Context) {

}

func (a Article) Update(c *gin.Context) {

}

func (a Article) Delete(c *gin.Context) {

}
