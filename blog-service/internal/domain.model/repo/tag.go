package repo

import (
    "context"

    "blog-service/api/http/dto"
    "blog-service/internal/domain.model/entity"
)

/**
 * @author Rancho
 * @date 2022/1/6
 */

type ITagRepo interface {
    Create(ctx context.Context, dto dto.CreateTagRequest) (*entity.Tag, error)
    Delete(ctx context.Context, ID int) error
    Update(ctx context.Context, entity *entity.Tag) error
    GetList(ctx context.Context, entity *entity.Tag, pageOffset, pageSize int) (entities []*entity.Tag, err error)
    Count(ctx context.Context, entity *entity.Tag) (int64, error)
}
