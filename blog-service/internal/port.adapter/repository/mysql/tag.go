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
    DeleteTag(int64) (int64, error)
    UpdateTag(where map[string]interface{}, update map[string]interface{}) (int64, error)
    GetTag(uint32, uint8) (*tag.Tag, error)
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

func (tr TagRepo) DeleteTag(i int64) (int64, error) {
    // TODO implement me
    panic("implement me")
}

func (tr TagRepo) UpdateTag(where map[string]interface{}, update map[string]interface{}) (int64, error) {
    // TODO implement me
    panic("implement me")
}

func (tr TagRepo) GetTag(u uint32, u2 uint8) (*tag.Tag, error) {
    // TODO implement me
    panic("implement me")
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
