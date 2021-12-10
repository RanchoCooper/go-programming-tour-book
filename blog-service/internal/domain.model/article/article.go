package article

import (
    "gorm.io/gorm"

    "go-programming-tour-book/blog-service/internal/port.adapter/repository"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type IArticle interface {
    CreateArticle(*Article) (*Article, error)
    DeleteArticle(*Article) error
    UpdateArticle(where map[string]interface{}, update map[string]interface{}) error
    GetArticle(*Article) (*Article, error)
    GetArticleList(*Article, int, int) ([]*Article, error)
    CountArticle(*Article) (int64, error)
}

type Article struct {
    gorm.Model
    Title         string `json:"title"`
    Desc          string `json:"desc"`
    Content       string `json:"content"`
    CoverImageUrl string `json:"cover_image_url"`
    State         *uint8 `json:"state"`
}

func (a Article) TableName() string {
    return "blog_article"
}

func (a *Article) GetArticle() (*Article, error) {
    return repository.MySQL.Article.GetArticle(a)
}

func (a *Article) CountArticle() (int64, error) {
    return repository.MySQL.Article.CountArticle(a)
}

func (a *Article) GetArticleList(offset, limit int) ([]*Article, error) {
    return repository.MySQL.Article.GetArticleList(a, offset, limit)
}

func (a *Article) CreateArticle() (*Article, error) {
    return repository.MySQL.Article.CreateArticle(a)
}

func (a *Article) UpdateArticle() error {
    where := map[string]interface{}{
        "id": a.ID,
    }
    var update map[string]interface{}
    if a.Title != "" {
        update["title"] = a.Title
    }
    if a.Desc != "" {
        update["desc"] = a.Desc
    }
    if a.Content != "" {
        update["content"] = a.Content
    }
    if a.CoverImageUrl != "" {
        update["cover_image_url"] = a.CoverImageUrl
    }
    if a.State != nil {
        update["state"] = a.State
    }
    return repository.MySQL.Article.UpdateArticle(where, update)
}

func (a *Article) DeleteArticle() error {
    return repository.MySQL.Article.DeleteArticle(a)
}
