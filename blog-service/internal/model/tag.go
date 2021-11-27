package model

import (
    "go-programming-tour-book/blog-service/pkg/app"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
    List []*Tag
    Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
