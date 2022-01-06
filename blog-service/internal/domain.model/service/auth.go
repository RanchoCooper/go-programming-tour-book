package service

import (
    "context"

    "github.com/pkg/errors"

    "blog-service/api/http/dto"
    "blog-service/internal/domain.model/repo"
    "blog-service/internal/port.adapter/repository"
    "blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2022/1/7
 */

type AuthService struct {
    Repository repo.IAuthRepo
}

func NewAuthService(ctx context.Context) *AuthService {
    srv := &AuthService{Repository: repository.Auth}
    logger.Log.Info(ctx, "auth service init successfully")
    return srv
}

func (a *AuthService) CheckAuth(ctx context.Context, dto *dto.AuthRequest) error {
    auth, err := a.Repository.Get(ctx, dto)
    if err != nil {
        return err
    }

    if auth.ID > 0 {
        return nil
    }

    return errors.New("auth not exists")
}
