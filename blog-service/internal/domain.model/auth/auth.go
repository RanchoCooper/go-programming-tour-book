package auth

import (
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
