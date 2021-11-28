package service

import (
    "context"

    "go-programming-tour-book/blog-service/global"
    "go-programming-tour-book/blog-service/internal/dao"
)

/**
 * @author Rancho
 * @date 2021/11/28
 */

type Service struct {
    ctx context.Context
    dao *dao.Dao
}

func New(ctx context.Context) Service {
    svc := Service{ctx: ctx}
    svc.dao = dao.New(global.DBEngine)

    return svc
}