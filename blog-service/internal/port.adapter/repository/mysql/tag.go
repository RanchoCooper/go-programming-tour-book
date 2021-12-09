package mysql

import (
    "gorm.io/gorm"

    "go-programming-tour-book/blog-service/internal/domain.model/tag"
)

/**
 * @author Rancho
 * @date 2021/11/28
 */

type ITag interface {
    CreateTag(*tag.Tag) (*tag.Tag, error)
    DeleteTag(*tag.Tag) error
    UpdateTag(where map[string]interface{}, update map[string]interface{}) error
    GetTag(*tag.Tag) (*tag.Tag, error)
    GetTagList(*tag.Tag, int, int) ([]*tag.Tag, error)
    CountTag(*tag.Tag) (int64, error)
}

type TagRepo struct {
    db *gorm.DB
}

func (tr TagRepo) CreateTag(t *tag.Tag) (*tag.Tag, error) {
    err := tr.db.Create(t).Error
    return t, err
}

func (tr TagRepo) DeleteTag(t *tag.Tag) error {
    return tr.db.Delete(t).Error
}

func (tr TagRepo) UpdateTag(where map[string]interface{}, update map[string]interface{}) error {
    return tr.db.Model(&tr).Where(where).Updates(update).Error
}

func (tr TagRepo) GetTag(t *tag.Tag) (*tag.Tag, error) {
    err := tr.db.Model(t).Last(t).Error
    return t, err
}

func (tr TagRepo) GetTagList(t *tag.Tag, limit, offset int) ([]*tag.Tag, error) {
    var tags []*tag.Tag
    var err error
    if limit >= 0 && offset >= 0 {
        tr.db.Limit(limit).Offset(offset)
    }
    err = tr.db.Model(t).Find(&tags).Error
    if err != nil {
        return nil, err
    }

    return tags, nil
}

func (tr TagRepo) CountTag(t *tag.Tag) (int64, error) {
    var count int64
    err := tr.db.Model(t).Count(&count).Error
    return count, err
}

var _ ITag = &TagRepo{}

func NewTagRepository(db *gorm.DB) *TagRepo {
    return &TagRepo{
        db: db,
    }
}
