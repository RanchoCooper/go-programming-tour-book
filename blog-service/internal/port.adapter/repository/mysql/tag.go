package mysql

import (
    "context"
    "errors"

    "blog-service/api/http/dto"
    "blog-service/internal/domain.model/entity"
    "blog-service/internal/domain.model/repo"
)

/**
 * @author Rancho
 * @date 2022/1/6
 */

type Tag struct {
    IMySQL
}

var _ repo.ITagRepo = &Tag{}

func NewTag(mysql IMySQL) *Tag {
    return &Tag{
        IMySQL: mysql,
    }
}

func (t *Tag) Create(ctx context.Context, dto dto.CreateTagRequest) (*entity.Tag, error) {
    record := &entity.Tag{}
    record.Name = dto.Name
    record.CreatedBy = dto.CreateBy
    record.State = dto.State
    err := t.GetDB(ctx).Create(record).Error
    if err != nil {
        return nil, err
    }

    return record, nil
}

func (t *Tag) Delete(ctx context.Context, ID int) error {
    if ID == 0 {
        return errors.New("delete fail. need ID")
    }
    err := t.GetDB(ctx).Delete(&entity.Tag{}, ID).Error
    return err
}

func (t *Tag) Update(ctx context.Context, entity *entity.Tag) error {
    return t.GetDB(ctx).Table(entity.TableName()).Where("id = ? AND is_del = 0", entity.ID).
        Updates(entity.GetChangeMap()).Error
}

func (t *Tag) GetList(ctx context.Context, entity *entity.Tag, pageOffset, pageSize int) (entities []*entity.Tag, err error) {
    db := t.GetDB(ctx).Table(entity.TableName())
    if pageOffset >= 0 && pageSize > 0 {
        db = db.Offset(pageOffset).Limit(pageSize)
    }

    if entity.Name != "" {
        db = db.Where("name = ?", entity.Name)
    }
    db = db.Where("state = ?", entity.State)
    err = db.Where("is_del = 0").Find(entities).Error
    if err != nil {
        return nil, err
    }

    return
}

func (t *Tag) Count(ctx context.Context, entity *entity.Tag) (int64, error) {
    var count int64
    db := t.GetDB(ctx).Table(entity.TableName())
    if entity.Name != "" {
        db = db.Where("name = ?", entity.Name)
    }

    err := db.Where("state = ? AND is_del = 0", entity.State).Count(&count).Error
    if err != nil {
        return 0, err
    }

    return count, nil
}
