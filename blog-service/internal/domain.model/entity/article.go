package entity

/**
 * @author Rancho
 * @date 2022/1/5
 */

type Article struct {
    *Model
    Title         string `json:"title"`
    Desc          string `json:"desc"`
    Content       string `json:"content"`
    CoverImageUrl string `json:"cover_image_url"`
    State         uint8  `json:"state"`
}

func (a Article) TableName() string {
    return "blog_article"
}
