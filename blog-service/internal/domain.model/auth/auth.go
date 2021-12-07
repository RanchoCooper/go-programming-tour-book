package auth

import (
    "errors"

    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/12/7
 */

type Auth struct {
    gorm.Model
    AppKey    string `gorm:"column:app_key" json:"app_key"`
    AppSecret string `gorm:"column:app_secret" json:"app_secret"`
}

func (a Auth) TableName() string {
    return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (*Auth, error) {
    var auth *Auth
    db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
    err := db.First(auth).Error
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return auth, err
    }

    return auth, nil
}
