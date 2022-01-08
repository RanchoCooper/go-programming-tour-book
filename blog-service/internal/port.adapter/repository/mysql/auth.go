package mysql

import (
    "context"

    "gorm.io/gorm"

    "blog-service/api/http/dto"
    "blog-service/internal/domain.model/entity"
    "blog-service/internal/domain.model/repo"
)

/**
 * @author Rancho
 * @date 2022/1/7
 */

type Auth struct {
    IMySQL
}

var _ repo.IAuthRepo = &Auth{}

func NewAuth(mysql IMySQL) *Auth {
    return &Auth{IMySQL: mysql}
}

func (a *Auth) Get(ctx context.Context, dto *dto.AuthRequest) (*entity.Auth, error) {
    record := &entity.Auth{}
    record.AppKey = dto.AppKey
    record.AppSecret = dto.AppSecret
    err := a.GetDB(ctx).Where("app_key = ? AND app_secret = ? AND is_del = 0", record.AppKey, record.AppSecret).
        First(record).Error

    if err != nil && err != gorm.ErrRecordNotFound {
        return record, err
    }

    return record, nil
}
