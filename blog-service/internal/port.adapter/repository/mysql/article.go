package mysql

import (
    "gorm.io/gorm"

    "go-programming-tour-book/blog-service/internal/domain.model/article"
)

/**
 * @author Rancho
 * @date 2021/12/9
 */

type IArticle interface {
    CreateArticle(*article.Article) (*article.Article, error)
    DeleteArticle(*article.Article) error
    UpdateArticle(where map[string]interface{}, update map[string]interface{}) error
    GetArticle(*article.Article) (*article.Article, error)
    GetArticleList(*article.Article, int, int) ([]*article.Article, error)
    CountArticle(*article.Article) (int64, error)
}

type ArticleRepo struct {
    db *gorm.DB
}

func (tr ArticleRepo) CreateArticle(t *article.Article) (*article.Article, error) {
    err := tr.db.Create(t).Error
    return t, err
}

func (tr ArticleRepo) DeleteArticle(t *article.Article) error {
    return tr.db.Delete(t).Error
}

func (tr ArticleRepo) UpdateArticle(where map[string]interface{}, update map[string]interface{}) error {
    return tr.db.Model(&tr).Where(where).Updates(update).Error
}

func (tr ArticleRepo) GetArticle(t *article.Article) (*article.Article, error) {
    err := tr.db.Model(t).Last(t).Error
    return t, err
}

func (tr ArticleRepo) GetArticleList(t *article.Article, limit, offset int) ([]*article.Article, error) {
    var Articles []*article.Article
    var err error
    if limit >= 0 && offset >= 0 {
        tr.db.Limit(limit).Offset(offset)
    }
    err = tr.db.Model(t).Find(&Articles).Error
    if err != nil {
        return nil, err
    }

    return Articles, nil
}

func (tr ArticleRepo) CountArticle(t *article.Article) (int64, error) {
    var count int64
    err := tr.db.Model(t).Count(&count).Error
    return count, err
}

var _ IArticle = &ArticleRepo{}

func NewArticleRepository(db *gorm.DB) *ArticleRepo {
    return &ArticleRepo{
        db: db,
    }
}
