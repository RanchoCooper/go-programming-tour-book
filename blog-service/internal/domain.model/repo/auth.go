package repo

import (
    "context"

    "blog-service/api/http/dto"
    "blog-service/internal/domain.model/entity"
)

/**
 * @author Rancho
 * @date 2022/1/7
 */

type IAuthRepo interface {
    Get(ctx context.Context, dto *dto.AuthRequest) (*entity.Auth, error)
}
