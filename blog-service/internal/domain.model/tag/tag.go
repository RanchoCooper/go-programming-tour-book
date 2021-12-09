package tag

import (
    "time"

    "gorm.io/gorm"

    "go-programming-tour-book/blog-service/internal/port.adapter/repository"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Tag struct {
    ID        uint   `gorm:"primarykey"`
    Name      string `json:"name"`
    State     *uint8 `json:"state"`
    CreatedBy string `json:"created_by"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t Tag) TableName() string {
    return "blog_tag"
}

func (t *Tag) GetTag() (*Tag, error) {
    return repository.MySQL.Tag.GetTag(t)
}

func (t *Tag) CountTag() (int64, error) {
    return repository.MySQL.Tag.CountTag(t)
}

func (t *Tag) GetTagList(offset, limit int) ([]*Tag, error) {
    return repository.MySQL.Tag.GetTagList(t, offset, limit)
}

func (t *Tag) CreateTag() (*Tag, error) {
    return repository.MySQL.Tag.CreateTag(t)
}

func (t *Tag) UpdateTag() error {
    where := map[string]interface{}{
        "id": t.ID,
    }
    var update map[string]interface{}
    if t.Name != "" {
        update["name"] = t.Name
    }
    if t.CreatedBy != "" {
        update["created_by"] = t.CreatedBy
    }
    if t.State != nil {
        update["state"] = t.State
    }
    return repository.MySQL.Tag.UpdateTag(where, update)
}

func (t *Tag) DeleteTag() error {
    return repository.MySQL.Tag.DeleteTag(t)
}
