package entity

/**
 * @author Rancho
 * @date 2022/1/5
 */

type ArticleTag struct {
    *Model
    TagID     uint32 `json:"tag_id"`
    ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
    return "blog_article_tag"
}
