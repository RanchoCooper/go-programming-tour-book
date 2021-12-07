package article_tag

import (
    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type ArticleTag struct {
    gorm.Model
    TagID     uint32 `json:"tag_id"`
    ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
    return "blog_article_tag"
}
