package service

import (
    "context"

    "blog-service/api/http/dto"
    "blog-service/internal/domain.model/repo"
    "blog-service/internal/port.adapter/repository"
    "blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2022/1/6
 */

type TagService struct {
    Repository repo.ITagRepo
}

func NewTagService(ctx context.Context) *TagService {
    srv := &TagService{Repository: repository.Tag}
    logger.Log.Info(ctx, "tag service init successfully")
    return srv
}

func (t *TagService) Create(ctx context.Context, dto dto.CreateTagRequest) error {
    _, err := t.Repository.Create(ctx, dto)
    if err != nil {
        return err
    }
    return nil
}
