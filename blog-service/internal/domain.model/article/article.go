package article

import (
    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Article struct {
    gorm.Model
    Title         string `json:"title"`
    Desc          string `json:"desc"`
    Content       string `json:"content"`
    CoverImageUrl string `json:"cover_image_url"`
    State         uint8  `json:"state"`
}

func (a Article) TableName() string {
    return "blog_article"
}
