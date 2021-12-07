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
    GetTag(uint32, uint8) (*tag.Tag, error)
    GetTagList(string, uint32, int) ([]*tag.Tag, error)
}

type TagRepo struct {
    db *gorm.DB
}

func (t TagRepo) GetTag(u uint32, u2 uint8) (*tag.Tag, error) {
    // TODO implement me
    panic("implement me")
}

func (t TagRepo) GetTagList(s string, u uint32, i int) ([]*tag.Tag, error) {
    // TODO implement me
    panic("implement me")
}

var _ ITag = &TagRepo{}

func NewTagRepository(db *gorm.DB) *TagRepo {
    return &TagRepo{
        db: db,
    }
}
