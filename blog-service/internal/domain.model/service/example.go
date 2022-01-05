package service

import (
    "context"

    "github.com/jinzhu/copier"

    "blog-service/internal/domain.model/dto"
    "blog-service/internal/domain.model/repo"
    "blog-service/internal/port.adapter/repository"
    "blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/24
 */

type ExampleService struct {
    Repository repo.IExampleRepo
}

func NewExampleService(ctx context.Context) *ExampleService {
    srv := &ExampleService{Repository: repository.Example}
    logger.Log.Info(ctx, "example service init successfully")
    return srv
}

func (e *ExampleService) Create(ctx context.Context, v dto.CreateExampleReq) (*dto.CreateExampleResp, error) {
    result := &dto.CreateExampleResp{}
    model, err := e.Repository.Create(ctx, v)
    if err != nil {
        return nil, err
    }
    err = copier.Copy(result, model)
    if err != nil {
        return nil, err
    }

    return result, nil
}
