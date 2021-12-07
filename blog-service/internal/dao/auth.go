package dao

import (
    "go-programming-tour-book/blog-service/internal/model"
)

/**
 * @author Rancho
 * @date 2021/12/7
 */

func (d *Dao) GetAuth(appKey, appSecret string) (*model.Auth, error) {
    auth := &model.Auth{
        AppKey:    appKey,
        AppSecret: appSecret,
    }

    return auth.Get(d.engine)
}
