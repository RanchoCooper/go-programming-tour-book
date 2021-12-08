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
    GetTagList(string, uint32, int) ([]*tag.Tag, error)
    CountTag() (int64, error)
}

type TagRepo struct {
    db *gorm.DB
}

func (t TagRepo) CreateTag(t2 *tag.Tag) (*tag.Tag, error) {
    err := t.db.Create(t2).Error
    return t2, err
}

func (t TagRepo) DeleteTag(i int64) (int64, error) {
    // TODO implement me
    panic("implement me")
}

func (t TagRepo) UpdateTag(where map[string]interface{}, update map[string]interface{}) (int64, error) {
    // TODO implement me
    panic("implement me")
}

func (t TagRepo) GetTag(u uint32, u2 uint8) (*tag.Tag, error) {
    // TODO implement me
    panic("implement me")
}

func (t TagRepo) GetTagList(s string, u uint32, i int) ([]*tag.Tag, error) {
    // TODO implement me
    panic("implement me")
}

func (t TagRepo) CountTag() (int64, error) {
    // TODO implement me
    panic("implement me")
}

var _ ITag = &TagRepo{}

func NewTagRepository(db *gorm.DB) *TagRepo {
    return &TagRepo{
        db: db,
    }
}
