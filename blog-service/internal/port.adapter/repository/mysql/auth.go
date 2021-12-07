package mysql

import (
    "errors"

    "gorm.io/gorm"

    "go-programming-tour-book/blog-service/internal/domain.model/auth"
)

/**
 * @author Rancho
 * @date 2021/12/7
 */

type IAuth interface {
    GetAuth(string, string) (auth.Auth, error)
}

type AuthRepo struct {
    db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepo {
    return &AuthRepo{
        db: db,
    }
}

func (a *AuthRepo) GetAuth(appKey, appSecret string) (*auth.Auth, error) {
    var auth *auth.Auth
    err := a.db.Where("app_key = ? AND app_secret = ? AND is_del = ?", appKey, appSecret, 0).First(auth).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }

    return auth, nil
}
