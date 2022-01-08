package service

import (
    "context"

    "github.com/RanchoCooper/structs"
    "github.com/jinzhu/copier"

    "blog-service/api/http/dto"
    "blog-service/internal/domain.model/entity"
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

func (t *TagService) Update(ctx context.Context, dto dto.UpdateTagRequest) error {
    tag := &entity.Tag{}
    _ = copier.Copy(tag, dto)
    tag.ChangeMap = structs.Map(tag)
    err := t.Repository.Update(ctx, tag)
    if err != nil {
        return err
    }
    return nil
}

func (t *TagService) Delete(ctx context.Context, dto dto.DeleteTagRequest) error {
    err := t.Repository.Delete(ctx, dto.ID)
    if err != nil {
        return err
    }
    return nil
}
