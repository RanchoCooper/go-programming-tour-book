package repo

import (
    "context"

    "blog-service/internal/domain.model/dto"
    "blog-service/internal/domain.model/entity"
)

/**
 * @author Rancho
 * @date 2021/12/24
 */

type IExampleRepo interface {
    Create(ctx context.Context, dto dto.CreateExampleReq) (*entity.Example, error)
    Delete(ctx context.Context, ID int) error
    Save(ctx context.Context, entity *entity.Example) error
    Get(ctx context.Context, ID int) (obj *entity.Example, e error)
    FindByName(ctx context.Context, name string) (*entity.Example, error)
}
