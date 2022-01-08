package handle

import (
    "context"
    "flag"
    "testing"

    "blog-service/config"
    "blog-service/internal/domain.model/service"
    "blog-service/internal/port.adapter/repository"
)

/**
 * @author Rancho
 * @date 2022/1/7
 */

var ctx = context.Background()

func TestMain(m *testing.M) {
    if err := flag.Set("cf", "../../config/config.yaml"); err != nil {
        panic(err)
    }
    config.Init()
    repository.Init(
        repository.WithMySQL(ctx),
        repository.WithRedis(ctx),
    )
    service.Init(ctx)

    m.Run()
}
