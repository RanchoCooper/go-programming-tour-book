package http

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"

    "blog-service/api/http/dto"
    "blog-service/api/http/errcode"
    "blog-service/api/http/handle"
    "blog-service/api/http/validator"
    "blog-service/internal/domain.model/service"
    "blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2022/1/5
 */

type Tag struct{}

func NewTag() Tag {
    return Tag{}
}

// List
// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
    param := dto.TagListRequest{}
    response := handle.NewResponse(c)
    valid, errs := validator.BindAndValid(c, &param, c.ShouldBindQuery)
    if !valid {
        logger.Log.Errorf(c, "tagList.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    response.ToResponse(gin.H{})
}

// Create
// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
    body := dto.CreateTagRequest{}
    response := handle.NewResponse(c)
    valid, errs := validator.BindAndValid(c, &body, c.ShouldBindJSON)
    if !valid {
        logger.Log.Errorf(c, "tagCreate.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }
    err := service.Service.TagService.Create(c, body)
    if err != nil {
        logger.Log.Errorf(c, "tagCreate.create fail, err: %v", err)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }
    response.ToResponse(gin.H{})
}

// Update
// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
    body := dto.UpdateTagRequest{
        ID: cast.ToUint32(c.Param("id")),
    }
    response := handle.NewResponse(c)
    valid, errs := validator.BindAndValid(c, &body, c.ShouldBindJSON)
    if !valid {
        logger.Log.Errorf(c, "tagUpdate.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }
    err := service.Service.TagService.Update(c, body)
    if err != nil {
        logger.Log.Errorf(c, "tagUpdate.update fail, err: %v", err)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }
    response.ToResponse(gin.H{})
}

// Delete
// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
    param := dto.DeleteTagRequest{}
    response := handle.NewResponse(c)
    valid, errs := validator.BindAndValid(c, &param, c.ShouldBindUri)
    if !valid {
        logger.Log.Errorf(c, "tagUpdate.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }
    err := service.Service.TagService.Delete(c, param)
    if err != nil {
        logger.Log.Errorf(c, "tagDelete.delete fail, err: %v", err)
        errResp := errcode.InvalidParams
        response.ToErrorResponse(errResp)
        return
    }
    response.ToResponse(gin.H{})
}
