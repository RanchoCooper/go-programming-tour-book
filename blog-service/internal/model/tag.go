package model

import (
    "errors"

    "go-programming-tour-book/blog-service/pkg/app"

    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
    List []*Tag
    Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
    var count int64
    if t.Name != "" {
        db = db.Where("name = ?", t.Name)
    }
    db = db.Where("state = ?", t.State)
    err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (t Tag) Get(db *gorm.DB) (*Tag, error) {
    var tag *Tag
    var err error
    if t.ID == 0 {
        return nil, errors.New("ID is empty")
    }
    err = db.Table(t.TableName()).Take(tag).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }

    return tag, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
    var tags []*Tag
    var err error

    if pageOffset >= 0 && pageSize > 0 {
        db = db.Offset(pageOffset).Limit(pageSize)
    }
    if t.Name != "" {
        db = db.Where("name = ?", t.Name)
    }
    db = db.Where("state = ?", t.State)
    if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
        return nil, err
    }

    return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
    return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
    return db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
    return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}