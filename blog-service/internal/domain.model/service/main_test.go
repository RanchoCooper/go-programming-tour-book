package service

import (
    "context"
    "testing"

    "blog-service/internal/port.adapter/repository"
)

/**
 * @author Rancho
 * @date 2021/12/25
 */

var ctx = context.Background()

func TestMain(m *testing.M) {
    repository.Init(
        repository.WithMySQL(ctx),
        repository.WithRedis(ctx),
    )
    m.Run()
}
