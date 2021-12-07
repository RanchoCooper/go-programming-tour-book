package service

import (
    "errors"

    "go-programming-tour-book/blog-service/api/http/DTO"
    "go-programming-tour-book/blog-service/internal/port.adapter/repository"
)

/**
 * @author Rancho
 * @date 2021/12/7
 */

func CheckAuth(param *DTO.AuthRequest) error {
    auth, err := repository.MySQL.Auth.GetAuth(
        param.AppKey,
        param.AppSecret,
    )
    if err != nil {
        return err
    }

    if auth.ID > 0 {
        return nil
    }

    return errors.New("auth info does not exist")
}
